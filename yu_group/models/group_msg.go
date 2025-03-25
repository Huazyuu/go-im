package models

import (
	"server/common/models"
	"server/common/models/commonType/cmsg"
)

// GroupMsgModel 群消息表
type GroupMsgModel struct {
	models.Model

	GroupID    uint       `json:"groupID"` // 群id
	GroupModel GroupModel `json:"-" gorm:"foreignKey:GroupID"`

	SendUserID uint `json:"sendUserID"` // 发送者id

	MsgPreview string `gorm:"64" json:"msgPreview"` // 消息预览
	// 消息类型 1 文本类型  2 图片消息  3 视频消息 4 文件消息 5 语音消息  6 语言通话  7 视频通话  8 撤回消息 9回复消息 10 引用消息 11 at(@)消息
	MsgType   cmsg.MsgType `json:"msgType"`
	Msg       cmsg.Msg     `json:"msg"`       // 消息内容
	SystemMsg *cmsg.SysMsg `json:"systemMsg"` // 系统提示
}

func (GroupMsgModel) TableName() string {
	return "group_msg"
}
