package handler

import (
	"net/http"

	"mall/service/pay/api/internal/logic"
	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CallbackHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), ctx)
		resp, err := l.Callback(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}