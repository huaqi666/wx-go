package pay

import (
	"github.com/cliod/wx-go/common"
)

type WxPayConfig struct {
	PayBaseUrl string `json:"pay_base_url"` // 微信支付接口请求地址域名部分.

	AppId     string    `json:"app_id"`      // 公众号/小程序 appid
	MchId     string    `json:"mch_id"`      // 微信支付商户号.
	MchKey    string    `json:"mch_key"`     // 微信支付商户密钥.
	NotifyUrl string    `json:"notify_url"`  // 微信支付异步回掉地址，通知url必须为直接可访问的url，不能携带参数.
	KeyPath   string    `json:"key_path"`    // apiclient_cert.p12文件的绝对路径，或者如果放在项目中，请以classpath:开头指定.
	SubAppId  string    `json:"sub_app_id"`  // 服务商模式下的子商户公众账号ID，普通模式请不要配置，请在配置文件中将对应项删除.
	SubMchId  string    `json:"sub_mch_id"`  // 服务商模式下的子商户号，普通模式请不要配置，最好是请在配置文件中将对应项删除.
	TradeType TradeType `json:"trade_type"`  // 交易类型.
	SignType  SignType  `json:"sign_type"`   // 签名方式.
	EntPayKey string    `json:"ent_pay_key"` // 企业支付密钥.

	PrivateKeyPath  string `json:"private_key_path"`  // apiv3 商户apiclient_key.pem
	PrivateCertPath string `json:"private_cert_path"` // apiv3 商户apiclient_cert.pem

	ApiV3Key           string `json:"api_v_3_key"`           // apiV3 秘钥值.
	CertSerialNo       string `json:"cert_serial_no"`        // apiV3 证书序列号值
	ServiceId          string `json:"service_id"`            // 微信支付分serviceId
	PayScoreNotifyUrl  string `json:"pay_score_notify_url"`  // 微信支付分回调地址
	PrivateKey         string `json:"private_key"`           // 私钥信息
	CertAutoUpdateTime int    `json:"cert_auto_update_time"` // 证书自动更新时间差(分钟)，默认一分钟

	HttpConnectionTimeout int `json:"http_connection_timeout"` // http请求连接超时时间 5000
	HttpTimeout           int `json:"http_timeout"`            // http请求数据读取等待时间 10000

	UseSandboxEnv     bool   `json:"use_sandbox_env"`
	IfSaveApiData     bool   `json:"if_save_api_data"`
	HttpProxyHost     string `json:"http_proxy_host"`
	HttpProxyPort     string `json:"http_proxy_port"`
	HttpProxyUsername string `json:"http_proxy_username"`
	HttpProxyPassword string `json:"http_proxy_password"`
}

func (c *WxPayConfig) GetPayBaseUrl() string {
	if c.PayBaseUrl == "" {
		c.PayBaseUrl = common.PayDefaultPayBaseUrl
	}
	return c.PayBaseUrl
}

func newWxPayV2Config(appId, mchId, mchKey, notifyUrl, keyPath string) *WxPayConfig {
	return &WxPayConfig{
		AppId:     appId,
		MchId:     mchId,
		MchKey:    mchKey,
		NotifyUrl: notifyUrl,
		KeyPath:   keyPath,
		TradeType: JSAPI,
		SignType:  MD5,

		PayBaseUrl: common.PayDefaultPayBaseUrl,
	}
}

func NewWxPayV2Config(appId, mchId, mchKey, notifyUrl, keyPath string) *WxPayConfig {
	return newWxPayV2Config(appId, mchId, mchKey, notifyUrl, keyPath)
}
