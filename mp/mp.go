// 微信公众号接口
package mp

import (
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
)

type WxMpService interface {
	common.WxService
	common.WxJsapi

	// 获取配置
	GetWxMpConfig() WxMpConfig
	// 设置配置
	SetWxMpConfig(WxMpConfig)

	// 验证消息的确来自微信服务器
	CheckSignature(timestamp, nonce, signature string) bool

	// 获取用户接口
	GetWxMpUserService() WxMpUserService
	// 设置(用户自定义的)用户接口
	SetWxMpUserService(WxMpUserService)
	// 获取二维码接口
	GetWxMpQrcodeService() WxMpQrcodeService
	// 设置(用户自定义的)二维码接口
	SetWxMpQrcodeService(WxMpQrcodeService)
}

type WxMpServiceImpl struct {
	common.WxServiceImpl

	config        WxMpConfig
	userService   WxMpUserService
	qrcodeService WxMpQrcodeService
}

func newWxMpService(config WxMpConfig) *WxMpServiceImpl {
	impl := &WxMpServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMpConfig(config)
	impl.userService = newWxMpUserService(impl)
	impl.qrcodeService = newWxMpQrcodeService(impl)
	return impl
}

func (s *WxMpServiceImpl) CheckSignature(timestamp, nonce, signature string) bool {
	return util.CheckSignature(s.GetWxMpConfig().GetToken(), timestamp, nonce, signature)
}

func (s *WxMpServiceImpl) GetWxMpUserService() WxMpUserService {
	return s.userService
}

func (s *WxMpServiceImpl) GetWxMpQrcodeService() WxMpQrcodeService {
	return s.qrcodeService
}

func (s *WxMpServiceImpl) SetWxMpUserService(userService WxMpUserService) {
	s.userService = userService
}

func (s *WxMpServiceImpl) SetWxMpQrcodeService(qrcodeService WxMpQrcodeService) {
	s.qrcodeService = qrcodeService
}

func (s *WxMpServiceImpl) GetWxMpConfig() WxMpConfig {
	return s.config
}

func (s *WxMpServiceImpl) SetWxMpConfig(config WxMpConfig) {
	s.SetWxConfig(config)
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func (s *WxMpServiceImpl) GetJsapiTicket() (*common.Ticket, error) {
	return s.ForceGetJsapiTicket(false)
}

func (s *WxMpServiceImpl) ForceGetJsapiTicket(forceRefresh bool) (*common.Ticket, error) {
	return s.ForceGetTicket(common.JSAPI, forceRefresh)
}

func (s *WxMpServiceImpl) GetTicket(ticketType common.TicketType) (*common.Ticket, error) {
	return s.ForceGetTicket(ticketType, false)
}

func (s *WxMpServiceImpl) ForceGetTicket(ticketType common.TicketType, forceRefresh bool) (*common.Ticket, error) {
	conf := s.GetWxMpConfig()
	b := conf.IsTicketExpired(ticketType)
	if forceRefresh || b {
		tt, err := s.getTicket(ticketType)
		conf.UpdateTicket(ticketType, tt)
		return tt, err
	}
	return conf.GetTicket(ticketType), nil
}

func (s *WxMpServiceImpl) getTicket(ticketType common.TicketType) (*common.Ticket, error) {
	var ticket common.Ticket

	err := s.GetFor(&ticket, common.MpGetTicketUrl, ticketType)
	return &ticket, err
}

func (s *WxMpServiceImpl) CreateJsapiSignature(url string) (*common.WxJsapiSignature, error) {
	jsapiTicket, _ := s.GetJsapiTicket()
	appId := s.GetWxMpConfig().GetAppID()
	return common.CreateJsapiSignature(url, appId, jsapiTicket.Ticket)
}

func NewWxMpServiceBy(appId, secret string) WxMpService {
	return newWxMpService(newWxMpConfig(appId, secret))
}

func NewWxMpService(config WxMpConfig) WxMpService {
	if config == nil {
		config = new(WxMpConfigImpl)
	}
	return newWxMpService(config)
}

// 将方法静态方便使用

// 获取accessToken
func GetAccessToken(appId, secret string) (*common.AccessToken, error) {
	return NewWxMpServiceBy(appId, secret).GetAccessToken()
}
