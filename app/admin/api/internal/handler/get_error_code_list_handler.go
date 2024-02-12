package handler

import (
	"github.com/luxun9527/gex/app/admin/api/internal/logic"
	"github.com/luxun9527/gex/app/admin/api/internal/svc"
	"github.com/luxun9527/gex/app/admin/api/internal/types"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetErrorCodeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetErrorCodeListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, errs.WarpMessage(errs.ParamValidateFailed, err.Error()))
			return
		}

		l := logic.NewGetErrorCodeListLogic(r.Context(), svcCtx)
		resp, err := l.GetErrorCodeList(&req)
		response.Response(w, r, resp, err)

	}
}
