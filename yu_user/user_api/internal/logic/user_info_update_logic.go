package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"server/common/models/commonType/cverify"
	"server/utils/umap"
	"server/yu_user/user_api/internal/svc"
	"server/yu_user/user_api/internal/types"
	"server/yu_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {
	// 开启事务
	tx := l.svcCtx.DB.Begin()
	if tx.Error != nil {
		return nil, errors.New("开启事务失败")
	}

	userMaps := umap.RefStructByTag(*req, "user")
	if len(userMaps) != 0 {
		var user user_models.UserModel
		err = tx.Where(req.UserID).Take(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("用户不存在")
		} else if err != nil {
			tx.Rollback()
			return nil, errors.New("系统错误")
		}

		err = tx.Model(&user).Updates(userMaps).Error
		if err != nil {
			logx.Error(err)
			logx.Info(userMaps)
			tx.Rollback()
			return nil, errors.New("系统错误 用户信息更新失败")
		}
	}
	userConfMaps := umap.RefStructByTag(*req, "user_conf")
	if len(userConfMaps) != 0 {
		var userconf user_models.UserConfModel
		err = tx.Where("user_id=?", req.UserID).Take(&userconf).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New("用户配置不存在")
		} else if err != nil {
			tx.Rollback()
			return nil, errors.New("系统错误")
		}

		verificationQuestions, ok := userConfMaps["verification_question"]
		if ok {
			delete(userConfMaps, "verification_question")
			data := cverify.VerificationQuestion{}
			umap.MapToStruct(verificationQuestions.(map[string]any), &data)
			tx.Model(&userconf).Updates(&user_models.UserConfModel{
				VerificationQuestion: &data,
			})
		}

		err = tx.Model(&userconf).Updates(userConfMaps).Error
		if err != nil {
			tx.Rollback()
			logx.Error(userConfMaps)
			logx.Error(err)
			return nil, errors.New("用户信息更新失败")
		}
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("提交事务失败")
	}
	return
}
