package mp

import (
	"strconv"
	"time"
	"wx-go/common"
)

type WxMpConfig interface {
	common.WxConfig

	GetToken() string
	GetAesKey() string
	GetTicket(TicketType) *Ticket
	UpdateTicket(TicketType, *Ticket)
	IsTicketExpired(TicketType) bool
	ExpireTicket(TicketType)
}

type WxMpConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token        string
	AesKey       string
	JsapiTicket  *Ticket
	SdkTicket    *Ticket
	WxCardTicket *Ticket
}

func newWxMpConfig(appId, secret string) WxMpConfig {
	return &WxMpConfigImpl{
		appId:  appId,
		secret: secret,
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

func (c *WxMpConfigImpl) GetTicket(ticketType TicketType) *Ticket {
	switch ticketType {
	case JSAPI:
		return c.JsapiTicket
	case SDK:
		return c.SdkTicket
	case WxCard:
		return c.WxCardTicket
	}
	return c.JsapiTicket
}

func (c *WxMpConfigImpl) UpdateTicket(ticketType TicketType, ticket *Ticket) {
	switch ticketType {
	case JSAPI:
		c.JsapiTicket = ticket
	case SDK:
		c.SdkTicket = ticket
	case WxCard:
		c.WxCardTicket = ticket
	}
}

func (c *WxMpConfigImpl) IsTicketExpired(ticketType TicketType) bool {
	tt := c.GetTicket(ticketType)
	if tt == nil {
		// 过期
		return true
	}
	ei := strconv.FormatUint(tt.ExpiresIn, 10)
	m, _ := time.ParseDuration(ei + "s")
	return tt.Time.Add(m).Before(time.Now())
}

func (c *WxMpConfigImpl) ExpireTicket(ticketType TicketType) {
	c.UpdateTicket(ticketType, nil)
}
