package ma

import (
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
)

type WxMaService interface {
	common.WxService
	// 获取配置
	GetWxMaConfig() WxMaConfig
	// 设置配置
	SetWxMaConfig(WxMaConfig)

	// jsCode换取openid
	JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error)
	// 验证消息的确来自微信服务器
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421135319&token=&lang=zh_CN
	CheckSignature(timestamp, nonce, signature string) bool
	/* 用户支付完成后，获取该用户的UnionId，无需用户授权。本接口支持第三方平台代理查询。
	   注意：调用前需要用户完成支付，且在支付后的五分钟内有效。
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/api/getPaidUnionId.html */
	GetPaidUnionId(openid, transactionId, mchId, outTradeNo string) (*WxMaUnionIdResult, error)

	// 获取用户接口
	GetWxMaUserService() WxMaUserService
	// 设置(用户自定义的)用户接口
	SetWxMaUserService(WxMaUserService)
	// 获取二维码接口
	GetWxMaQrcodeService() WxMaQrcodeService
	// 获取订阅接口
	GetWxMaSubscribeService() WxMaSubscribeService
	// 获取分享接口
	GetWxMaShareService() WxMaShareService
	// 获取消息接口
	GetWxMaMessageService() WxMaMsgService
	// 获取设置jsapi接口
	GetWxMaJsapiService() WxMaJsapiService
	// 设置(用户自定义的)二维码接口
	SetWxMaQrcodeService(WxMaQrcodeService)
	// 设置(用户自定义的)订阅接口
	SetWxMaSubscribeService(WxMaSubscribeService)
	// 设置(用户自定义的)分享接口
	SetWxMaShareService(WxMaShareService)
	// 设置(用户自定义的)消息接口
	SetWxMaMsgService(WxMaMsgService)
	// 设置jsapi接口
	SetWxMaJsapiService(WxMaJsapiService)
}

type WxMaServiceImpl struct {
	common.WxServiceImpl

	config           WxMaConfig
	userService      WxMaUserService
	qrCodeService    WxMaQrcodeService
	subscribeService WxMaSubscribeService
	shareService     WxMaShareService
	msgService       WxMaMsgService
	liveService      WxMaLiveService
	jsapiService     WxMaJsapiService
}

func newWxMaService(config WxMaConfig) *WxMaServiceImpl {
	impl := &WxMaServiceImpl{}
	impl.SetHttpService(common.NewService())
	impl.SetWxMaConfig(config)
	impl.userService = newWxMaUserService(impl)
	impl.qrCodeService = newWxMaQrcodeService(impl)
	impl.subscribeService = newWxMaSubscribeService(impl)
	impl.msgService = newWxMaMsgService(impl)
	impl.shareService = newWxMaShareService(impl)
	impl.liveService = newWxMaLiveService(impl)
	impl.jsapiService = newWxMaJsapiService(impl)
	return impl
}

func (s *WxMaServiceImpl) JsCode2SessionInfo(jsCode string) (*JsCode2SessionResult, error) {
	var jsr JsCode2SessionResult
	err := s.GetFor(&jsr, common.MaSessionInfoUrl, s.config.GetAppID(), s.config.GetSecret(), jsCode)
	return &jsr, err
}

func (s *WxMaServiceImpl) CheckSignature(timestamp, nonce, signature string) bool {
	return util.CheckSignature(s.GetWxMaConfig().GetToken(), timestamp, nonce, signature)
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
	return &res, s.GetFor(&res, url)
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

func (s *WxMaServiceImpl) GetWxMaShareService() WxMaShareService {
	return s.shareService
}

func (s *WxMaServiceImpl) GetWxMaMessageService() WxMaMsgService {
	return s.msgService
}

func (s *WxMaServiceImpl) GetWxMaJsapiService() WxMaJsapiService {
	return s.jsapiService
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

func (s *WxMaServiceImpl) SetWxMaShareService(shareService WxMaShareService) {
	s.shareService = shareService
}

func (s *WxMaServiceImpl) SetWxMaMsgService(msgService WxMaMsgService) {
	s.msgService = msgService
}

func (s *WxMaServiceImpl) SetWxMaJsapiService(jsapiService WxMaJsapiService) {
	s.jsapiService = jsapiService
}

func (s *WxMaServiceImpl) GetWxMaConfig() WxMaConfig {
	return s.config
}

func (s *WxMaServiceImpl) SetWxMaConfig(config WxMaConfig) {
	s.SetWxConfig(config)
	s.config = config
	_, _ = s.ForceGetAccessToken(true)
}

func NewWxMaServiceBy(appId, secret string) WxMaService {
	return newWxMaService(newWxMaConfig(appId, secret))
}

func NewWxMaService(config WxMaConfig) WxMaService {
	if config == nil {
		config = new(WxMaConfigImpl)
	}
	return newWxMaService(config)
}
