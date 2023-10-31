package bootstrap

import (
	"context"
	"github.com/luxun9527/gex/app/match/rpc/internal/consumer"
	"github.com/luxun9527/gex/app/match/rpc/internal/engine"
	"github.com/luxun9527/gex/app/match/rpc/internal/svc"
	"github.com/luxun9527/gex/app/order/rpc/orderservice"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/proto/enum"
	"github.com/luxun9527/gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

func Start(sc *svc.ServiceContext) {
	loadOrder(sc)
	consumer.InitMatchConsumer(sc)
}
func loadOrder(sc *svc.ServiceContext) {
	stream, err := sc.OrderClient.GetOrderAllPendingOrder(context.Background(), &orderservice.OrderEmpty{})
	if err != nil {
		logx.Severef("call GetOrderAllPendingOrder failed", logger.ErrorField(err))
	}
	for {
		order, err := stream.Recv()
		if order.Done {
			break
		}
		if err != nil {
			//没有加载完则panic
			logx.Severef("read order from order service failed err = %v", err)
		}
		logx.Debugw("init load order", logx.Field("order", order))
		o := &engine.Order{
			Uid:            order.Uid,
			OrderID:        order.OrderId,
			SequenceId:     order.SequenceId,
			CreateTime:     0,
			IsCancel:       false,
			Price:          utils.NewFromStringMaxPrec(order.Price),
			Qty:            utils.NewFromStringMaxPrec(order.Qty),
			OrderType:      order.OrderType,
			Amount:         utils.NewFromStringMaxPrec(order.Amount),
			Side:           order.Side,
			OrderStatus:    enum.OrderStatus_NewCreated,
			UnfilledQty:    utils.NewFromStringMaxPrec(order.UnFilledQty),
			FilledQty:      utils.DecimalZeroMaxPrec,
			UnfilledAmount: utils.NewFromStringMaxPrec(order.UnFilledAmount),
		}

		sc.MatchEngine.HandleOrder(o)

	}
}
