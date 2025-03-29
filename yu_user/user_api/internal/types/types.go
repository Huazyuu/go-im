// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"Role"`
}

type UserInfoResponse struct {
	UserID               uint                  `json:"userID"`
	UserName             string                `json:"userName"`
	UserNickname         string                `json:"userNickname"`
	UserAbstract         string                `json:"userAbstract"`
	UserAvatar           string                `json:"userAvatar"`
	RecallMessage        *string               `json:"recallMessage"`        // 撤回消息提示内容
	IsFriendOnlineNotify bool                  `json:"isFriendOnlineNotify"` // 是否通知好友上线
	IsSound              bool                  `json:"isSound"`              // 提醒声音
	IsSecureLink         bool                  `json:"isSecureLink"`         // 安全链接,不直接转跳
	IsSavePwd            bool                  `json:"isSavePwd"`            // 保存密码
	SearchUser           int8                  `json:"searchUser"`           // 查到到我的方式 0不允许别人查找 1用户号查 2可以用昵称
	Verification         int8                  `json:"verification"`         // 好友验证 0不允许添加 1验证消息 2回答问题 4需要正确回答问题
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion"` // 验证问题 3,4需要
}

type UserInfoUpdateRequest struct {
	UserID               uint                  `header:"User-ID"`
	UserNickname         string                `json:"userNickname,optional" user:"user_nickname"`
	UserAbstract         string                `json:"userAbstract,optional" user:"user_abstract"`
	UserAvatar           string                `json:"userAvatar,optional" user:"user_avatar"`
	RecallMessage        *string               `json:"recallMessage,optional" user_conf:"recall_msg"`                     // 撤回消息提示内容
	IsFriendOnlineNotify bool                  `json:"isFriendOnlineNotify,optional" user_conf:"is_friend_online_notify"` // 是否通知好友上线
	IsSound              bool                  `json:"isSound,optional" user_conf:"is_sound"`                             // 提醒声音
	IsSecureLink         bool                  `json:"isSecureLink,optional" user_conf:"is_secure_link"`                  // 安全链接,不直接转跳
	IsSavePwd            bool                  `json:"isSavePwd,optional" user_conf:"is_save_pwd"`                        // 保存密码
	SearchUser           int8                  `json:"searchUser,optional" user_conf:"search_user"`                       // 查到到我的方式 0不允许别人查找 1用户号查 2可以用昵称
	Verification         int8                  `json:"verification,optional" user_conf:"verification"`                    // 好友验证 0不允许添加 1验证消息 2回答问题 4需要正确回答问题
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion,optional" user_conf:"verification_question"`   // 验证问题 3,4需要
}

type UserInfoUpdateResponse struct {
}

type VerificationQuestion struct {
	Question1 *string `json:"question1,optional" user_conf:"question1"`
	Question2 *string `json:"question2,optional" user_conf:"question2"`
	Question3 *string `json:"question3,optional" user_conf:"question3"`
	Answer1   *string `json:"answer1,optional" user_conf:"answer1"`
	Answer2   *string `json:"answer2,optional" user_conf:"answer2"`
	Answer3   *string `json:"answer3,optional" user_conf:"answer3"`
}
