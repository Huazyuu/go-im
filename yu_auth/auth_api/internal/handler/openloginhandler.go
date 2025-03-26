package handler

import (
	"net/http"
	"server/common/response"
	"server/yu_auth/auth_api/internal/logic"
	"server/yu_auth/auth_api/internal/svc"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login()
		response.Response(r, w, resp, err)

	}
}
