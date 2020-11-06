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

// 用户数据
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

// 水印数据
type Watermark struct {
	Timestamp uint64 `json:"timestamp"`
	AppId     string `json:"appid"`
}

// 手机号数据
type PhoneNumberInfo struct {
	PhoneNumber     string    `json:"phone_number"`
	PurePhoneNumber string    `json:"pure_phone_number"`
	CountryCode     string    `json:"country_code"`
	Watermark       Watermark `json:"watermark"`
}

// 二维码数据
type BaseCode struct {
	Width     int           `json:"width"`
	AutoColor bool          `json:"auto_color"`
	IsHyaline bool          `json:"is_hyaline"`
	LineColor CodeLineColor `json:"line_color"`
}

// 二维码数据
type WxaCode struct {
	BaseCode
	Path string `json:"path"`
}

// 二维码数据
type WxaCodeUnlimited struct {
	BaseCode
	Scene string `json:"scene"`
	Page  string `json:"page"`
}

// 二维码线颜色
type CodeLineColor struct {
	R string
	B string
	G string
}

// 默认白色
func DefaultCodeLineColor() CodeLineColor {
	return CodeLineColor{R: "0", B: "0", G: "0"}
}
