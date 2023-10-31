package logic

import (
	"context"
	"github.com/luxun9527/gex/app/quotes/api/internal/svc"
	"github.com/luxun9527/gex/app/quotes/api/internal/types"
	klinepb "github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetKlineListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKlineListLogic {
	return &GetKlineListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKlineListLogic) GetKlineList(req *types.KlineListReq) (resp *types.KlineListResp, err error) {
	conn, ok := l.svcCtx.KlineClients.GetConn(req.Symbol)
	if !ok {
		logx.Sloww("symbol not found", logx.Field("symbol", req.Symbol))
		return nil, errs.Internal
	}
	client := l.svcCtx.GetKlineClient(conn)
	klineResp, err := client.GetKline(l.ctx, &klinepb.GetKlineReq{
		StartTime: req.StartTime,
		EntTime:   req.EndTime,
		KlineType: klinepb.KlineType(req.KlineType),
		Symbol:    req.Symbol,
	})
	if err != nil {
		return nil, err
	}
	klines := make([]*types.Kline, 0, len(klineResp.KlineList))
	for _, v := range klineResp.KlineList {
		k := &types.Kline{
			Open:       v.Open,
			High:       v.High,
			Low:        v.Low,
			Close:      v.Close,
			Volume:     v.Volume,
			Amount:     v.Amount,
			StartTime:  v.StartTime,
			EndTime:    v.EndTime,
			PriceRange: v.Range,
			Symbol:     v.Symbol,
		}
		klines = append(klines, k)
	}
	resp = &types.KlineListResp{}
	resp.KineList = klines
	return
}
