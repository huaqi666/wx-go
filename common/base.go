package common

import (
	"fmt"
	"time"
)

// 接口返回错误
type Err struct {
	ErrCode uint64 `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("错误代码：%d, 错误信息：%s", e.ErrCode, e.ErrMsg)
}

type ErrMsg struct {
	Err
	Msg string `json:"msg"`
}

func (e *ErrMsg) Error() string {
	return fmt.Sprintf("错误代码：%d, 错误信息：%s，信息：%s", e.ErrCode, e.ErrMsg, e.Msg)
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
