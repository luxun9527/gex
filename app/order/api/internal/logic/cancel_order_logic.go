package logic

import (
	"context"
	orderpb "github.com/luxun9527/gex/app/order/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/spf13/cast"

	"github.com/luxun9527/gex/app/order/api/internal/svc"
	"github.com/luxun9527/gex/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.CancelOrderReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	conn, ok := l.svcCtx.OrderClients.GetConn(req.SymbolName)
	if !ok {
		logx.Sloww("symbol not found", logx.Field("symbol", req.SymbolName))
		return nil, errs.Internal
	}
	client := l.svcCtx.GetOrderClient(conn)
	_, err = client.CancelOrder(l.ctx, &orderpb.CancelOrderReq{
		Id:  cast.ToInt64(req.ID),
		Uid: cast.ToInt64(uid),
	})
	if err != nil {
		return nil, err
	}
	return &types.Empty{}, nil

}
