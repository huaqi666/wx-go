package common

import (
	"fmt"
	"github.com/cliod/wx-go/common/util"
	"time"
)

// 公共接口返回错误
type Err struct {
	ErrCode uint64 `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Err) Error() string {
	var errorMsg string
	if msg := e.ErrCode; msg != 0 {
		errorMsg = fmt.Sprintf("错误代码：%d", msg)
	}
	if msg := util.TrimToUpper(e.ErrMsg); msg != "" {
		errorMsg += "，错误信息：" + msg
	}
	return errorMsg
}

// 自定义error
type ErrMsg struct {
	Err
	Msg string `json:"msg"`
}

func (e *ErrMsg) Error() string {
	var errorMsg string
	if msg := e.ErrCode; msg != 0 {
		errorMsg = fmt.Sprintf("错误代码：%d", msg)
	}
	if msg := util.TrimToUpper(e.ErrMsg); msg != "" {
		errorMsg += "，错误信息：" + msg
	}
	if msg := util.TrimToUpper(e.Msg); msg != "" {
		errorMsg += "，详细信息：" + msg
	}
	return errorMsg
}

func ErrorOf(msg string, params ...interface{}) *ErrMsg {
	return &ErrMsg{
		Msg: fmt.Sprintf(msg, params...),
	}
}

//小程序access_token
type AccessToken struct {
	Err
	AccessToken string    `json:"access_token"`
	ExpiresIn   uint64    `json:"expires_in"`
	Time        time.Time `json:"time"`
}

// ticket类型
type TicketType string

const (
	JSAPI  TicketType = "jsapi"
	SDK    TicketType = "2"
	WxCard TicketType = "wx_card"
)

// 授权页ticket
type Ticket struct {
	Err
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
