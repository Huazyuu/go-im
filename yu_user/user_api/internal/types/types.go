// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"Role"`
}

type UserInfoResponse struct {
	UserID         uint   `json:"userID"`
	UserRole       int8   `json:"userRole"`
	UserName       string `json:"userName"`
	UserNickname   string `json:"userNickname"`
	UserAbstract   string `json:"userAbstract"`
	UserAvatar     string `json:"userAvatar"`
	RegisterSource string `json:"registerSource"`
}
