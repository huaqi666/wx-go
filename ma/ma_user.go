package ma

import (
	"encoding/json"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
)

type WxMaUserService interface {
	// jsCode换取openid
	GetSessionInfo(jsCode string) (*WxMaJsCode2SessionResult, error)
	// 解密用户敏感数据
	GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error)
	// 解密用户手机号信息.
	GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error)
	// 验证用户信息完整性
	CheckUserInfo(sessionKey, rawData, signature string) bool
	/* 上报用户数据后台接口.
	   小游戏可以通过本接口上报key-value数据到用户的CloudStorage。
	   文档参考https://developers.weixin.qq.com/minigame/dev/document/open-api/data/setUserStorage.html */
	SetUserStorage(kvMap map[string]string, sessionKey, openid string) error
}

type WxMaUserServiceImpl struct {
	service WxMaService
}

func newWxMaUserService(service WxMaService) *WxMaUserServiceImpl {
	return &WxMaUserServiceImpl{
		service: service,
	}
}

func (u *WxMaUserServiceImpl) GetSessionInfo(jsCode string) (*WxMaJsCode2SessionResult, error) {
	return u.service.JsCode2SessionInfo(jsCode)
}

func (u *WxMaUserServiceImpl) GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error) {
	return GetUserInfo(sessionKey, encryptedData, ivStr)
}

func (u *WxMaUserServiceImpl) GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error) {
	return GetPhoneNoInfo(sessionKey, encryptedData, ivStr)
}

func (u *WxMaUserServiceImpl) CheckUserInfo(sessionKey, rawData, signature string) bool {
	return CheckUserInfo(sessionKey, rawData, signature)
}

func (u *WxMaUserServiceImpl) SetUserStorage(kvMap map[string]string, sessionKey, openid string) error {
	c := u.service.GetWxMaConfig()
	var param = map[string]interface{}{}
	var arr []interface{}
	for k, v := range kvMap {
		arr = append(arr, map[string]interface{}{"key": k, "value": v})
	}
	param["kv_list"] = arr
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}
	sign := util.HmacSha256(string(b), sessionKey)
	_, err = u.service.Post(common.MaSetUserStorage, common.PostJsonContentType, param, c.GetAppID(), sign, openid, "hmac_sha256")
	return err
}
