package logic

import (
	"context"
	"github.com/luxun9527/gex/app/account/rpc/internal/dao/model"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/utils"

	"github.com/luxun9527/gex/app/account/rpc/internal/svc"
	"github.com/luxun9527/gex/app/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	user := &model.User{
		Username:    in.Username,
		Password:    utils.BcryptHash(in.Password),
		PhoneNumber: in.PhoneNumber,
		Status:      1,
	}
	if err := l.svcCtx.Query.User.WithContext(l.ctx).Create(user); err != nil {
		logx.Errorw("create user failed", logger.ErrorField(err))
		return nil, errs.ExecSqlFailed
	}
	return &pb.Empty{}, nil
}
