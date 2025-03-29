package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/yu_user/user_models"

	"server/yu_user/user_rpc/internal/svc"
	"server/yu_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {

	// 开启事务
	tx := l.svcCtx.DB.Begin()
	if tx.Error != nil {
		return nil, errors.New("开启事务失败")
	}

	var user user_models.UserModel
	err := tx.Where("user_name=?", in.Username).Take(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, errors.New("该用户已存在")
	}

	user = user_models.UserModel{
		UserName:       in.Username,
		UserNickname:   in.Username,
		UserPwd:        in.Password,
		UserAvatar:     in.Avatar,
		UserRole:       int8(in.Role),
		OpenID:         in.OpenId,
		RegisterSource: in.RegisterSource,
	}

	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}

	fmt.Println(user)

	err = tx.Create(&user_models.UserConfModel{
		UserID:               user.ID,
		RecallMsg:            nil,
		IsFriendOnlineNotify: false,
		IsOnline:             true, // 用户在线
		IsSound:              false,
		IsSecureLink:         false,
		IsSavePwd:            false, // 不保存密码
		SearchUser:           2,     // id 昵称搜索
		Verification:         2,     // 需要验证
	}).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("创建用户失败")
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("提交事务失败")
	}
	return &user_rpc.UserCreateResponse{
		Username: user.UserName,
		UserId:   int32(user.ID),
	}, nil
}
