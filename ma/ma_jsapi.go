package ma

import "github.com/cliod/wx-go/common"

type WxMaJsapiService interface {
	// 获得卡券api_ticket,不强制刷新api_ticket
	GetCardApiTicket() (*Ticket, error)
	// 获得卡券api_ticket
	// 获得时会检查apiToken是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	ForceGetCardApiTicket(forceRefresh bool) (*Ticket, error)
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
}

type WxMaJsapiServiceImpl struct {
	service WxMaService
}

func newWxMaJsapiService(service WxMaService) *WxMaJsapiServiceImpl {
	return &WxMaJsapiServiceImpl{
		service: service,
	}
}

func (s *WxMaJsapiServiceImpl) GetCardApiTicket() (*Ticket, error) {
	return s.ForceGetCardApiTicket(false)
}

func (s *WxMaJsapiServiceImpl) ForceGetCardApiTicket(forceRefresh bool) (*Ticket, error) {
	return s.ForceGetTicket(WxCard, forceRefresh)
}

func (s *WxMaJsapiServiceImpl) GetJsapiTicket() (*Ticket, error) {
	return s.ForceGetJsapiTicket(false)
}

func (s *WxMaJsapiServiceImpl) ForceGetJsapiTicket(forceRefresh bool) (*Ticket, error) {
	return s.ForceGetTicket(JSAPI, forceRefresh)
}

func (s *WxMaJsapiServiceImpl) GetTicket(ticketType TicketType) (*Ticket, error) {
	return s.ForceGetTicket(ticketType, false)
}

func (s *WxMaJsapiServiceImpl) ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error) {
	conf := s.service.GetWxMaConfig()
	b := conf.IsTicketExpired(ticketType)
	if forceRefresh || b {
		tt, err := s.getTicket(ticketType)
		conf.UpdateTicket(ticketType, tt)
		return tt, err
	}
	return conf.GetTicket(ticketType), nil
}

func (s *WxMaJsapiServiceImpl) getTicket(ticketType TicketType) (*Ticket, error) {
	var ticket Ticket

	err := s.service.GetFor(&ticket, common.MpGetTicketUrl, ticketType)
	return &ticket, err
}

func (s *WxMaJsapiServiceImpl) CreateJsapiSignature(url string) (*WxJsapiSignature, error) {
	jsapiTicket, _ := s.GetJsapiTicket()
	appId := s.service.GetWxMaConfig().GetAppID()
	return CreateJsapiSignature(url, appId, jsapiTicket.Ticket)
}
