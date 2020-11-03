package ma

import (
	"fmt"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
)

type WxMaService interface {
	common.WxService

	// 获取配置
	GetWxMaConfig() WxMaConfig
	// 设置配置
	SetWxMaConfig(config WxMaConfig)

	// jsCode换取openid
	JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error)
	// 验证消息的确来自微信服务器
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421135319&token=&lang=zh_CN
	CheckSignature(timestamp, nonce, signature string) bool

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

func newWxMaServiceFor(config WxMaConfig) *WxMaServiceImpl {
	impl := WxMaServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMaConfig(config)
	impl.userService = newWxMaUserService(&impl)
	impl.qrCodeService = newWxMaQrcodeService(&impl)
	return &impl
}

func newWxMaService(appId, secret string) *WxMaServiceImpl {
	return newWxMaServiceFor(newWxMaConfig(appId, secret))
}

func (s *WxMaServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.GetFor(&jsr, common.MaSessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *WxMaServiceImpl) CheckSignature(timestamp, nonce, signature string) bool {
	arr := []string{s.GetWxMaConfig().GetToken(), timestamp, nonce}
	gen, err := util.Gen(arr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return gen == signature
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

func NewWxMaServiceFor(config WxMaConfig) WxMaService {
	return newWxMaServiceFor(config)
}
