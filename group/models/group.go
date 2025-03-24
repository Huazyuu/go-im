package models

import (
	"server/common/models"
	"server/common/models/commonType/cverify"
)

type GroupModel struct {
	models.Model
	GroupName     string `json:"groupName"`
	GroupAbstract string `json:"groupAbstract"` // 简介
	GroupAvatar   string `json:"groupAvatar"`
	GroupCreator  string `json:"groupCreator"`
	GroupSize     int    `json:"groupSize"` // 群聊规模 20 100 200 1000

	GroupIsSearch     bool `json:"groupIsSearch"`     // 是否可以搜索加群
	GroupIsInvite     bool `json:"groupIsInvite"`     // 是否可邀请人加入
	GroupIsTmpSession bool `json:"groupIsTmpSession"` // 是否可以开启临时消息
	GroupIsProhibit   bool `json:"groupIsProhibit"`   // 是否全员禁言

	GroupVerify         cverify.VerifyType            `json:"groupVerify"` // 0不允许添加 1验证消息 2回答问题 4需要正确回答问题
	GroupVerifyQuestion *cverify.VerificationQuestion `json:"groupVerifyQuestion"`
}
