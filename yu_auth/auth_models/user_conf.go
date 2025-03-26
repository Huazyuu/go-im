package auth_models

import (
	"server/common/models"
	"server/common/models/commonType/cverify"
)

// UserConfModel 用户配置表
type UserConfModel struct {
	models.Model
	UserID    uint    `json:"userID"`
	RecallMsg *string `gorm:"size:32" json:"recallMsg"` // 撤回消息提示内容

	IsFriendOnlineNotify bool `json:"isFriendOnlineNotify"` // 好友上线提醒
	IsOnline             bool `json:"isOnline"`             // 是否在线
	IsSound              bool `json:"isSound"`              // 提醒声音
	IsSecureLink         bool `json:"isSecureLink"`         // 安全链接,不直接转跳
	IsSavePwd            bool `json:"isSavePwd"`            // 保存密码

	SearchUser int8 `json:"searchUser"` // 查到到我的方式 0不允许别人查找 1用户号查 2可以用昵称

	Verification         cverify.VerifyType            `json:"verification"`         // 好友验证 0不允许添加 1验证消息 2回答问题 4需要正确回答问题
	VerificationQuestion *cverify.VerificationQuestion `json:"verificationQuestion"` // 验证问题 3,4需要
}
