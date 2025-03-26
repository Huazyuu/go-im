package logic

import (
	"context"
	"errors"
	"fmt"
	"server/utils/jwt"
	"server/utils/pwd"
	"server/yu_auth/auth_models"

	"server/yu_auth/auth_api/internal/svc"
	"server/yu_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 查找用户
	var user auth_models.UserModel
	err = l.svcCtx.DB.Where("id = ?", req.Username).Take(&user).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	// 检查密码
	if !pwd.VerifyPwd(user.UserPwd, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}
	// gen token
	token, err := jwt.GenerateToken(jwt.JwtPayload{
		UserID:   user.ID,
		Nickname: user.UserNickname,
		Role:     user.UserRole,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("系统服务错误")
	}
	return &types.LoginResponse{Token: token}, nil
}
