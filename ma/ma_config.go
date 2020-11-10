package ma

import (
	"github.com/cliod/wx-go/common"
	"strconv"
	"time"
)

type WxMaConfig interface {
	common.WxConfig

	GetToken() string
	GetAesKey() string
	GetMsgDataFormat() string

	GetTicket(TicketType) *Ticket
	UpdateTicket(TicketType, *Ticket)
	IsTicketExpired(TicketType) bool
	ExpireTicket(TicketType)
}

type WxMaConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token         string
	AesKey        string
	MsgDataFormat string

	JsapiTicket  *Ticket
	SdkTicket    *Ticket
	WxCardTicket *Ticket
}

func newWxMaConfig(appId, secret string) *WxMaConfigImpl {
	return &WxMaConfigImpl{
		appId:  appId,
		secret: secret,
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

func (c *WxMaConfigImpl) GetTicket(ticketType TicketType) *Ticket {
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

func (c *WxMaConfigImpl) UpdateTicket(ticketType TicketType, ticket *Ticket) {
	switch ticketType {
	case JSAPI:
		c.JsapiTicket = ticket
	case SDK:
		c.SdkTicket = ticket
	case WxCard:
		c.WxCardTicket = ticket
	}
}

func (c *WxMaConfigImpl) IsTicketExpired(ticketType TicketType) bool {
	tt := c.GetTicket(ticketType)
	if tt == nil {
		// 过期
		return true
	}
	ei := strconv.FormatUint(tt.ExpiresIn, 10)
	m, _ := time.ParseDuration(ei + "s")
	return tt.Time.Add(m).Before(time.Now())
}

func (c *WxMaConfigImpl) ExpireTicket(ticketType TicketType) {
	c.UpdateTicket(ticketType, nil)
}

func NewWxMaConfig(appId, secret string) WxMaConfig {
	return newWxMaConfig(appId, secret)
}
