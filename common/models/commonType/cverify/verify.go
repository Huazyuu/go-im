package cverify

type VerifyType int8

const (
	VerifyUnableAdd   VerifyType = iota // 不允许
	VerifyAbleAdd                       // 允许添加
	VerifyMsg                           // 验证消息
	VerifyAnswer                        // 回答问题
	VerifyRightAnswer                   // 回答正确问题
)
