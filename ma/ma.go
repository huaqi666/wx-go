package ma

import (
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

func newWxMaService(appId, secret string) *WxMaServiceImpl {
	impl := WxMaServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMaConfig(newWxMaConfig(appId, secret))
	impl.userService = newWxMaUserService(&impl)
	impl.qrCodeService = newWxMaQrcodeService(&impl)
	return &impl
}

func (s *WxMaServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.GetFor(&jsr, common.MaSessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
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
	s.SetWxConfig(config)
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func NewWxMaService(appId, secret string) WxMaService {
	return newWxMaService(appId, secret)
}
