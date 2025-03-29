package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"server/utils/jwt"
	"server/utils/open_login"
	"server/yu_auth/auth_api/internal/svc"
	"server/yu_auth/auth_api/internal/types"
	"server/yu_auth/auth_models"
	"server/yu_user/user_rpc/types/user_rpc"
	"strconv"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(code string) (resp *types.LoginResponse, err error) {

	gitInfo, openError := open_login.NewGiteeLogin(code, open_login.GiteeConf{
		ClientID:     l.svcCtx.Config.Gitee.ClientID,
		ClientSecret: l.svcCtx.Config.Gitee.ClientSecret,
		Redirect:     l.svcCtx.Config.Gitee.Redirect,
	})
	if openError != nil {
		logx.Errorf("Gitee登录失败: %v", openError)
		return nil, errors.New("gitee登录失败")
	}

	// 处理GitHub昵称 赋值login 唯一
	username := gitInfo.Login

	type OpenInfo struct {
		Username string
		OpenID   string
		Avatar   string
	}

	info := OpenInfo{
		OpenID:   strconv.Itoa(gitInfo.ID), // 转换数字ID为字符串
		Username: username,
		Avatar:   gitInfo.AvatarURL,
	}

	// 查询用户信息
	var user auth_models.UserModel

	err = l.svcCtx.DB.Where("user_name = ?", info.Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 自动注册逻辑
		logx.Info("注册 rpc 逻辑")
		res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
			Username:       info.Username,
			Nickname:       info.Username,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			OpenId:         info.OpenID,
			RegisterSource: "Gitee",
		})
		if err != nil {
			logx.Errorf(err.Error())
			return nil, errors.New("登陆失败 err:" + err.Error())
		}
		user.ID = uint(res.UserId)
		user.UserRole = 2
		user.UserName = info.Username
	}

	token, err := jwt.GenerateToken(jwt.JwtPayload{
		UserID:   user.ID,
		Username: user.UserName,
		Role:     user.UserRole,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Errorf(err.Error())
		return nil, errors.New("登陆失败 err:" + err.Error())
	}
	return &types.LoginResponse{Token: token}, nil
}
