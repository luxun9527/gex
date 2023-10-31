package logic

import (
	"context"
	"github.com/luxun9527/gex/app/account/rpc/accountservice"
	"github.com/luxun9527/gex/common/pkg/logger"

	"github.com/luxun9527/gex/app/account/api/internal/svc"
	"github.com/luxun9527/gex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.AccountRpcClient.Login(l.ctx, &accountservice.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorw("call account login failed", logger.ErrorField(err))
		return nil, err
	}
	resp = &types.LoginResp{
		Uid:        loginResp.Uid,
		Username:   loginResp.Username,
		Token:      loginResp.Token,
		ExpireTime: loginResp.ExpireTime,
	}
	return
}
