package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/yu_auth/auth_api/internal/logic"
	"server/yu_auth/auth_api/internal/svc"
	"server/yu_auth/auth_api/internal/types"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticationRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication(&req)
		response.Response(r, w, resp, err)
	}
}
