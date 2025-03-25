package models

import "server/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	UserNickname string `gorm:"size:32" json:"userNickname"`
	UserPwd      string `gorm:"size:64" json:"userPwd"`
	UserAbstract string `gorm:"size:128" json:"userAbstract"`
	UserAvatar   string `gorm:"size:256" json:"userAvatar"`
	UserIP       string `gorm:"size:32" json:"userIP"`
	UserAddr     string `gorm:"size:64" json:"userAddr"`
}
