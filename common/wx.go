package common

import (
	"strconv"
	"time"
)

// 配置的接口
type WxConfig interface {
	// 获取appId
	GetAppID() string
	// 获取密钥
	GetSecret() string
	// 获取access_token
	GetAccessToken() *AccessToken
	// 设置access_token
	SetAccessToken(at *AccessToken)
}

type WxAccessToken interface {
	// 获取access_token
	GetAccessToken() (*AccessToken, error)
	// access_token是否过期, ture：是
	IsAccessTokenExpired() bool
	// 强制过期access_token
	ExpireAccessToken()
}

// http请求，获取accessToken
type WxService interface {
	Service
	// 获取access_token
	GetAccessToken() (*AccessToken, error)
	// 是否强制获取access_token
	ForceGetAccessToken(forceRefresh bool) (*AccessToken, error)

	// 配置读取
	GetWxConfig() WxConfig
	// 配置设置
	SetWxConfig(WxConfig)

	// 设置http请求方式
	SetHttpService(Service)
}

// http请求默认实现
type WxServiceImpl struct {
	// 配置
	config WxConfig
	http   Service
}

func (s *WxServiceImpl) Get(url string, args ...interface{}) ([]byte, error) {
	return s.http.Get(url, args...)
}

func (s *WxServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	return s.http.Post(url, contentType, data, args...)
}

func (s *WxServiceImpl) GetFor(v interface{}, url string, args ...interface{}) error {
	return s.http.GetFor(v, url, args...)
}

func (s *WxServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	return s.http.PostFor(v, url, contentType, data, args...)
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

func (s *WxServiceImpl) GetWxConfig() WxConfig {
	return s.config
}

func (s *WxServiceImpl) SetWxConfig(config WxConfig) {
	s.config = config
}

func (s *WxServiceImpl) SetHttpService(service Service) {
	s.http = service
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
