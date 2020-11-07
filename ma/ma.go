package ma

import (
	"github.com/cliod/wx-go/common"
)

type WxMaService interface {
	common.WxService
	// 携带access_token执行
	Do(url string, res error) error
	// 获取配置
	GetWxMaConfig() WxMaConfig
	// 设置配置
	SetWxMaConfig(config WxMaConfig)

	// jsCode换取openid
	JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error)
	// 验证消息的确来自微信服务器
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421135319&token=&lang=zh_CN
	CheckSignature(timestamp, nonce, signature string) bool
	/* 用户支付完成后，获取该用户的 UnionId，无需用户授权。本接口支持第三方平台代理查询。
	   注意：调用前需要用户完成支付，且在支付后的五分钟内有效。
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/api/getPaidUnionId.html */
	GetPaidUnionId(openid, transactionId, mchId, outTradeNo string) (*WxMaUnionIdResult, error)

	// 获取用户接口
	GetWxMaUserService() WxMaUserService
	// 设置(用户自定义的)用户接口
	SetWxMaUserService(userService WxMaUserService)
	// 获取二维码接口
	GetWxMaQrcodeService() WxMaQrcodeService
	// 获取订阅接口
	GetWxMaSubscribeService() WxMaSubscribeService
	// 设置(用户自定义的)二维码接口
	SetWxMaQrcodeService(qrcodeService WxMaQrcodeService)
	// 设置(用户自定义的)订阅接口
	SetWxMaSubscribeService(subscribeService WxMaSubscribeService)
}

type WxMaServiceImpl struct {
	common.WxServiceImpl

	config           WxMaConfig
	userService      WxMaUserService
	qrCodeService    WxMaQrcodeService
	subscribeService WxMaSubscribeService
}

func newWxMaService(config WxMaConfig) *WxMaServiceImpl {
	impl := WxMaServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMaConfig(config)
	impl.userService = newWxMaUserService(&impl)
	impl.qrCodeService = newWxMaQrcodeService(&impl)
	impl.subscribeService = newWxMaSubscribeService(&impl)
	return &impl
}

func (s *WxMaServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.GetFor(&jsr, common.MaSessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *WxMaServiceImpl) CheckSignature(timestamp, nonce, signature string) bool {
	return CheckSignature(s.GetWxMaConfig().GetToken(), timestamp, nonce, signature)
}

func (s *WxMaServiceImpl) GetPaidUnionId(openid, transactionId, mchId, outTradeNo string) (*WxMaUnionIdResult, error) {
	param := map[string]string{}
	param["openid"] = openid
	if transactionId != "" {
		param["transaction_id"] = transactionId
	}
	if mchId != "" {
		param["mch_id"] = mchId
	}
	if outTradeNo != "" {
		param["out_trade_no"] = outTradeNo
	}

	url := common.MaGetPaidUnionIdUrl
	for k, v := range param {
		url += "&" + k + "=" + v
	}
	var res WxMaUnionIdResult
	return &res, s.Do(url, &res)
}

func (s *WxMaServiceImpl) GetWxMaUserService() WxMaUserService {
	return s.userService
}

func (s *WxMaServiceImpl) GetWxMaQrcodeService() WxMaQrcodeService {
	return s.qrCodeService
}

func (s *WxMaServiceImpl) GetWxMaSubscribeService() WxMaSubscribeService {
	return s.subscribeService
}

func (s *WxMaServiceImpl) SetWxMaUserService(userService WxMaUserService) {
	s.userService = userService
}

func (s *WxMaServiceImpl) SetWxMaQrcodeService(qrcodeService WxMaQrcodeService) {
	s.qrCodeService = qrcodeService
}

func (s *WxMaServiceImpl) SetWxMaSubscribeService(subscribeService WxMaSubscribeService) {
	s.subscribeService = subscribeService
}

func (s *WxMaServiceImpl) GetWxMaConfig() WxMaConfig {
	return s.config
}

func (s *WxMaServiceImpl) SetWxMaConfig(config WxMaConfig) {
	s.SetWxConfig(config)
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func (s *WxMaServiceImpl) Do(url string, res error) (err error) {
	at, err := s.GetAccessToken()
	if err == nil {
		err = s.GetFor(&res, url, at.AccessToken)
	}
	return
}

func NewWxMaService(appId, secret string) WxMaService {
	return newWxMaService(newWxMaConfig(appId, secret))
}

func NewWxMaServiceBy(config WxMaConfig) WxMaService {
	if config == nil {
		config = new(WxMaConfigImpl)
	}
	return newWxMaService(config)
}
