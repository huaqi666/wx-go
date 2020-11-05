package ma

import (
	"github.com/cliod/wx-go/common"
)

// js_code session_info
type JsCode2SessionResult struct {
	common.Err
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
}

type UserInfo struct {
	Openid    string    `json:"openid"`
	Nickname  string    `json:"nickname"`
	Gender    string    `json:"gender"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatar_url"`
	UnionID   string    `json:"union_id"`
	Watermark Watermark `json:"watermark"`
}

type Watermark struct {
	Timestamp uint64 `json:"timestamp"`
	AppId     string `json:"appid"`
}

type PhoneNumberInfo struct {
	PhoneNumber     string    `json:"phone_number"`
	PurePhoneNumber string    `json:"pure_phone_number"`
	CountryCode     string    `json:"country_code"`
	Watermark       Watermark `json:"watermark"`
}
