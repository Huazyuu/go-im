package models

import (
	"server/common/models"
	"server/common/models/commonType/cmsg"
)

type ChatModel struct {
	models.Model

	ChatSenderID   uint `json:"chatSenderID"`
	ChatRecvUserID uint `json:"chatRecvUserID"`
	// 1:text 2:img 3:video 4:file 5:voice 6:voice_call 7:video:call 8:withdraw 9:replyMsg 10:quoteMsg
	ChatMsgType    int8         `json:"chatMsgType"`
	ChatMsgPreview string       `gorm:"size:64" json:"chatMsgPreview"`
	ChatMsg        cmsg.Msg     `json:"chatMsg"`
	ChatSysMsg     *cmsg.SysMsg `json:"chatSysMsg"` // 系统提示
}

func (ChatModel) TableName() string {
	return "chat"
}
