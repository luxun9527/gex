package logic

import (
	"context"
	matchpb "github.com/luxun9527/gex/app/match/rpc/pb"
	"github.com/luxun9527/gex/app/quotes/api/internal/svc"
	"github.com/luxun9527/gex/app/quotes/api/internal/types"
	"github.com/luxun9527/gex/common/errs"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepthListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepthListLogic {
	return &GetDepthListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepthListLogic) GetDepthList(req *types.GetDepthListReq) (resp *types.GetDepthListResp, err error) {
	// todo: add your logic here and delete this line
	conn, ok := l.svcCtx.MatchClients.GetConn(req.Symbol)
	if !ok {
		logx.Sloww("symbol not found", logx.Field("symbol", req.Symbol))
		return nil, errs.Internal
	}
	client := l.svcCtx.GetMatchClient(conn)
	depthResp, err := client.GetDepth(l.ctx, &matchpb.GetDepthReq{
		Symbol: req.Symbol,
		Level:  req.Level,
	})
	if err != nil {
		return nil, err
	}
	asks := make([]*types.Position, 0, len(depthResp.Asks))
	bids := make([]*types.Position, 0, len(depthResp.Bids))
	for _, v := range depthResp.Asks {
		position := &types.Position{
			Qty:    v.Qty,
			Price:  v.Price,
			Amount: v.Amount,
		}
		asks = append(asks, position)
	}
	for _, v := range depthResp.Bids {
		position := &types.Position{
			Qty:    v.Qty,
			Price:  v.Price,
			Amount: v.Amount,
		}
		bids = append(bids, position)
	}
	resp = &types.GetDepthListResp{
		Version: cast.ToString(depthResp.Version),
		Asks:    asks,
		Bids:    bids,
	}
	return
}
