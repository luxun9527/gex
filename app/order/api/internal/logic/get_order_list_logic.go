package logic

import (
	"context"
	orderpb "github.com/luxun9527/gex/app/order/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	enum "github.com/luxun9527/gex/common/proto/enum"
	"github.com/spf13/cast"

	"github.com/luxun9527/gex/app/order/api/internal/svc"
	"github.com/luxun9527/gex/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListReq) (resp *types.GetOrderListResp, err error) {
	// todo: add your logic here and delete this line
	uid := l.ctx.Value("uid")
	conn, ok := l.svcCtx.OrderClients.GetConn(req.SymbolName)
	if !ok {
		logx.Sloww("symbol not found", logx.Field("symbol", req.SymbolName))
		return nil, errs.Internal
	}
	client := l.svcCtx.GetOrderClient(conn)
	statusList := make([]enum.OrderStatus, 0, len(req.Status))
	for _, v := range req.Status {
		statusList = append(statusList, enum.OrderStatus(v))
	}
	orderList, err := client.GetOrderList(l.ctx, &orderpb.GetOrderListByUserReq{
		UserId:     cast.ToInt64(uid),
		StatusList: statusList,
	})
	if err != nil {
		return nil, err
	}
	orderInfoList := make([]*types.OrderInfo, 0, len(orderList.OrderList))
	for _, v := range orderList.OrderList {
		orderInfo := &types.OrderInfo{
			Id:             cast.ToString(v.Id),
			OrderId:        v.OrderId,
			UserId:         v.UserId,
			SymbolName:     v.SymbolName,
			Price:          v.Price,
			Qty:            v.Qty,
			Amount:         v.Amount,
			Side:           int32(v.Side),
			Status:         int32(v.Status),
			OrderType:      int32(v.OrderType),
			FilledQty:      v.FilledQty,
			FilledAmount:   v.FilledAmount,
			FilledAvgPrice: v.FilledAvgPrice,
			CreatedAt:      v.CreatedAt,
		}
		orderInfoList = append(orderInfoList, orderInfo)
	}
	resp = &types.GetOrderListResp{OrderList: orderInfoList}
	return
}
