package mp

import (
	"fmt"
	"strconv"
	"time"
	"wx-go/common"
	"wx-go/common/util"
)

type WxMpService interface {
	common.WxService

	// 获取配置
	GetWxMpConfig() WxMpConfig
	// 设置配置
	SetWxMpConfig(WxMpConfig)

	// 验证消息的确来自微信服务器
	CheckSignature(timestamp, nonce, signature string) bool

	// 获得jsapi_ticket,不强制刷新jsapi_ticket.
	GetJsapiTicket() (*Ticket, error)
	// 获得jsapi_ticket.
	// 获得时会检查jsapiToken是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	// 详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	ForceGetJsapiTicket(bool) (*Ticket, error)

	// 获得ticket,不强制刷新ticket.
	GetTicket(ticketType TicketType) (*Ticket, error)
	// 获得时会检查 Token是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error)

	// 创建调用jsapi时所需要的签名.
	// 详情请见：http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115&token=&lang=zh_CN
	CreateJsapiSignature(url string) (*WxJsapiSignature, error)

	// 获取用户接口
	GetWxMpUserService() WxMpUserService
	// 获取二维码接口
	GetWxMpQrcodeService() WxMpQrcodeService
}

// jspai signature.
type WxJsapiSignature struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}

type WxMpServiceImpl struct {
	common.WxServiceImpl

	config        WxMpConfig
	userService   WxMpUserService
	qrCodeService WxMpQrcodeService
}

func newWxMpService(appId, secret string) *WxMpServiceImpl {
	impl := WxMpServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMpConfig(newWxMpConfig(appId, secret))
	impl.userService = newWxMpUserService(&impl)
	impl.qrCodeService = newWxMpQrcodeService(&impl)
	return &impl
}

func (s *WxMpServiceImpl) CheckSignature(timestamp, nonce, signature string) bool {
	arr := []string{s.GetWxMpConfig().GetToken(), timestamp, nonce}
	gen, err := util.Gen(arr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return gen == signature
}

func (s *WxMpServiceImpl) GetJsapiTicket() (*Ticket, error) {
	return s.ForceGetJsapiTicket(false)
}

func (s *WxMpServiceImpl) ForceGetJsapiTicket(forceRefresh bool) (*Ticket, error) {
	return s.ForceGetTicket(JSAPI, forceRefresh)
}

func (s *WxMpServiceImpl) GetTicket(ticketType TicketType) (*Ticket, error) {
	return s.ForceGetTicket(ticketType, false)
}

func (s *WxMpServiceImpl) ForceGetTicket(ticketType TicketType, forceRefresh bool) (*Ticket, error) {
	b := s.GetWxMpConfig().IsTicketExpired(ticketType)
	if forceRefresh || b {
		tt, err := s.getTicket(ticketType)
		s.GetWxMpConfig().UpdateTicket(ticketType, tt)
		return tt, err
	}
	return s.GetWxMpConfig().GetTicket(ticketType), nil
}

func (s *WxMpServiceImpl) getTicket(ticketType TicketType) (*Ticket, error) {
	at, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	var ticket Ticket

	err = s.GetFor(&ticket, common.MpGetTicketUrl, at.AccessToken, ticketType)
	return &ticket, err
}

func (s *WxMpServiceImpl) CreateJsapiSignature(url string) (*WxJsapiSignature, error) {
	timestamp := strconv.Itoa(time.Now().Second())
	randomStr := util.RandSeq(16)
	jsapiTicket, _ := s.GetJsapiTicket()
	arr := []string{"jsapi_ticket=" + jsapiTicket.Ticket, "noncestr=" + randomStr, "timestamp=" + timestamp, "url=" + url}
	signature, err := util.GenWithAmple(arr)
	if err != nil {
		return nil, err
	}
	return &WxJsapiSignature{
		AppId:     s.GetWxMpConfig().GetAppID(),
		Timestamp: timestamp,
		NonceStr:  randomStr,
		Url:       url,
		Signature: signature,
	}, nil
}

func (s *WxMpServiceImpl) GetWxMpUserService() WxMpUserService {
	return s.userService
}

func (s *WxMpServiceImpl) GetWxMpQrcodeService() WxMpQrcodeService {
	return s.qrCodeService
}

func (s *WxMpServiceImpl) GetWxMpConfig() WxMpConfig {
	return s.config
}

func (s *WxMpServiceImpl) SetWxMpConfig(config WxMpConfig) {
	s.SetWxConfig(config)
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func NewWxMpService(appId, secret string) WxMpService {
	return newWxMpService(appId, secret)
}
