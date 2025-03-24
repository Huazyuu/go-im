package cmsg

import "time"

type SysMsg struct {
	Type uint8 `json:"type"` // 违规 1:黄 2:恐 3:政 4:不正当言论
}

type MsgType int8

const (
	content MsgType = iota + 1
	imageMsg
	videoMsg
	fileMsg
	voiceMsg
	voiceCallMsg
	videoCallMsg
	withdrawMsg
	replyMsg
	quoteMsg
	atMsg
)

// Msg 定义消息，包含不同类型消息的通用信息
// 消息类型说明：
// 1: text 文本消息
// 2: img 图片消息
// 3: video 视频消息
// 4: file 文件消息
// 5: voice 语音消息
// 6: voice_call 语音通话消息
// 7: video_call 视频通话消息
// 8: withdraw 撤回消息
// 9: replyMsg 回复消息
// 10: quoteMsg 引用消息
// 11: @ 用户的消息（群聊才有）
type Msg struct {
	Type         MsgType       `json:"type"`         // 消息类型，和 msgType 一致
	Content      *string       `json:"content"`      // 文本消息，当消息类型为 1 时使用
	ImageMsg     *ImgMsg       `json:"imageMsg"`     // 图片消息
	VideoMsg     *VideoMsg     `json:"videoMsg"`     // 视频消息
	FileMsg      *FileMsg      `json:"fileMsg"`      // 文件消息
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`     // 语音消息
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"` // 语音通话消息
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"` // 视频通话消息
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`  // 撤回消息
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`     // 回复消息
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`     // 引用消息
	AtMsg        *AtMsg        `json:"atMsg"`        // @ 用户的消息，群聊时使用
}

// ImgMsg 定义图片消息
type ImgMsg struct {
	Title string `json:"title"` // 图片标题
	Src   string `json:"src"`   // 图片的源地址
}

// VideoMsg 定义视频消息
type VideoMsg struct {
	Title    string `json:"title"`    // 视频标题
	Src      string `json:"src"`      // 视频的源地址
	Duration int    `json:"duration"` // 视频时长，单位为秒
}

// FileMsg 定义文件消息
type FileMsg struct {
	Title string `json:"title"` // 文件标题
	Src   string `json:"src"`   // 文件的源地址
	Size  int64  `json:"size"`  // 文件大小
	Type  string `json:"type"`  // 文件后缀
}

// VoiceMsg 定义语音消息
type VoiceMsg struct {
	Src      string `json:"src"`      // 语音文件的源地址
	Duration int    `json:"duration"` // 语音时长，单位为秒
}

// VoiceCallMsg 定义语音通话消息
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` // 语音通话开始时间
	EndTime   time.Time `json:"endTime"`   // 语音通话结束时间
	EndReason int8      `json:"endReason"` // 语音通话结束原因
	// 结束原因说明：
	// 0: 发起方挂断
	// 1: 接收方挂断
	// 2: 网络原因挂断
}

// VideoCallMsg 定义视频通话消息
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` // 视频通话开始时间
	EndTime   time.Time `json:"endTime"`   // 视频通话结束时间
	EndReason int8      `json:"endReason"` // 视频通话结束原因
	// 结束原因说明：
	// 0: 发起方挂断
	// 1: 接收方挂断
	// 2: 网络原因挂断
	// 3: 未打通
}

// WithdrawMsg 定义撤回消息
type WithdrawMsg struct {
	Content   string `json:"content"` // 撤回的提示词
	OriginMsg *Msg   `json:"-"`       // 原消息，不参与 JSON 序列化
}

// ReplyMsg 定义回复消息
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息 ID
	Content string `json:"content"` // 回复的文本消息，目前仅支持文本回复
	Msg     *Msg   `json:"msg"`     // 被回复的消息
}

// QuoteMsg 定义引用消息
type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息 ID
	Content string `json:"content"` // 引用的文本消息，目前仅支持文本引用
	Msg     *Msg   `json:"msg"`     // 被引用的消息
}

// AtMsg 定义 @ 用户消息
type AtMsg struct {
	UserID  uint   `json:"userID"`  // 被 @ 的用户 ID
	Content string `json:"content"` // 包含 @ 的文本消息
	Msg     *Msg   `json:"msg"`     // 关联的消息
}
