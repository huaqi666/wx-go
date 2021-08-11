package common

import (
	"fmt"
	"github.com/cliod/wx-go/common/util"
	"time"
)

// WechatResponse 微信API响应
type WechatResponse interface {
	ErrMsg() string
	HasErrOccurred() bool
}

type WxBaseResponse struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (e *WxBaseResponse) ErrMsg() string {
	return e.Errmsg
}

func (e *WxBaseResponse) HasErrOccurred() bool {
	return e.Errcode != 0
}

// WxCommonErr 公共接口返回错误
type WxCommonErr struct {
	WxBaseResponse
}

func (e *WxCommonErr) Error() string {
	var errorMsg string
	if msg := e.Errcode; msg != 0 {
		errorMsg = fmt.Sprintf("错误代码：%d", msg)
	}
	if msg := util.TrimToUpper(e.Errmsg); msg != "" {
		errorMsg += "，错误信息：" + msg
	}
	return errorMsg
}

// WxPayErrMsg 自定义error
type WxPayErrMsg struct {
	WxCommonErr
	Msg string `json:"msg"`
}

func (e *WxPayErrMsg) Error() string {
	var errorMsg string
	if msg := e.Errcode; msg != 0 {
		errorMsg = fmt.Sprintf("错误代码：%d", msg)
	}
	if msg := util.TrimToUpper(e.Errmsg); msg != "" {
		errorMsg += "，错误信息：" + msg
	}
	if msg := util.TrimToUpper(e.Msg); msg != "" {
		errorMsg += "，详细信息：" + msg
	}
	return errorMsg
}

func ErrorOf(msg string, params ...interface{}) *WxPayErrMsg {
	return &WxPayErrMsg{
		Msg: fmt.Sprintf(msg, params...),
	}
}

// AccessToken 小程序access_token
type AccessToken struct {
	WxCommonErr
	AccessToken string    `json:"access_token"`
	ExpiresIn   uint64    `json:"expires_in"`
	Time        time.Time `json:"time"`
}

// TicketType ticket类型
type TicketType string

func (t TicketType) String() string {
	return string(t)
}

const (
	JSAPI  TicketType = "jsapi"
	SDK    TicketType = "sdk"
	WxCard TicketType = "wx_card"
)

// Ticket 授权页ticket
type Ticket struct {
	WxCommonErr
	Ticket    string    `json:"ticket"`
	ExpiresIn uint64    `json:"expires_in"`
	Time      time.Time `json:"time"`
	Type      string    `json:"type"`
}

// WxJsapiSignature jsapi signature.
type WxJsapiSignature struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}
