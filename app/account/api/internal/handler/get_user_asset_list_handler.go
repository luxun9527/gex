package handler

import (
	"github.com/luxun9527/gex/app/account/api/internal/logic"
	"github.com/luxun9527/gex/app/account/api/internal/svc"
	"github.com/luxun9527/gex/common/pkg/response"
	"net/http"
)

func GetUserAssetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserAssetListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAssetList()
		response.Response(w, r, resp, err) //

	}
}
