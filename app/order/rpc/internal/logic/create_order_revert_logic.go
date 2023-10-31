package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/luxun9527/gex/app/order/rpc/internal/dao/model"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/proto/enum"
	"gorm.io/gorm"

	"github.com/luxun9527/gex/app/order/rpc/internal/svc"
	"github.com/luxun9527/gex/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRevertLogic {
	return &CreateOrderRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 下单补偿
func (l *CreateOrderRevertLogic) CreateOrderRevert(in *pb.CreateOrderReq) (*pb.OrderEmpty, error) {
	logx.Infow("CreateOrderRevert invoke")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Errorw("BarrierFromGrpc db failed", logger.ErrorField(err))
		return nil, errs.DTMFailed
	}
	entrustOrder := l.svcCtx.Query.EntrustOrder
	db, err := entrustOrder.WithContext(l.ctx).UnderlyingDB().DB()
	if err != nil {
		logx.Errorw("get UnderlyingDB failed", logger.ErrorField(err))
		return nil, errs.ExecSqlFailed
	}
	//修改订单的状态为无效
	order := &model.EntrustOrder{
		Status: int32(enum.OrderStatus_Wasted),
	}
	//构建sql
	toSQL := entrustOrder.WithContext(l.ctx).UnderlyingDB().ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Select(entrustOrder.Status.ColumnName().String()).
			Clauses(entrustOrder.OrderID.Eq(in.OrderId)).
			Updates(order)
	})

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if _, err := tx.Exec(toSQL); err != nil {
			return err
		}
		logx.Infow("CreateOrderRevert CallWithDB invoke")
		return nil
	}); err != nil {
		logx.Errorw("CallWithDB failed", logger.ErrorField(err))
		return nil, errs.Internal

	}
	return &pb.OrderEmpty{}, nil
}
