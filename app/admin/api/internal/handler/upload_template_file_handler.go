package handler

import (
	"github.com/luxun9527/gex/app/admin/api/internal/logic"
	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func UploadTemplateFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.MultipartForm == nil {
			if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
				logx.Errorw("parse multipart form failed", logx.Field("err", err))
				response.Response(w, r, nil, errs.ParamValidateFailed)
				return
			}
		}
		f, fh, err := r.FormFile("template")
		if err != nil {
			logx.Errorw("parse multipart form failed", logger.ErrorField(err))
			response.Response(w, r, nil, errs.ParamValidateFailed)
			return
		}
		serviceName := r.PostFormValue("service_name")
		symbol := r.PostFormValue("symbol")
		l := logic.NewUploadTemplateFileLogic(r.Context(), svcCtx)
		req := &types.UploadTemplateFileReq{
			ServiceName: serviceName,
			Symbol:      symbol,
		}
		resp, err := l.UploadTemplateFile(req, &f, fh)
		response.Response(w, r, resp, err) //

	}
}
