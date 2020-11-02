package pay

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

type WxPayService interface {
	// 执行Post请求
	Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error)
	// Post 执行Post请求并将结果转成对象
	PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error
	// Post 携带证书执行Post请求
	PostKey(url string, data interface{}, certPath, keyPath string, args ...interface{}) ([]byte, error)
	// Post 携带证书执行Post请求并将结果转成对象
	PostKeyFor(v interface{}, url string, data interface{}, certPath, keyPath string, args ...interface{}) error

	// 获取微信支付请求url前缀，沙箱环境可能不一样.
	GetPayBaseUr() string

	// 统一支付
	// 调用统一下单接口，并组装生成支付所需参数对象.
	UnifyPay(WxPayUnifiedOrderRequest) ([]byte, error)
	/** 统一下单
	  (详见https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1)
	  在发起微信支付前，需要调用统一下单接口，获取"预支付交易会话标识"
	  接口地址：https://api.mch.weixin.qq.com/pay/unifiedorder */
	UnifyOrder(WxPayUnifiedOrderRequest) (*WxPayUnifiedOrderResult, error)

	CloseOrder(string) (*WxPayOrderCloseResult, error)

	QueryOrder(string) (*WxPayOrderQueryResult, error)

	// 获取配置
	GetWxPayConfig() *WxPayConfig
	// 设置配置
	SetWxPayConfig(*WxPayConfig)

	// 获取提现接口
	GetEntPayService() WxEntPayService
}

type WxPayV2ServiceImpl struct {
	http          common.Service
	config        *WxPayConfig
	entPayService WxEntPayService
}

func newWxPayService(config *WxPayConfig) *WxPayV2ServiceImpl {
	impl := &WxPayV2ServiceImpl{}
	impl.http = common.NewXmlService()
	impl.SetWxPayConfig(config)
	impl.entPayService = newWxEntPayService(impl)
	return impl
}

func (p *WxPayV2ServiceImpl) Post(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	return p.http.Post(url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	return p.http.PostFor(v, url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostKey(url string, data interface{}, certPath, keyPath string, args ...interface{}) ([]byte, error) {
	cli, err := util.NewTLSClient(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	ser := common.NewXmlServiceFor(cli)
	return ser.Post(url, "application/xml; charset=utf-8", data, args...)
}

func (p *WxPayV2ServiceImpl) PostKeyFor(v interface{}, url string, data interface{}, certPath, keyPath string, args ...interface{}) error {
	b, err := p.PostKey(url, data, certPath, keyPath, args...)
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, v)
}

func (p *WxPayV2ServiceImpl) GetPayBaseUr() string {
	url := p.GetWxPayConfig().PayBaseUrl
	if p.GetWxPayConfig().UseSandboxEnv {
		url += "/sandboxnew"
	}
	return url
}

func (p *WxPayV2ServiceImpl) UnifyPay(request WxPayUnifiedOrderRequest) ([]byte, error) {
	v, err := p.UnifyOrder(request)
	if err != nil {
		return nil, err
	}
	prepayId := v.PrepayId
	if prepayId == "" {
		return nil, fmt.Errorf("无法获取prepay id，错误代码： '%s'，信息：%s。", v.ErrCode, v.ErrCodeDes)
	}

	timestamp := strconv.Itoa(time.Now().Second())
	nonceStr := v.NonceStr

	switch request.TradeType {
	case NATIVE:
		return []byte(v.MwebUrl), nil
	case H5:
		return []byte(v.CodeURL), nil
	case APP:
		// APP支付绑定的是微信开放平台上的账号，APPID为开放平台上绑定APP后发放的参数
		appId := v.AppId
		if v.SubAppId != "" {
			appId = v.SubAppId
		}
		// 此map用于参与调起sdk支付的二次签名,格式全小写，timestamp只能是10位,格式固定，切勿修改
		partnerId := v.MchId
		if v.SubMchId != "" {
			partnerId = v.SubMchId
		}
		configMap := map[string]string{
			"prepayid":  prepayId,
			"partnerid": partnerId,
			"package":   "Sign=WXPay",
			"timestamp": timestamp,
			"noncestr":  nonceStr,
			"appid":     appId,
		}
		sign := p.sign(configMap, request.SignType, p.GetWxPayConfig().MchKey)
		configMap["sign"] = sign
		return json.Marshal(configMap)
	case JSAPI:
		st := request.SignType
		if st == "" {
			st = MD5
		}
		appId := v.AppId
		if v.SubAppId != "" {
			appId = v.SubAppId
		}
		configMap := map[string]string{
			"package":   "prepay_id=" + prepayId,
			"timestamp": timestamp,
			"noncestr":  nonceStr,
			"appid":     appId,
			"sign_type": string(st),
		}
		sign := p.sign(configMap, request.SignType, p.GetWxPayConfig().MchKey)
		configMap["sign"] = sign
		return json.Marshal(configMap)
	default:
		return nil, fmt.Errorf("该交易类型暂不支持")
	}
}

func (p *WxPayV2ServiceImpl) UnifyOrder(request WxPayUnifiedOrderRequest) (*WxPayUnifiedOrderResult, error) {
	url := p.GetPayBaseUr() + "/pay/unifiedorder"
	var res WxPayUnifiedOrderResult
	err := p.PostFor(&res, url, "", request)
	return &res, err
}

func (p *WxPayV2ServiceImpl) CloseOrder(string) (*WxPayOrderCloseResult, error) {
	//todo 关闭交易
	return nil, nil
}

func (p *WxPayV2ServiceImpl) QueryOrder(string) (*WxPayOrderQueryResult, error) {
	//todo 查询交易
	return nil, nil
}

func (p *WxPayV2ServiceImpl) SetWxPayConfig(config *WxPayConfig) {
	p.config = config
}

func (p *WxPayV2ServiceImpl) GetWxPayConfig() *WxPayConfig {
	return p.config
}

func (p *WxPayV2ServiceImpl) GetEntPayService() WxEntPayService {
	return p.entPayService
}

// 微信支付签名算法(详见:https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3).
func (p WxPayV2ServiceImpl) sign(params map[string]string, st SignType, sk string, ignoreParams ...string) string {
	signStr := buildSign(params, sk, ignoreParams...)
	var sign string
	switch st {
	case HmacSha256:
		sign = util.HmacSha256(signStr, sk)
	case MD5:
		sign = util.Md5(signStr)
	}
	return sign
}
