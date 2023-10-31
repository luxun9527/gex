package response

import (
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {

	lang := r.Header.Get("language")
	var body Body
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			logx.Errorw("unknown error", logger.ErrorField(err))
		}
		body.Code = int(e.Code())
		code := e.Code()
		if body.Code < int(errs.Start) {
			code = codes.Code(errs.DefaultCode)
		}
		body.Msg = errs.Code(code).Translate(lang)

		if e.Message() != "" && int(e.Code()) > int(errs.Start) {
			body.Msg += ":" + e.Message()
		}

		body.Data = struct{}{}
	} else {
		body.Msg = "OK"
		if resp == nil {
			resp = struct{}{}
		}
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
