package mp

import (
	"github.com/cliod/wx-go/common"
)

type WxMpService interface {
	common.WxService

	// 获取配置
	GetWxMpConfig() WxMpConfig
	// 设置配置
	SetWxMpConfig(WxMpConfig)

	// 验证消息的确来自微信服务器
	CheckSignature(timestamp, nonce, signature string) bool

	// 获得jsapi_ticket,不强制刷新jsapi_ticket.
	GetJsapiTicket() (*Ticket, error)
	// 获得jsapi_ticket.
	// 获得时会检查jsapiToken是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	// 详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	ForceGetJsapiTicket(bool) (*Ticket, error)

	// 获得ticket,不强制刷新ticket.
	GetTicket(TicketType) (*Ticket, error)
	// 获得时会检查 Token是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error)

	// 创建调用jsapi时所需要的签名.
	// 详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	CreateJsapiSignature(url string) (*WxJsapiSignature, error)

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
	return CheckSignature(s.GetWxMpConfig().GetToken(), timestamp, nonce, signature)
}

func (s *WxMpServiceImpl) GetJsapiTicket() (*Ticket, error) {
	return s.ForceGetJsapiTicket(false)
}

func (s *WxMpServiceImpl) ForceGetJsapiTicket(forceRefresh bool) (*Ticket, error) {
	return s.ForceGetTicket(JSAPI, forceRefresh)
}

func (s *WxMpServiceImpl) GetTicket(ticketType TicketType) (*Ticket, error) {
	return s.ForceGetTicket(ticketType, false)
}

func (s *WxMpServiceImpl) ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error) {
	conf := s.GetWxMpConfig()
	b := conf.IsTicketExpired(ticketType)
	if forceRefresh || b {
		tt, err := s.getTicket(ticketType)
		conf.UpdateTicket(ticketType, tt)
		return tt, err
	}
	return conf.GetTicket(ticketType), nil
}

func (s *WxMpServiceImpl) getTicket(ticketType TicketType) (*Ticket, error) {
	var ticket Ticket

	err := s.GetFor(&ticket, common.MpGetTicketUrl, ticketType)
	return &ticket, err
}

func (s *WxMpServiceImpl) CreateJsapiSignature(url string) (*WxJsapiSignature, error) {
	jsapiTicket, _ := s.GetJsapiTicket()
	appId := s.GetWxMpConfig().GetAppID()
	return CreateJsapiSignature(url, appId, jsapiTicket.Ticket)
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

func NewWxMpServiceBy(appId, secret string) WxMpService {
	return newWxMpService(newWxMpConfig(appId, secret))
}

func NewWxMpService(config WxMpConfig) WxMpService {
	if config == nil {
		config = new(WxMpConfigImpl)
	}
	return newWxMpService(config)
}
