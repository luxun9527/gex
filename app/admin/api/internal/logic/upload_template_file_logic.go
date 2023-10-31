package logic

import (
	"context"
	"mime/multipart"

	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadTemplateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadTemplateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadTemplateFileLogic {
	return &UploadTemplateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadTemplateFileLogic) UploadTemplateFile(req *types.UploadTemplateFileReq, file *multipart.File, header *multipart.FileHeader) (resp *types.Empty, err error) {

	//io.Copy(file)

	return
}
