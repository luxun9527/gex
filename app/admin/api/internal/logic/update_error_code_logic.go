package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/model"
	"github.com/luxun9527/gex/app/admin/api/internal/dao/query"
	"github.com/luxun9527/gex/common/errs"
	"gopkg.in/yaml.v3"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateErrorCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateErrorCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateErrorCodeLogic {
	return &UpdateErrorCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateErrorCodeLogic) UpdateErrorCode(req *types.UpdateErrorCodeReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	code := &model.ErrorCode{
		ID:            int32(req.Id),
		ErrorCodeID:   req.ErrorCodeId,
		ErrorCodeName: req.ErrorCodeName,
		Language:      req.Language,
	}
	errorCode := l.svcCtx.Query.ErrorCode
	if err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		if _, err := tx.WithContext(l.ctx).ErrorCode.Updates(code); err != nil {
			return err
		}

		codes, err := tx.WithContext(l.ctx).ErrorCode.Where(errorCode.Language.Eq(req.Language)).Find()
		if err != nil {
			return err
		}
		m := make(map[int32]string)
		for _, v := range codes {
			m[v.ErrorCodeID] = v.ErrorCodeName
		}
		d, err := yaml.Marshal(m)
		if err != nil {
			return err
		}
		if _, err := l.svcCtx.EtcdCli.Put(l.ctx, errs.EtcdPrefixKey+req.Language, string(d)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return
}
