package logic

import (
	"context"
	"database/sql"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/luxun9527/gex/app/order/rpc/internal/dao/model"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	enum "github.com/luxun9527/gex/common/proto/enum"
	matchMq "github.com/luxun9527/gex/common/proto/mq/match"
	commonWs "github.com/luxun9527/gex/common/proto/ws"
	gpush "github.com/luxun9527/gpush/proto"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/utils"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"time"

	"github.com/luxun9527/gex/app/order/rpc/internal/svc"
	"github.com/luxun9527/gex/app/order/rpc/pb"
	commonUtils "github.com/luxun9527/gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建订单
func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.OrderEmpty, error) {
	if in.OrderType == enum.OrderType_MO {
		if in.Side == enum.Side_Buy {
			in.Price = "0"
			in.Qty = "0"
		} else {
			in.Price = "0"
			in.Amount = "0"
		}
	}

	or := &model.EntrustOrder{
		ID:             l.svcCtx.SnowflakeGenerator.GetId(),
		OrderID:        utils.NewUuid(),
		UserID:         in.UserId,
		SymbolID:       in.SymbolId,
		SymbolName:     in.SymbolName,
		Qty:            in.Qty,
		Price:          in.Price,
		Side:           int32(in.Side),
		Amount:         in.Amount,
		Status:         int32(enum.OrderStatus_NewCreated),
		OrderType:      int32(in.OrderType),
		FilledQty:      "0",
		UnFilledQty:    in.Qty,
		FilledAvgPrice: "0",
		FilledAmount:   "0",
		UnFilledAmount: in.Amount,
		CreatedAt:      time.Now().Unix(),
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Errorw("CreateOrder BarrierFromGrpc db failed", logger.ErrorField(err))
		return nil, errs.CastToDtmError(errs.DTMFailed)
	}
	entrustOrder := l.svcCtx.Query.EntrustOrder
	db, err := entrustOrder.WithContext(l.ctx).UnderlyingDB().DB()
	if err != nil {
		logx.Errorw("CreateOrder get UnderlyingDB failed", logger.ErrorField(err))
		return nil, errs.CastToDtmError(errs.ExecSqlFailed)
	}
	toSQL := entrustOrder.WithContext(l.ctx).UnderlyingDB().ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(or)
	})
	//测试执行失败的情况

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if _, err := tx.Exec(toSQL); err != nil {
			return err
		}

		return nil
	}); err != nil {
		logx.Errorw("CreateOrder CallWithDB failed", logger.ErrorField(err))
		return nil, errs.CastToDtmError(errs.Internal)

	}
	//构建消息发送
	msg := &matchMq.MatchReq{Operate: &matchMq.MatchReq_NewOrder{
		NewOrder: &matchMq.NewOrderOperate{
			OrderId:    or.OrderID,
			SequenceId: or.ID,
			Uid:        or.UserID,
			Side:       in.Side,
			Price:      in.Price,
			Qty:        in.Qty,
			Amount:     in.Amount,
			OrderType:  in.OrderType,
		},
	}}
	data, _ := proto.Marshal(msg)
	if _, err := l.svcCtx.MatchProducer.Send(l.ctx, &pulsar.ProducerMessage{
		Payload: data,
	}); err != nil {
		return nil, errs.CastToDtmError(errs.PulsarFailed)
	}
	//发送ws数据
	wsOrder := &commonWs.Order{
		Id:             cast.ToString(or.ID),
		OrderId:        cast.ToString(or.OrderID),
		SymbolName:     or.SymbolName,
		Price:          commonUtils.PrecCut(or.Price, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
		Qty:            commonUtils.PrecCut(or.Qty, l.svcCtx.Config.SymbolInfo.BaseCoinPrec),
		Amount:         commonUtils.PrecCut(or.Amount, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
		Side:           int8(or.Side),
		Status:         int8(enum.OrderStatus_NewCreated),
		OrderType:      int8(or.OrderType),
		FilledAmount:   commonUtils.PrecCut(or.FilledAmount, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
		FilledQty:      commonUtils.PrecCut(or.FilledQty, l.svcCtx.Config.SymbolInfo.BaseCoinPrec),
		FilledAvgPrice: commonUtils.PrecCut(or.FilledAvgPrice, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
		Uid:            cast.ToString(or.UserID),
		CreatedAt:      or.CreatedAt,
	}
	l.pushWsData(wsOrder)
	return &pb.OrderEmpty{}, nil
}
func (l *CreateOrderLogic) pushWsData(order *commonWs.Order) {
	msg := commonWs.Message[commonWs.Order]{
		Topic:   commonWs.OrderPrefix.WithParam(l.svcCtx.Config.SymbolInfo.SymbolName),
		Payload: *order,
	}
	_, err := l.svcCtx.WsClient.PushData(context.Background(), &gpush.Data{
		Uid:   order.Uid,
		Topic: msg.Topic,
		Data:  msg.ToBytes(),
	})
	if err != nil {
		logx.Errorw("push ws data failed", logger.ErrorField(err))
	}
}
