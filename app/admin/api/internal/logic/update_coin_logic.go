package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/model"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/query"
	"github.com/luxun9527/gex/common/proto/define"
	"gopkg.in/yaml.v3"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCoinLogic {
	return &UpdateCoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCoinLogic) UpdateCoin(req *types.UpdateCoinReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	symbol := l.svcCtx.Query.Symbol
	c := &model.Coin{
		CoinName: req.CoinName,
		Prec:     req.Prec,
		ID:       req.ID,
	}
	if err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		if _, err := tx.WithContext(l.ctx).Coin.Updates(c); err != nil {
			return err
		}
		//更新交易对中的配置
		if _, err := tx.WithContext(l.ctx).Symbol.
			Where(symbol.BaseCoinID.Eq(req.ID)).
			UpdateColumnSimple(symbol.BaseCoinPrec.Value(req.Prec)); err != nil {
			return err
		}

		if _, err := tx.WithContext(l.ctx).Symbol.
			Where(symbol.QuoteCoinID.Eq(req.ID)).
			UpdateColumnSimple(symbol.QuoteCoinPrec.Value(req.Prec)); err != nil {
			return err
		}

		symbols, err := tx.Symbol.WithContext(l.ctx).Where(symbol.BaseCoinID.Eq(req.ID)).Or(symbol.QuoteCoinID.Eq(req.ID)).Find()
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

		coinInfo := define.CoinInfo{
			CoinID:   int32(req.ID),
			CoinName: req.CoinName,
			Prec:     req.Prec,
		}
		data, err := yaml.Marshal(coinInfo)
		if err != nil {
			return err
		}
		if _, err := l.svcCtx.EtcdCli.Put(l.ctx, define.EtcdCoinPrefix+req.CoinName, string(data)); err != nil {
			logx.Errorw("put config to etcd failed", logx.Field("err", err))
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}
