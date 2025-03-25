package models

import "server/common/models"

// FriendsModel 好友表
type FriendsModel struct {
	models.Model
	SendUserID uint   `json:"sendUserID"`             // 发起验证方
	RecvUserID uint   `json:"recvUserID"`             // 接受验证方
	Notice     string `gorm:"size:128" json:"notice"` // 备注
}
