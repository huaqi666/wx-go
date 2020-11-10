package ma

import (
	"github.com/cliod/wx-go/common"
	"time"
)

// 小程序版本
type MiniProgramState string

// 小程序语言
type MiniProgramLang string

type TicketType string

const (
	JSAPI  TicketType = "jsapi"
	SDK    TicketType = "2"
	WxCard TicketType = "wx_card"
)

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

// 用户的UnionId
type WxMaUnionIdResult struct {
	common.Err

	UnionId string `json:"unionid"`
}

// 授权页ticket
type Ticket struct {
	common.Err
	Ticket    string    `json:"ticket"`
	ExpiresIn uint64    `json:"expires_in"`
	Time      time.Time `json:"time"`
	Type      string    `json:"type"`
}

// jspai signature.
type WxJsapiSignature struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}
