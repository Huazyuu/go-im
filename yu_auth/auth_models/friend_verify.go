package auth_models

import (
	"server/common/models"
	"server/common/models/commonType/cverify"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                          `json:"sendUserID"`                    // 发起验证方
	RecvUserID           uint                          `json:"recvUserID"`                    // 接受验证方
	Status               int8                          `json:"status"`                        // 0未操作 1同意 2拒绝 3忽略
	AdditionalMsg        string                        `gorm:"size:128" json:"additionalMsg"` // 附加消息
	VerificationQuestion *cverify.VerificationQuestion `json:"verificationQuestion"`          // 验证问题
}
