package ma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Service interface {
	// 获取access_token
	GetAccessToken() (*AccessToken, error)
	ForceGetAccessToken(forceRefresh bool) (*AccessToken, error)
	// jsCode换取openid
	JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error)

	// 获取用户接口
	GetUserService() UserService
	// 获取二维码接口
	GetQrCodeService() QrCodeService

	// 执行http请求
	get(url string, args ...interface{}) ([]byte, error)
	post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error)
}

type ServiceImpl struct {
	config        Config
	userService   UserService
	qrCodeService QrCodeService
}

func newService(appId, secret string) Service {
	impl := ServiceImpl{
		config: newConfig(appId, secret),
	}
	at, _ := impl.getAccessToken()
	impl.config.SetAccessToken(at)
	impl.userService = newUserService(&impl)
	impl.qrCodeService = newQrCodeService(&impl)
	return &impl
}

func (s *ServiceImpl) GetAccessToken() (*AccessToken, error) {
	return s.ForceGetAccessToken(false)
}

func (s *ServiceImpl) ForceGetAccessToken(forceRefresh bool) (*AccessToken, error) {
	c := s.config
	tok := c.GetAccessToken()
	if !forceRefresh && tok != nil {
		s := strconv.FormatUint(tok.ExpiresIn, 10)
		m, _ := time.ParseDuration(s + "s")
		if tok.Time.Add(m).After(time.Now()) {
			return tok, nil
		}
	}
	at, err := s.getAccessToken()
	s.config.SetAccessToken(at)
	return at, err
}

func (s *ServiceImpl) getAccessToken() (*AccessToken, error) {
	var at AccessToken
	err := s.getFor(&at, AccessTokenUrl, s.config.GetAppID(), s.config.GetSecret())
	at.Time = time.Now()
	return &at, err
}

func (s *ServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.getFor(&jsr, SessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *ServiceImpl) GetUserService() UserService {
	return s.userService
}

func (s *ServiceImpl) GetQrCodeService() QrCodeService {
	return s.qrCodeService
}

func (s *ServiceImpl) get(url string, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	body, err := json.Marshal(data)
	res, err := http.Post(uri, contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (s *ServiceImpl) getFor(v interface{}, url string, args ...interface{}) error {
	res, err := s.get(url, args...)
	if err != nil {
		return err
	}
	return s.convert(v, res)
}

func (s *ServiceImpl) postFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	res, err := s.post(url, contentType, data, args...)
	if err != nil {
		return err
	}
	return s.convert(v, res)
}

func (s *ServiceImpl) convert(v interface{}, body []byte) error {
	if v != nil && reflect.ValueOf(v).Kind() == reflect.Ptr {
		err := json.Unmarshal(body, v)
		return err
	}
	return fmt.Errorf("对象不是结构体")
}

func NewService(appId, secret string) Service {
	return newService(appId, secret)
}
