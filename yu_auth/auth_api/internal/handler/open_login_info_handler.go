package handler

import (
	"net/http"
	"server/common/response"
	"server/yu_auth/auth_api/internal/logic"
	"server/yu_auth/auth_api/internal/svc"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		response.Response(r, w, resp, err)

	}
}
