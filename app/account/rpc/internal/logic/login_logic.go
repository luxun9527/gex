package logic

import (
	"context"
	"errors"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/utils"
	"gorm.io/gorm"

	"github.com/luxun9527/gex/app/account/rpc/internal/svc"
	"github.com/luxun9527/gex/app/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	user := l.svcCtx.Query.User

	result, err := user.WithContext(l.ctx).
		Where(user.Username.Eq(in.Username)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.LoginFailed
		}
		logx.Errorw("find user failed", logger.ErrorField(err))
		return nil, errs.Internal
	}
	if !utils.BcryptCheck(in.Password, result.Password) {
		return nil, errs.LoginFailed
	}
	claims := l.svcCtx.JwtClient.CreateClaims(utils.JwtContent{
		UserID:   int64(result.ID),
		Username: result.Username,
		NickName: "",
	})
	token, err := l.svcCtx.JwtClient.CreateToken(claims)
	if err != nil {
		logx.Errorw("create token failed", logger.ErrorField(err))
		return nil, errs.Internal
	}

	return &pb.LoginResp{
		Uid:        int64(result.ID),
		Username:   in.Username,
		Token:      token,
		ExpireTime: claims.ExpiresAt.Unix(),
	}, nil

}
