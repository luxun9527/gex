package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/model"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/query"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/proto/define"
	"gopkg.in/yaml.v3"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSymbolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSymbolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSymbolLogic {
	return &AddSymbolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSymbolLogic) AddSymbol(req *types.AddSymbolReq) (resp *types.AddSymbolResp, err error) {

	var (
		coin   = l.svcCtx.Query.Coin
		symbol = l.svcCtx.Query.Symbol
	)
	baseCoinInfo, err := coin.WithContext(l.ctx).Where(coin.ID.Eq(req.BaseCoinID)).Take()
	if err != nil {
		return nil, err
	}
	quoteCoinInfo, err := coin.WithContext(l.ctx).Where(coin.ID.Eq(uint32(req.QuoteCoinID))).Take()
	if err != nil {
		return nil, err
	}
	symbolName := baseCoinInfo.CoinName + "_" + quoteCoinInfo.CoinName
	c := &model.Symbol{
		SymbolName:    symbolName,
		BaseCoinID:    req.BaseCoinID,
		BaseCoinName:  baseCoinInfo.CoinName,
		BaseCoinPrec:  baseCoinInfo.Prec,
		QuoteCoinID:   uint32(req.QuoteCoinID),
		QuoteCoinName: quoteCoinInfo.CoinName,
		QuoteCoinPrec: quoteCoinInfo.Prec,
	}

	count, err := symbol.WithContext(l.ctx).Where(symbol.SymbolName.Eq(symbolName)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errs.DuplicateDataErr
	}
	if err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		if err := tx.Symbol.WithContext(l.ctx).Create(c); err != nil {
			logx.Errorw("create symbol failed", logx.Field("err", err))
			return err
		}
		symbols, err := tx.Symbol.WithContext(l.ctx).Find()
		if err != nil {
			return err
		}
		for _, v := range symbols {
			symbolInfo := &define.SymbolInfo{
				SymbolName:         v.SymbolName,
				SymbolID:           int32(v.ID),
				BaseCoinName:       v.BaseCoinName,
				BaseCoinID:         int32(v.BaseCoinID),
				QuoteCoinName:      v.QuoteCoinName,
				QuoteCoinID:        int32(v.QuoteCoinID),
				BaseCoinPrecValue:  v.BaseCoinPrec,
				QuoteCoinPrecValue: v.QuoteCoinPrec,
			}
			data, err := yaml.Marshal(symbolInfo)
			if err != nil {
				return err
			}
			if _, err := l.svcCtx.EtcdCli.Put(l.ctx, define.EtcdSymbolPrefix+v.SymbolName, string(data)); err != nil {
				logx.Errorw("put config to etcd failed", logx.Field("err", err))
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return
}
