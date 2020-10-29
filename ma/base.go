package ma

import (
	"wx-go/common"
)

// js_code session_info
type JsCode2SessionResult struct {
	common.Err
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
}