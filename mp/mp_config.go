package mp

import (
	"github.com/cliod/wx-go/common"
)

type WxMpConfig interface {
	common.WxConfig
	common.WxTicket

	GetToken() string
	GetAesKey() string
}

type WxMpConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token  string
	AesKey string
	ticket common.WxTicket
}

func newWxMpConfig(appId, secret string) WxMpConfig {
	return &WxMpConfigImpl{
		appId:  appId,
		secret: secret,
		ticket: new(common.WxTicketImpl),
	}
}

func (c *WxMpConfigImpl) GetAppID() string {
	return c.appId
}

func (c *WxMpConfigImpl) GetSecret() string {
	return c.secret
}

func (c *WxMpConfigImpl) GetAccessToken() *common.AccessToken {
	return c.AccessToken
}

func (c *WxMpConfigImpl) SetAccessToken(at *common.AccessToken) {
	c.AccessToken = at
}

func (c *WxMpConfigImpl) GetToken() string {
	return c.Token
}

func (c *WxMpConfigImpl) GetAesKey() string {
	return c.AesKey
}

func (c *WxMpConfigImpl) GetWxTicket() common.WxTicket {
	return c.ticket
}

func (c WxMpConfigImpl) GetTicket(ticketType common.TicketType) *common.Ticket {
	return c.ticket.GetTicket(ticketType)
}

func (c *WxMpConfigImpl) UpdateTicket(ticketType common.TicketType, ticket *common.Ticket) {
	c.ticket.UpdateTicket(ticketType, ticket)
}

func (c *WxMpConfigImpl) IsTicketExpired(ticketType common.TicketType) bool {
	return c.ticket.IsTicketExpired(ticketType)
}

func (c *WxMpConfigImpl) ExpireTicket(ticketType common.TicketType) {
	c.UpdateTicket(ticketType, nil)
}

func NewWxMpConfig(appId, secret string) WxMpConfig {
	return newWxMpConfig(appId, secret)
}
