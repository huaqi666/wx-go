package ma

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"wx-go/common/util"
)

type WxMaUserService interface {
	// jsCode换取openid
	GetSessionInfo(jsCode string) (*JsCode2SessionResult, error)
	// 用户用户信息
	GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error)
	// 解密用户手机号信息.
	GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error)
	// 验证用户信息完整性
	checkUserInfo(sessionKey, rawData, signature string) bool
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
	Timestamp string `json:"timestamp"`
	AppId     string `json:"appid"`
}

type PhoneNumberInfo struct {
	PhoneNumber     string    `json:"phone_number"`
	PurePhoneNumber string    `json:"pure_phone_number"`
	CountryCode     string    `json:"country_code"`
	Watermark       Watermark `json:"watermark"`
}

type WxMaUserServiceImpl struct {
	service WxMaService
}

func newWxMaUserService(service WxMaService) *WxMaUserServiceImpl {
	return &WxMaUserServiceImpl{
		service: service,
	}
}

func (u *WxMaUserServiceImpl) GetSessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	return u.service.JsCode2SessionInfo(jsCode)
}

func (u *WxMaUserServiceImpl) GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error) {
	var usrInfo UserInfo
	err := util.Decrypt(&usrInfo, sessionKey, encryptedData, ivStr)
	return &usrInfo, err
}

func (u *WxMaUserServiceImpl) GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error) {
	var info PhoneNumberInfo
	err := util.Decrypt(&info, sessionKey, encryptedData, ivStr)
	return &info, err
}

func (u *WxMaUserServiceImpl) checkUserInfo(sessionKey, rawData, signature string) bool {
	sum := sha1.Sum([]byte(rawData + sessionKey))
	return strings.ToLower(hex.EncodeToString(sum[:])) == signature
}
