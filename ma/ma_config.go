package ma

import (
	"github.com/cliod/wx-go/common"
)

type WxMaConfig interface {
	common.WxConfig
	common.WxTicket

	GetToken() string
	GetAesKey() string
	GetMsgDataFormat() string
}

type WxMaConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token         string
	AesKey        string
	MsgDataFormat string

	ticket common.WxTicket
}

func newWxMaConfig(appId, secret string) *WxMaConfigImpl {
	return &WxMaConfigImpl{
		appId:  appId,
		secret: secret,
		ticket: new(common.WxTicketImpl),
	}
}

func (c *WxMaConfigImpl) GetAppID() string {
	return c.appId
}

func (c *WxMaConfigImpl) GetSecret() string {
	return c.secret
}

func (c *WxMaConfigImpl) GetAccessToken() *common.AccessToken {
	return c.AccessToken
}

func (c *WxMaConfigImpl) SetAccessToken(at *common.AccessToken) {
	c.AccessToken = at
}

func (c *WxMaConfigImpl) GetToken() string {
	return c.Token
}

func (c *WxMaConfigImpl) GetAesKey() string {
	return c.AesKey
}

func (c *WxMaConfigImpl) GetMsgDataFormat() string {
	return c.MsgDataFormat
}

func (c WxMaConfigImpl) GetWxTicket() common.WxTicket {
	return c.ticket
}

func (c WxMaConfigImpl) GetTicket(ticketType common.TicketType) *common.Ticket {
	return c.ticket.GetTicket(ticketType)
}

func (c *WxMaConfigImpl) UpdateTicket(ticketType common.TicketType, ticket *common.Ticket) {
	c.ticket.UpdateTicket(ticketType, ticket)
}

func (c *WxMaConfigImpl) IsTicketExpired(ticketType common.TicketType) bool {
	return c.ticket.IsTicketExpired(ticketType)
}

func (c *WxMaConfigImpl) ExpireTicket(ticketType common.TicketType) {
	c.UpdateTicket(ticketType, nil)
}

func NewWxMaConfig(appId, secret string) WxMaConfig {
	return newWxMaConfig(appId, secret)
}
