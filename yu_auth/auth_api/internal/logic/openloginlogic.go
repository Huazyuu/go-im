package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"server/utils/open_login"
	"server/yu_auth/auth_api/internal/svc"
	"server/yu_auth/auth_api/internal/types"
	"server/yu_auth/auth_models"
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
	type OpenInfo struct {
		Username string
		OpenID   string
		Avatar   string
	}

	var info OpenInfo
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

	info = OpenInfo{
		OpenID:   strconv.Itoa(gitInfo.ID), // 转换数字ID为字符串
		Username: username,
		Avatar:   gitInfo.AvatarURL,
	}

	// 查询用户信息
	var user auth_models.UserModel

	err = l.svcCtx.DB.Where("user_name = ?", info.Username).First(&user).Error
	if err != nil {
		// 自动注册逻辑
		logx.Info("注册 rpc 逻辑")
		/*createRes, err := l.svcCtx.DB.UserCreate(l.ctx, &user_rpc.UserCreateRequest{
			UserNickName:       info.Username,
			UserPassword:       "", // 第三方登录不需要密码
			UserRole:           2,  // 普通用户
			UserAvatar:         info.Avatar,
			UserOpenId:         info.OpenID,
			UserRegisterSource: strings.ToLower(req.Flag), // 注册来源
		})
		if err != nil {
			logx.Errorf("用户创建失败: %v", err)
			return nil, errors.New("自动注册失败")
		}

		// 填充新建用户信息
		user = auth_models.UserModel{
			Model:    gorm.Model{ID: uint(createRes.UserId)},
			Nickname: info.Username,
			Role:     2,
			OpenID:   info.OpenID,
		}*/
	}

	// 生成JWT
	/*	token, err := jwt.GenerateToken(jwt.JwtPayload{
			UserID:   user.ID,
			Username: user.UserName,
			Role:     user.UserRole,
		}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
		if err != nil {
			logx.Errorf("JWT生成失败: %v", err)
			return nil, errors.New("系统错误")
		}*/

	/*	return &types.LoginResponse{
		Token: token,
		UserInfo: types.UserInfo{
			UserId:   int64(user.ID),
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	}, nil*/
	return
}
