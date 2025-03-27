package handler

import (
	"net/http"
	"server/common/response"
	"server/yu_auth/auth_api/internal/logic"
	"server/yu_auth/auth_api/internal/svc"
)

func login_giteeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogin_giteeLogic(r.Context(), svcCtx)
		redirectURL, err := l.Login_gitee()
		if err != nil {
			// 错误处理
			response.Response(r, w, nil, err)
			return
		}
		// 执行重定向
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
