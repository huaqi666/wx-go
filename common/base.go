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
