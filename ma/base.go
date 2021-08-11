package ma

import (
	"github.com/cliod/wx-go/common"
)

// MiniProgramState 小程序版本
//
// 小程序订阅消息跳转小程序类型
// developer为开发版；trial为体验版；formal为正式版；默认为正式版
type MiniProgramState string

// MiniProgramLang 小程序语言
//
// 进入小程序查看的语言类型
// 支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
type MiniProgramLang string

const (
	DEVELOPER MiniProgramState = "developer" // 开发版
	TRIAL     MiniProgramState = "trial"     // 体验版
	FORMAL    MiniProgramState = "formal"    // 正式版
)

const (
	ZhCn MiniProgramLang = "zh_CN" // 简体中文
	EnUs MiniProgramLang = "en_US" // 英文
	ZhHk MiniProgramLang = "zh_HK" // 繁体中文
	ZhTw MiniProgramLang = "zh_TW" // 繁体中文
)

// WxMaJsCode2SessionResult code换取session_key接口的响应
// 文档地址：https://mp.weixin.qq.com/debug/wxadoc/dev/api/api-login.html#wxloginobject
type WxMaJsCode2SessionResult struct {
	common.WxCommonErr

	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
}

// WxMaUnionIdResult 用户的UnionId接口的响应
type WxMaUnionIdResult struct {
	common.WxCommonErr

	UnionId string `json:"unionid"`
}
