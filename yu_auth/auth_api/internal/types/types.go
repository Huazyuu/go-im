// Code generated by goctl. DO NOT EDIT.
package types

type AuthenticationReponse struct {
	UserID uint `json:"userID"`
	Role   int  `json:"role"`
}

type AuthenticationRequest struct {
	Token     string `header:"token,optional"`     // 请求token
	ValidPath string `header:"validpath,optional"` // 请求的地址
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OpenLoginInfoResponse struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"` // 跳转地址
}
