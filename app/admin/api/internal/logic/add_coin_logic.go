package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/model"
	"github.com/luxun9527/gex/common/errs"

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
	if err := coin.WithContext(l.ctx).Create(c); err != nil {
		logx.Errorw("create coin failed", logx.Field("err", err))
		return nil, err
	}

	return
}
