package logic

import (
	"context"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMatchListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMatchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMatchListLogic {
	return &GetMatchListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMatchListLogic) GetMatchList(req *types.GetMatchListReq) (resp *types.GetMatchListResp, err error) {

	return
}
