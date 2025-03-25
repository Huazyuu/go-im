package models

import "server/common/models"

// FriendsModel 好友表
type FriendsModel struct {
	models.Model

	SendUserID uint      `json:"sendUserID"` // 发起验证方
	SendUser   UserModel `json:"-" gorm:"foreignKey:SendUserID"`
	RecvUserID uint      `json:"recvUserID"` // 接受验证方
	RecvUser   UserModel `json:"-" gorm:"foreignKey:RecvUserID"`

	Notice string `gorm:"size:128" json:"notice"` // 备注
}

func (FriendsModel) TableName() string {
	return "friends"
}
