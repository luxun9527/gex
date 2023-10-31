package logic

import (
	"context"
	"github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/svc"
	"github.com/luxun9527/gex/app/quotes/api/internal/types"
	"github.com/luxun9527/gex/common/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTickerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTickerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTickerListLogic {
	return &GetTickerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTickerListLogic) GetTickerList(req *types.GetTickerListReq) (resp *types.GetTickerListResp, err error) {
	conn, ok := l.svcCtx.MatchClients.GetConn(req.Symbol)
	if !ok {
		logx.Sloww("symbol not found", logx.Field("symbol", req.Symbol))
		return nil, errs.Internal
	}
	client := l.svcCtx.GetMatchClient(conn)
	tickerResp, err := client.GetTicker(l.ctx, &pb.GetTickerReq{Symbol: req.Symbol})
	if err != nil {
		return nil, err
	}
	tickerList := make([]*types.Ticker, 0, len(tickerResp.TickerList))
	for _, v := range tickerResp.TickerList {
		ticker := &types.Ticker{
			LastPrice:   v.LatestPrice,
			High:        v.High,
			Low:         v.Low,
			Volume:      v.Volume,
			Amount:      v.Amount,
			PriceRange:  v.PriceRange,
			Symbol:      v.Symbol,
			Last24Price: v.Last24Price,
		}
		tickerList = append(tickerList, ticker)
	}

	resp = &types.GetTickerListResp{}
	resp.TickerList = tickerList

	return
}
