package logic

import (
	"context"
	"github.com/luxun9527/gex/app/account/rpc/accountservice"
	"github.com/luxun9527/gex/common/pkg/logger"

	"github.com/luxun9527/gex/app/account/api/internal/svc"
	"github.com/luxun9527/gex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.AccountRpcClient.Register(l.ctx, &accountservice.RegisterReq{
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}); err != nil {
		logx.Errorw("register failed", logger.ErrorField(err))
		return nil, err
	}
	return
}
