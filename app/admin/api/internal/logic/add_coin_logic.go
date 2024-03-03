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

type AddCoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCoinLogic {
	return &AddCoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCoinLogic) AddCoin(req *types.AddCoinReq) (resp *types.AddCoinResp, err error) {

	coin := l.svcCtx.Query.Coin
	c := &model.Coin{
		CoinName: req.CoinName,
		Prec:     req.Prec,
	}
	count, err := coin.WithContext(l.ctx).Where(coin.CoinName.Eq(req.CoinName)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errs.DuplicateDataErr
	}

	if err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		if err := tx.Coin.WithContext(l.ctx).Create(c); err != nil {
			logx.Errorw("create coin failed", logx.Field("err", err))
			return err
		}
		coins, err := tx.Coin.WithContext(l.ctx).Find()
		if err != nil {
			return err
		}
		for _, v := range coins {
			coinInfo := &define.CoinInfo{
				CoinID:   int32(v.ID),
				CoinName: v.CoinName,
				Prec:     v.Prec,
			}
			data, err := yaml.Marshal(coinInfo)
			if err != nil {
				return err
			}
			if _, err := l.svcCtx.EtcdCli.Put(l.ctx, define.EtcdCoinPrefix+v.CoinName, string(data)); err != nil {
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
