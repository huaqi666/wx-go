package common

import "time"

// 接口返回错误
type Err struct {
	ErrCode uint64 `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//小程序access_token
type AccessToken struct {
	Err
	AccessToken string    `json:"access_token"`
	ExpiresIn   uint64    `json:"expires_in"`
	Time        time.Time `json:"time"`
}
