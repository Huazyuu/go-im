package logic

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/zeromicro/go-zero/core/logx"
	"server/yu_auth/auth_api/internal/svc"
)

type Login_giteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogin_giteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Login_giteeLogic {
	return &Login_giteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Login_giteeLogic) Login_gitee() (resp string, err error) {
	conf := l.svcCtx.Config.Gitee
	// 参数校验
	if conf.ClientID == "" || conf.Redirect == "" {
		return "", errors.New("gitee oauth configuration incomplete")
	}

	// 构建授权URL
	baseURL := "https://gitee.com/oauth/authorize"
	params := url.Values{}
	params.Add("client_id", conf.ClientID)
	params.Add("redirect_uri", conf.Redirect)
	params.Add("response_type", "code")

	authURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	return authURL, nil
}
