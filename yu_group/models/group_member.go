package models

// ChatMemberModel 群成员表
type ChatMemberModel struct {
	GroupID    uint       `json:"groupID"`
	GroupModel GroupModel `json:"-" gorm:"foreignKey:GroupID"`

	UserID     uint   `json:"userID"`
	MemberName string `gorm:"size:32" json:"memberName"`

	Role            int  `json:"role"`            // 1:群主 2:管理 3:normal
	ProhibitionTime *int `json:"prohibitionTime"` // 禁言时间 单位分钟
}

func (ChatMemberModel) TableName() string {
	return "chat_member"
}
