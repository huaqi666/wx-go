package ma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	get(v interface{}, url string, args ...interface{}) error
	post(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error
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
	err := s.get(&at, AccessTokenUrl, s.config.GetAppID(), s.config.GetSecret())
	at.Time = time.Now()
	return &at, err
}

func (s *ServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.get(&jsr, SessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *ServiceImpl) GetUserService() UserService {
	return s.userService
}

func (s *ServiceImpl) GetQrCodeService() QrCodeService {
	return s.qrCodeService
}

func (s *ServiceImpl) get(v interface{}, url string, args ...interface{}) error {
	uri := fmt.Sprintf(url, args)
	res, err := http.Get(uri)
	if err != nil {
		return err
	}
	return s.convert(v, res.Body)
}

func (s *ServiceImpl) post(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	uri := fmt.Sprintf(url, args)
	body, err := json.Marshal(data)
	res, err := http.Post(uri, contentType, bytes.NewReader(body))
	if err != nil {
		return err
	}
	return s.convert(v, res.Body)
}

func (s ServiceImpl) convert(v interface{}, reader io.Reader) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	switch v.(type) {
	case []byte:
		v = &body
		return nil
	case io.Reader:
		v = reader
		return nil
	case struct{}:
		err = json.Unmarshal(body, v)
		return err
	default:
		return fmt.Errorf("不是 %s 或者 %s 类型", "struct{}", "[]byte")
	}
}

func NewService(appId, secret string) Service {
	return newService(appId, secret)
}
