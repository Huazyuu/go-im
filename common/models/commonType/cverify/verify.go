package cverify

const (
	VerifyUnableAdd   int8 = iota // 不允许
	VerifyAbleAdd                 // 允许添加
	VerifyMsg                     // 验证消息
	VerifyAnswer                  // 回答问题
	VerifyRightAnswer             // 回答正确问题
)
