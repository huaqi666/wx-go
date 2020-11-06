package ma

type WxMaUserService interface {
	// jsCode换取openid
	GetSessionInfo(jsCode string) (*JsCode2SessionResult, error)
	// 解密用户敏感数据
	GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error)
	// 解密用户手机号信息.
	GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error)
	// 验证用户信息完整性
	CheckUserInfo(sessionKey, rawData, signature string) bool
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
	return GetUserInfo(sessionKey, encryptedData, ivStr)
}

func (u *WxMaUserServiceImpl) GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error) {
	return GetPhoneNoInfo(sessionKey, encryptedData, ivStr)
}

func (u *WxMaUserServiceImpl) CheckUserInfo(sessionKey, rawData, signature string) bool {
	return CheckUserInfo(sessionKey, rawData, signature)
}
