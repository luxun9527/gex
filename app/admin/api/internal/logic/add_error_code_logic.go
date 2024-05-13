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

type AddErrorCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddErrorCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddErrorCodeLogic {
	return &AddErrorCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddErrorCodeLogic) AddErrorCode(req *types.AddErrorCodeReq) (resp *types.Empty, err error) {

	errorCode := l.svcCtx.Query.ErrorCode

	if err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		count, err := errorCode.WithContext(l.ctx).
			Where(errorCode.Language.Eq(req.Language), errorCode.ErrorCodeID.Eq(req.ErrorCodeId)).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return errs.DuplicateDataErr
		}
		code := &model.ErrorCode{
			ID:            0,
			ErrorCodeID:   req.ErrorCodeId,
			ErrorCodeName: req.ErrorCodeName,
			Language:      req.Language,
			CreatedAt:     0,
			UpdatedAt:     0,
			DeletedAt:     0,
		}

		if err := tx.WithContext(l.ctx).ErrorCode.Create(code); err != nil {
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
