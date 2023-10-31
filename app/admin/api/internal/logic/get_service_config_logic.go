package logic

import (
	"context"
	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"gopkg.in/yaml.v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceConfigLogic {
	return &GetServiceConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceConfigLogic) GetServiceConfig(req *types.GetServiceConfigReq) (resp *types.GetServiceConfigResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.EtcdCli.Get(l.ctx, req.Key)
	if err != nil {
		logx.Errorw("get config from etcd failed", logx.Field("err", err))
		return nil, errs.Internal
	}
	//验证

	kvs := data.Kvs
	if len(kvs) > 0 {
		resp = &types.GetServiceConfigResp{ConfigData: string(kvs[0].Value)}
		m := make(map[string]interface{})
		if err := yaml.Unmarshal(kvs[0].Value, m); err != nil {
			logx.Errorw("unmashal data failed", logger.ErrorField(err))
		}
		logx.Infow("detail", logx.Field("detail", m))
	} else {
		resp = &types.GetServiceConfigResp{ConfigData: ""}

	}
	return
}
