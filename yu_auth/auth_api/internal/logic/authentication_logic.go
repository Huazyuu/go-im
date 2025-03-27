package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"server/utils/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"server/yu_auth/auth_api/internal/svc"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp string, err error) {
	if token == "" {
		return "", errors.New("请传入token")
	}

	_, err = jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return "", errors.New("认证失败")
	}

	// 是否在注销redis
	key := fmt.Sprintf("logout_%s", token)
	_, err = l.svcCtx.Redis.Get(key).Result()
	if !errors.Is(err, redis.Nil) {
		return "", errors.New("认证失败,认证过期")
	}
	resp = "认证通过"
	return resp, nil
}
