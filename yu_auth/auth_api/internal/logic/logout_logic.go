package logic

import (
	"context"
	"errors"
	"fmt"
	"server/utils/jwt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"server/yu_auth/auth_api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	if token == "" {
		return "", errors.New("请传入token")
	}

	payload, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return "", errors.New("token错误")
	}
	now := time.Now()
	expiration := payload.ExpiresAt.Sub(now)

	key := fmt.Sprintf("logout_%s", token)
	res := l.svcCtx.Redis.SetNX(key, "", expiration)
	if res.Err() != nil {
		return "", errors.New("系统错误: " + res.Err().Error())
	}

	resp = "注销成功"
	return
}
