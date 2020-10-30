package common

import (
	"strconv"
	"time"
)

// http请求，获取accessToken
type WxService interface {
	Service
	// 获取access_token
	GetAccessToken() (*AccessToken, error)
	// 是否强制获取access_token
	ForceGetAccessToken(forceRefresh bool) (*AccessToken, error)

	GetWxConfig() WxConfig
	SetWxConfig(WxConfig)

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
	c := s.GetWxConfig()
	tok := c.GetAccessToken()
	if !forceRefresh && tok != nil {
		s := strconv.FormatUint(tok.ExpiresIn, 10)
		m, _ := time.ParseDuration(s + "s")
		if tok.Time.Add(m).After(time.Now()) {
			return tok, nil
		}
	}
	at, err := s.getAccessToken()
	s.GetWxConfig().SetAccessToken(at)
	return at, err
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
