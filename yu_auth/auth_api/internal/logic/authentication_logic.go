package logic

import (
	"context"
	"errors"
	"fmt"
	"server/utils/jwt"
	"server/utils/ulist"

	"server/yu_auth/auth_api/internal/svc"
	"server/yu_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp *types.AuthenticationReponse, err error) {
	if ulist.InList(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logx.Infof("%s 在白名单中", req.ValidPath)
		return
	}

	if req.Token == "" {
		logx.Error("token为空")
		err = errors.New("认证失败")
		return
	}

	claims, err := jwt.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Error(err.Error())
		err = errors.New("认证失败")
		return
	}

	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%s", req.Token)).Result()
	if err == nil {
		logx.Error("在黑名单中")
		err = errors.New("认证失败")
		return
	}

	res := &types.AuthenticationReponse{
		UserID: claims.UserID,
		Role:   int(claims.Role),
	}
	return res, nil
}
