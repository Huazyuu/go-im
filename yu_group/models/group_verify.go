package models

import (
	"server/common/models"
	"server/common/models/commonType/cverify"
)

// GroupVerifyModel 群验证表
type GroupVerifyModel struct {
	models.Model

	GroupID    uint       `json:"groupID"` // 群
	GroupModel GroupModel `json:"-" gorm:"foreignKey:GroupID"`

	UserID               uint                          `json:"userID"`                            // 需要加群或者是退群的用户id
	Status               int8                          `json:"status"`                            // 状态 0 未操作 1 同意 2 拒绝 3 忽略
	Type                 int8                          `json:"type"`                              // 类型 1 加群  2 退群
	AdditionalMessages   string                        `gorm:"size:32" json:"additionalMessages"` // 附加消息
	VerificationQuestion *cverify.VerificationQuestion `json:"verificationQuestion"`              // 验证问题  为3和4的时候需要
}

func (GroupVerifyModel) TableName() string {
	return "group_verify"
}
