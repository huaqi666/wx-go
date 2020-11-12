package ma

import (
	"github.com/cliod/wx-go/common"
)

type WxMaJsapiService interface {
	common.WxJsapi
	// 获得卡券api_ticket,不强制刷新api_ticket
	GetCardApiTicket() (*common.Ticket, error)
	// 获得卡券api_ticket
	// 获得时会检查apiToken是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	ForceGetCardApiTicket(forceRefresh bool) (*common.Ticket, error)
}

type WxMaJsapiServiceImpl struct {
	service WxMaService
}

func newWxMaJsapiService(service WxMaService) *WxMaJsapiServiceImpl {
	return &WxMaJsapiServiceImpl{
		service: service,
	}
}

func (s *WxMaJsapiServiceImpl) GetCardApiTicket() (*common.Ticket, error) {
	return s.ForceGetCardApiTicket(false)
}

func (s *WxMaJsapiServiceImpl) ForceGetCardApiTicket(forceRefresh bool) (*common.Ticket, error) {
	return s.ForceGetTicket(common.WxCard, forceRefresh)
}

func (s *WxMaJsapiServiceImpl) GetJsapiTicket() (*common.Ticket, error) {
	return s.ForceGetJsapiTicket(false)
}

func (s *WxMaJsapiServiceImpl) ForceGetJsapiTicket(forceRefresh bool) (*common.Ticket, error) {
	return s.ForceGetTicket(common.JSAPI, forceRefresh)
}

func (s *WxMaJsapiServiceImpl) GetTicket(ticketType common.TicketType) (*common.Ticket, error) {
	return s.ForceGetTicket(ticketType, false)
}

func (s *WxMaJsapiServiceImpl) ForceGetTicket(ticketType common.TicketType, forceRefresh bool) (*common.Ticket, error) {
	conf := s.service.GetWxMaConfig()
	b := conf.IsTicketExpired(ticketType)
	if forceRefresh || b {
		tt, err := s.getTicket(ticketType)
		conf.UpdateTicket(ticketType, tt)
		return tt, err
	}
	return conf.GetTicket(ticketType), nil
}

func (s *WxMaJsapiServiceImpl) getTicket(ticketType common.TicketType) (*common.Ticket, error) {
	var ticket common.Ticket

	err := s.service.GetFor(&ticket, common.MpGetTicketUrl, ticketType)
	return &ticket, err
}

func (s *WxMaJsapiServiceImpl) CreateJsapiSignature(url string) (*common.WxJsapiSignature, error) {
	jsapiTicket, _ := s.GetJsapiTicket()
	appId := s.service.GetWxMaConfig().GetAppID()
	return common.CreateJsapiSignature(url, appId, jsapiTicket.Ticket)
}
