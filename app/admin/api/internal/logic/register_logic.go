package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/model"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/utils"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

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
	if req.ConfirmPassword != req.Password {
		return nil, errs.PasswordNotMatch
	}

	u := l.svcCtx.Query.User
	count, err := u.WithContext(l.ctx).Where(u.Username.Eq(req.Username)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errs.DuplicateDataErr
	}
	p := utils.BcryptHash(req.Password)
	user := &model.User{
		ID:        0,
		Username:  req.Username,
		Password:  p,
		CreatedAt: 0,
		UpdatedAt: 0,
		DeletedAt: 0,
	}
	if err := u.WithContext(l.ctx).Create(user); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
