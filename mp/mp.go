package mp

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
	"wx-go/common"
)

type WxMpService interface {
	common.WxService

	// 获取配置
	GetWxMpConfig() WxMpConfig
	// 设置配置
	SetWxMpConfig(WxMpConfig)

	// 验证消息的确来自微信服务器
	CheckSignature(timestamp, nonce, signature string) bool

	// 获得ticket,不强制刷新ticket.
	//GetTicket(ticketType TicketType) error
	// 获得时会检查 Token是否过期，如果过期了，那么就刷新一下，否则就什么都不干
	//ForceGetTicket(ticketType TicketType, forceRefresh bool) error

	// 获取用户接口
	GetWxMpUserService() WxMpUserService
	// 获取二维码接口
	GetWxMpQrcodeService() WxMpQrcodeService
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
	sort.Strings(arr)
	var str string
	for _, v := range arr {
		str += v
	}
	sum := sha1.Sum([]byte(str))
	return strings.ToLower(hex.EncodeToString(sum[:])) == signature
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
