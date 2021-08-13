package handler

import (
	"net/http"

	"hothothot/service/gateway/internal/logic/password"
	"hothothot/service/gateway/internal/svc"
	"hothothot/service/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ResetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PwdResetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewResetLogic(r.Context(), ctx)
		resp, err := l.Reset(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
