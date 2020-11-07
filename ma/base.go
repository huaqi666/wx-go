package ma

import (
	"github.com/cliod/wx-go/common"
)

type MiniProgramState string

type MiniProgramLang string

const (
	DEVELOPER MiniProgramState = "developer"
	TRIAL     MiniProgramState = "trial"
	FORMAL    MiniProgramState = "formal"
)

const (
	ZhCn MiniProgramLang = "zh_CN"
	EnUs MiniProgramLang = "en_US"
	ZhHk MiniProgramLang = "zh_HK"
	ZhTw MiniProgramLang = "zh_TW"
)

// js_code session_info
type JsCode2SessionResult struct {
	common.Err

	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
}

type WxMaUnionIdResult struct {
	common.Err

	UnionId string `json:"unionid"`
}
