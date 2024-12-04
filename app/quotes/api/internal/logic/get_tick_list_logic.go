package logic

import (
	"context"
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/svc"
	"github.com/luxun9527/gex/app/quotes/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTickListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTickListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTickListLogic {
	return &GetTickListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTickListLogic) GetTickList(req *types.GetTickReq) (resp *types.GetTickResp, err error) {

	tickListResp, err := l.svcCtx.MatchClients.GetTick(l.ctx, &matchpb.GetTickReq{
		Symbol: req.Symbol,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}
	tickerList := make([]*types.TickInfo, 0, len(tickListResp.TickList))
	for _, v := range tickListResp.TickList {
		ticker := &types.TickInfo{
			Price:        v.Price,
			Qty:          v.Qty,
			Amount:       v.Amount,
			Timestamp:    v.Timestamp / 1e9,
			Symbol:       v.Symbol,
			TakerIsBuyer: v.TakerIsBuyer,
		}
		tickerList = append(tickerList, ticker)
	}

	resp = &types.GetTickResp{
		TickList: tickerList,
	}

	return
}
