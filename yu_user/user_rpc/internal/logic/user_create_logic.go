package logic

import (
	"context"
	"errors"
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

	var user user_models.UserModel
	err := l.svcCtx.DB.Where("user_name=?", in.Username).Take(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
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
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}
	return &user_rpc.UserCreateResponse{
		Username: user.UserName,
		UserId:   int32(user.ID),
	}, nil
}
