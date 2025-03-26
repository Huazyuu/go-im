package handler

import (
	"net/http"
	"server/common/response"
	"server/yu_auth/auth_api/internal/logic"
	"server/yu_auth/auth_api/internal/svc"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		response.Response(r, w, resp, err)

	}
}
