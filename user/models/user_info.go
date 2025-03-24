package models

import "server/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	UserNickname string `json:"userNickname"`
	UserPwd      string `json:"userPwd"`
	UserAbstract string `json:"userAbstract"`
	UserAvatar   string `json:"userAvatar"`
	UserIP       string `json:"userIP"`
	UserAddr     string `json:"userAddr"`
}
