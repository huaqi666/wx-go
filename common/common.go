package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// http请求
type Service interface {
	// 执行Get请求
	Get(url string, args ...interface{}) ([]byte, error)
	// 执行Post请求
	Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error)

	// Get 执行Get请求并将结果转成对象
	GetFor(v interface{}, url string, args ...interface{}) error
	// Post 执行Post请求并将结果转成对象
	PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error
}

// http请求，获取accessToken
type WxService interface {
	Service
	// 获取access_token
	GetAccessToken() (*AccessToken, error)
	// 是否强制获取access_token
	ForceGetAccessToken(forceRefresh bool) (*AccessToken, error)

	GetWxConfig() WxConfig
	SetWxConfig(WxConfig)
}

// http请求默认实现
type ServiceImpl struct {
}

// http请求默认实现
type WxServiceImpl struct {
	ServiceImpl
	// 配置
	config WxConfig
}

func (s *ServiceImpl) Get(url string, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	body, err := json.Marshal(data)
	res, err := http.Post(uri, contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) GetFor(v interface{}, url string, args ...interface{}) error {
	res, err := s.Get(url, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, v)
}

func (s *ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	res, err := s.Post(url, contentType, data, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, v)
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

func (s WxServiceImpl) SetWxConfig(config WxConfig) {
	s.config = config
}
