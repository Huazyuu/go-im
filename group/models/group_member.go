package models

// ChatMemberModel 群成员表
type ChatMemberModel struct {
	GroupID    uint   `json:"groupID"`
	UserID     uint   `json:"userID"`
	MemberName string `json:"memberName"`
	Role       int    `json:"role"`     // 1:群主 2:管理 3:normal
	Prohibit   *int   `json:"prohibit"` // 禁言时间min
}
