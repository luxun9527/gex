package logic

import (
	"context"
	"github.com/luxun9527/gex/app/order/rpc/internal/svc"
	"github.com/luxun9527/gex/app/order/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	enum "github.com/luxun9527/gex/common/proto/enum"
	"github.com/luxun9527/gex/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderList 获取用户订单列表
func (l *GetOrderListLogic) GetOrderList(in *pb.GetOrderListByUserReq) (*pb.GetOrderListByUserResp, error) {
	entrustOrder := l.svcCtx.Query.EntrustOrder
	statusCond := make([]int32, 0, len(in.StatusList))
	for _, v := range in.StatusList {
		statusCond = append(statusCond, int32(v))
	}
	result, err := entrustOrder.WithContext(l.ctx).
		Omit(entrustOrder.UpdatedAt).
		Where(entrustOrder.UserID.Eq(in.UserId)).
		Where(entrustOrder.Status.In(statusCond...)).
		Order(entrustOrder.ID.Desc()).
		Find()
	if err != nil {
		logx.Errorw("GetOrderList query user order list failed", logger.Error(err))
		return nil, errs.ExecSqlFailed
	}
	orders := make([]*pb.Order, 0, len(result))

	for _, v := range result {
		order := &pb.Order{
			Id:             v.ID,
			OrderId:        v.OrderID,
			UserId:         v.UserID,
			SymbolId:       v.SymbolID,
			SymbolName:     v.SymbolName,
			Qty:            utils.PrecCut(v.Qty, l.svcCtx.Config.SymbolInfo.BaseCoinPrec),
			Price:          utils.PrecCut(v.Price, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Amount:         utils.PrecCut(v.Amount, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Side:           enum.Side(v.Side),
			Status:         enum.OrderStatus(v.Status),
			OrderType:      enum.OrderType(v.OrderType),
			FilledQty:      utils.PrecCut(v.FilledQty, l.svcCtx.Config.SymbolInfo.BaseCoinPrec),
			FilledAmount:   utils.PrecCut(v.FilledAmount, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			FilledAvgPrice: utils.PrecCut(v.FilledAvgPrice, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			CreatedAt:      v.CreatedAt,
		}
		orders = append(orders, order)
	}
	return &pb.GetOrderListByUserResp{OrderList: orders}, nil
}
