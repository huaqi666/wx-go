package ma

import (
	"encoding/json"
	"fmt"
	"wx-go/common"
)

type WxMaService interface {
	common.WxService

	// 获取配置
	GetWxMaConfig() WxMaConfig
	// 设置配置
	SetWxMaConfig(config WxMaConfig)

	// jsCode换取openid
	JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error)

	// 获取用户接口
	GetWxMaUserService() WxMaUserService
	// 获取二维码接口
	GetWxMaQrcodeService() WxMaQrcodeService
}

type WxMaServiceImpl struct {
	common.WxServiceImpl

	config        WxMaConfig
	userService   WxMaUserService
	qrCodeService WxMaQrcodeService
}

func newService(appId, secret string) *WxMaServiceImpl {
	impl := WxMaServiceImpl{
		config: newWxMaConfig(appId, secret),
	}
	at, _ := impl.GetAccessToken()
	impl.config.SetAccessToken(at)
	impl.userService = newWxMaUserService(&impl)
	impl.qrCodeService = newWxMaQrcodeService(&impl)
	return &impl
}

func (s *WxMaServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.GetFor(&jsr, common.SessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *WxMaServiceImpl) GetWxMaUserService() WxMaUserService {
	return s.userService
}

func (s *WxMaServiceImpl) GetWxMaQrcodeService() WxMaQrcodeService {
	return s.qrCodeService
}

func (s *WxMaServiceImpl) GetWxMaConfig() WxMaConfig {
	return s.config
}

func (s *WxMaServiceImpl) SetWxMaConfig(config WxMaConfig) {
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func (s *WxMaServiceImpl) GetWxConfig() common.WxConfig {
	return s.GetWxMaConfig()
}

func (s *WxMaServiceImpl) SetWxConfig(config common.WxConfig) {
	var c WxMaConfigImpl
	b, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s.SetWxMaConfig(&c)
}

func NewService(appId, secret string) WxMaService {
	return newService(appId, secret)
}
