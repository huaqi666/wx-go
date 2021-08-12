package common

import (
	"strconv"
	"strings"
	"time"
)

// WxConfig 微信配置
type WxConfig interface {
	// GetAppID 获取appId
	GetAppID() string
	// GetSecret 获取密钥
	GetSecret() string

	// GetAccessToken 获取access_token
	GetAccessToken() *AccessToken
	// SetAccessToken 设置access_token
	SetAccessToken(*AccessToken)

	// GetWxTicket 获取Ticket
	GetWxTicket() WxTicket
}

// WxAccessToken API授权
type WxAccessToken interface {
	// GetAccessToken 获取access_token
	GetAccessToken() (*AccessToken, error)
	// IsAccessTokenExpired access_token是否过期, ture：是
	IsAccessTokenExpired() bool
	// ExpireAccessToken 强制过期access_token
	ExpireAccessToken()
}

type WxTicket interface {
	// GetTicket 获取Ticket
	GetTicket(TicketType) *Ticket
	// UpdateTicket 更新Ticket
	UpdateTicket(TicketType, *Ticket)
	// IsTicketExpired Ticket是否过期
	IsTicketExpired(TicketType) bool
	// ExpireTicket 直接过期Ticket
	ExpireTicket(TicketType)
}

type WxJsapi interface {
	// GetJsapiTicket 获得jsapi_ticket,不强制刷新jsapi_ticket.
	GetJsapiTicket() (*Ticket, error)
	// ForceGetJsapiTicket 获得jsapi_ticket.
	//   获得时会检查jsapiToken是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	//   详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	ForceGetJsapiTicket(bool) (*Ticket, error)
	// GetTicket 获得ticket,不强制刷新ticket.
	GetTicket(TicketType) (*Ticket, error)
	// ForceGetTicket 获得时会检查 Token是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error)
	// CreateJsapiSignature
	//   创建调用jsapi时所需要的签名.
	//   详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	CreateJsapiSignature(url string) (*WxJsapiSignature, error)
}

// WxService wxAPI，获取accessToken，配置管理
type WxService interface {
	Service
	// GetAccessToken 获取access_token
	GetAccessToken() (*AccessToken, error)
	// ForceGetAccessToken 是否强制获取access_token
	ForceGetAccessToken(forceRefresh bool) (*AccessToken, error)

	// GetWxConfig 配置读取
	GetWxConfig() WxConfig
	// SetWxConfig 配置设置
	SetWxConfig(WxConfig)

	// SetHttpService 设置http请求方式
	SetHttpService(Service)
}

// WxServiceImpl http请求默认实现
type WxServiceImpl struct {
	// 配置
	config WxConfig
	http   Service
}

func (s *WxServiceImpl) Get(url string, args ...interface{}) ([]byte, error) {
	return s.http.Get(url, s.attachAccessToken(url, args)...)
}

func (s *WxServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	return s.http.Post(url, contentType, data, s.attachAccessToken(url, args)...)
}

func (s *WxServiceImpl) GetFor(v interface{}, url string, args ...interface{}) error {
	return s.http.GetFor(v, url, s.attachAccessToken(url, args)...)
}

func (s *WxServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	return s.http.PostFor(v, url, contentType, data, s.attachAccessToken(url, args)...)
}

func (s *WxServiceImpl) attachAccessToken(url string, args []interface{}) []interface{} {
	var params []interface{}
	if strings.Contains(url, ApiSuffix) {
		accessToken := ""
		at, err := s.GetAccessToken()
		if err == nil {
			accessToken = at.AccessToken
		}
		params = append(params, accessToken)
	}
	params = append(params, args...)
	return params
}

func (s *WxServiceImpl) getAccessToken() (*AccessToken, error) {
	var at AccessToken
	err := s.GetFor(&at, AccessTokenUrl, s.GetWxConfig().GetAppID(), s.GetWxConfig().GetSecret())
	at.Time = time.Now()
	return &at, err
}

func (s *WxServiceImpl) GetAccessToken() (*AccessToken, error) {
	return s.ForceGetAccessToken(false)
}

func (s *WxServiceImpl) ForceGetAccessToken(forceRefresh bool) (*AccessToken, error) {
	b := s.IsAccessTokenExpired()
	if b || forceRefresh {
		at, err := s.getAccessToken()
		s.GetWxConfig().SetAccessToken(at)
		return at, err
	}
	return s.GetWxConfig().GetAccessToken(), nil
}

func (s *WxServiceImpl) IsAccessTokenExpired() bool {
	c := s.GetWxConfig()
	tok := c.GetAccessToken()
	if tok == nil {
		// 过期
		return true
	}
	ei := strconv.FormatUint(tok.ExpiresIn, 10)
	m, _ := time.ParseDuration(ei + "s")
	return tok.Time.Add(m).Before(time.Now())
}

func (s *WxServiceImpl) ExpireAccessToken() {
	s.config.SetAccessToken(nil)
}

func (s *WxServiceImpl) GetWxConfig() WxConfig {
	return s.config
}

func (s *WxServiceImpl) SetWxConfig(config WxConfig) {
	s.config = config
}

func (s *WxServiceImpl) SetHttpService(service Service) {
	s.http = service
}

type WxTicketImpl struct {
	JsapiTicket  *Ticket
	SdkTicket    *Ticket
	WxCardTicket *Ticket
}

func (c *WxTicketImpl) GetTicket(ticketType TicketType) *Ticket {
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

func (c *WxTicketImpl) UpdateTicket(ticketType TicketType, ticket *Ticket) {
	switch ticketType {
	case JSAPI:
		c.JsapiTicket = ticket
	case SDK:
		c.SdkTicket = ticket
	case WxCard:
		c.WxCardTicket = ticket
	}
}

func (c *WxTicketImpl) IsTicketExpired(ticketType TicketType) bool {
	tt := c.GetTicket(ticketType)
	if tt == nil {
		// 过期
		return true
	}
	ei := strconv.FormatUint(tt.ExpiresIn, 10)
	m, _ := time.ParseDuration(ei + "s")
	return tt.Time.Add(m).Before(time.Now())
}

func (c *WxTicketImpl) ExpireTicket(ticketType TicketType) {
	c.UpdateTicket(ticketType, nil)
}

func NewWxService(config WxConfig) WxService {
	return &WxServiceImpl{
		config: config,
		http:   NewService(),
	}
}
