package pay

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"strings"
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
	UnifyPay(*WxPayUnifiedOrderRequest) ([]byte, error)
	/** 统一下单
	  (详见https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1)
	  在发起微信支付前，需要调用统一下单接口，获取"预支付交易会话标识"
	  接口地址：https://api.mch.weixin.qq.com/pay/unifiedorder */
	UnifyOrder(*WxPayUnifiedOrderRequest) (*WxPayUnifiedOrderResult, error)

	/* 关闭订单.
	   应用场景
	   以下情况需要调用关单接口：
	   1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
	   2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
	   注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。
	   接口地址：https://api.mch.weixin.qq.com/pay/closeorder
	   是否需要证书：   不需要。 */
	CloseOrderBy(string) (*WxPayOrderCloseResult, error)
	/* 关闭订单.
	   应用场景
	   以下情况需要调用关单接口：
	   1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
	   2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
	   注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。
	   接口地址：https://api.mch.weixin.qq.com/pay/closeorder
	   是否需要证书：   不需要。 */
	CloseOrder(*WxPayOrderCloseRequest) (*WxPayOrderCloseResult, error)
	/* 查询订单（适合于需要自定义子商户号和子商户appid的情形）.
	   详见https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
	   该接口提供所有微信支付订单的查询，商户可以通过查询订单接口主动查询订单状态，完成下一步的业务逻辑。
	   需要调用查询接口的情况：
	   ◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；
	   ◆ 调用支付接口后，返回系统错误或未知交易状态情况；
	   ◆ 调用被扫支付API，返回USERPAYING的状态；
	   ◆ 调用关单或撤销接口API之前，需确认支付状态； */
	QueryOrderBy(outTradeNo, transactionId string) (*WxPayOrderQueryResult, error)
	/* 查询订单（适合于需要自定义子商户号和子商户appid的情形）.
	   详见https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
	   该接口提供所有微信支付订单的查询，商户可以通过查询订单接口主动查询订单状态，完成下一步的业务逻辑。
	   需要调用查询接口的情况：
	   ◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；
	   ◆ 调用支付接口后，返回系统错误或未知交易状态情况；
	   ◆ 调用被扫支付API，返回USERPAYING的状态；
	   ◆ 调用关单或撤销接口API之前，需确认支付状态； */
	QueryOrder(*WxPayOrderQueryRequest) (*WxPayOrderQueryResult, error)

	// 获取配置
	GetWxPayConfig() *WxPayConfig
	// 设置配置
	SetWxPayConfig(*WxPayConfig)

	// 获取提现接口
	GetEntPayService() WxEntPayService

	GetSandboxSignKey(*BaseWxPayRequest) (*WxPaySandboxSignKeyResult, error)
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
	contentType = "application/xml; charset=utf-8"
	return p.http.Post(url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	contentType = "application/xml; charset=utf-8"
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
	if url == "" {
		url = common.PayDefaultPayBaseUrl
	}
	if p.GetWxPayConfig().UseSandboxEnv {
		url += "/sandboxnew"
	}
	return url
}

func (p *WxPayV2ServiceImpl) UnifyPay(request *WxPayUnifiedOrderRequest) ([]byte, error) {
	if request == nil {
		return nil, fmt.Errorf("参数为空")
	}
	v, err := p.UnifyOrder(request)
	if err != nil {
		return nil, err
	}
	prepayId := v.PrepayId
	if prepayId == "" {
		if v.ErrCode != "" {
			return nil, fmt.Errorf("无法获取prepay id，错误代码： '%s'，信息：%s。", v.ErrCode, v.ErrCodeDes)
		} else {
			msg := v.ReturnMsg
			if msg == "" {
				msg = v.RetMsg
			}
			return nil, fmt.Errorf("无法获取prepay id，错误代码： '%s'，信息：%s。", v.ReturnCode, msg)
		}
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
		configMap := map[string]interface{}{
			"prepayid":  prepayId,
			"partnerid": partnerId,
			"package":   "Sign=WXPay",
			"timestamp": timestamp,
			"noncestr":  nonceStr,
			"appid":     appId,
		}
		sign := p.signForMap(configMap, request.SignType, p.GetWxPayConfig().MchKey)
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
		configMap := map[string]interface{}{
			"package":   "prepay_id=" + prepayId,
			"timestamp": timestamp,
			"noncestr":  nonceStr,
			"appid":     appId,
			"sign_type": string(st),
		}
		sign := p.signForMap(configMap, request.SignType, p.GetWxPayConfig().MchKey)
		configMap["sign"] = sign
		return json.Marshal(configMap)
	default:
		return nil, fmt.Errorf("该交易类型暂不支持")
	}
}

func (p *WxPayV2ServiceImpl) UnifyOrder(request *WxPayUnifiedOrderRequest) (*WxPayUnifiedOrderResult, error) {
	if request == nil {
		return nil, fmt.Errorf("参数为空")
	}

	c := p.GetWxPayConfig()
	if request.NotifyUrl == "" && c.NotifyUrl == "" {
		return nil, fmt.Errorf("参数为空")
	}
	p.convert(request)

	request.NonceStr = util.RandSeq(32)
	if c.UseSandboxEnv {
		re, err := p.GetSandboxSignKey(&request.BaseWxPayRequest)
		if err != nil {
			return nil, err
		}
		request.Sign = re.SandboxSignkey
	} else {
		request.Sign = p.signForObj(request, c.SignType, c.MchKey)
	}

	url := p.GetPayBaseUr() + "/pay/unifiedorder"

	var res WxPayUnifiedOrderResult
	err := p.PostFor(&res, url, "", request)

	return &res, err
}

func (p *WxPayV2ServiceImpl) CloseOrderBy(outTradeNo string) (*WxPayOrderCloseResult, error) {
	if outTradeNo == "" {
		return nil, fmt.Errorf("outTradeNo不能为空")
	}
	return p.CloseOrder(&WxPayOrderCloseRequest{
		OutTradeNo: outTradeNo,
	})
}

func (p *WxPayV2ServiceImpl) CloseOrder(request *WxPayOrderCloseRequest) (*WxPayOrderCloseResult, error) {
	if request == nil || request.OutTradeNo == "" {
		return nil, fmt.Errorf("outTradeNo不能为空")
	}

	url := p.GetPayBaseUr() + "/pay/closeorder"

	request.Sign = p.signForObj(request, p.GetWxPayConfig().SignType, p.GetWxPayConfig().MchKey)

	var res WxPayOrderCloseResult
	err := p.PostFor(&res, url, "", request)
	return &res, err
}

func (p *WxPayV2ServiceImpl) QueryOrderBy(outTradeNo, transactionId string) (*WxPayOrderQueryResult, error) {
	if outTradeNo != "" || transactionId != "" {
		return nil, fmt.Errorf("参数为空")
	}
	return p.QueryOrder(&WxPayOrderQueryRequest{
		OutTradeNo:    outTradeNo,
		TransactionId: transactionId,
	})
}

func (p *WxPayV2ServiceImpl) QueryOrder(request *WxPayOrderQueryRequest) (*WxPayOrderQueryResult, error) {

	url := p.GetPayBaseUr() + "/pay/orderquery"

	request.Sign = p.signForObj(request, p.GetWxPayConfig().SignType, p.GetWxPayConfig().MchKey)
	request.Version = "1.0"

	var res WxPayOrderQueryResult
	err := p.PostFor(&res, url, "", request)
	return &res, err
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

func (p *WxPayV2ServiceImpl) GetSandboxSignKey(request *BaseWxPayRequest) (*WxPaySandboxSignKeyResult, error) {
	url := "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

	var res WxPaySandboxSignKeyResult
	//request.Sign = p.signForObj(request, request.SignType, p.GetWxPayConfig().MchKey)
	//err := p.PostFor(&res, url, "", request)
	data := map[string]interface{}{
		"mch_id":    request.MchId,
		"nonce_str": request.NonceStr,
	}
	data["sign"] = p.signForMap(data, request.SignType, p.GetWxPayConfig().MchKey)
	err := p.PostFor(&res, url, "", data)

	return &res, err
}

// 微信支付签名算法(详见:https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3).
func (p *WxPayV2ServiceImpl) signForMap(params map[string]interface{}, st SignType, sk string, ignoreParams ...string) string {
	signStr := buildSign(params, sk, ignoreParams...)
	var sign string
	switch st {
	case HmacSha256:
		sign = util.HmacSha256(signStr, sk)
	case MD5:
		sign = util.Md5(signStr)
	}
	return strings.ToUpper(sign)
}

// 微信支付签名算法(详见:https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3).
func (p *WxPayV2ServiceImpl) signForObj(params interface{}, st SignType, sk string, ignoreParams ...string) string {
	signStr := buildSignFor(params, sk, ignoreParams...)
	var sign string
	switch st {
	case HmacSha256:
		sign = util.HmacSha256(signStr, sk)
	case MD5:
		sign = util.Md5(signStr)
	}
	return strings.ToUpper(sign)
}

func (p *WxPayV2ServiceImpl) convert(request *WxPayUnifiedOrderRequest) {

	c := p.GetWxPayConfig()

	if request.NotifyUrl == "" {
		request.NotifyUrl = p.GetWxPayConfig().NotifyUrl
	}
	if request.AppId == "" {
		request.AppId = c.AppId
	}
	if request.SignType == "" {
		request.SignType = c.SignType
	}
	if request.MchId == "" {
		request.MchId = c.MchId
	}
	if request.SubAppId == "" {
		request.SubAppId = c.SubAppId
	}
	if request.SubMchId == "" {
		request.SubMchId = c.SubMchId
	}
	if request.SpbillCreateIp == "" {
		request.SpbillCreateIp = "127.0.0.1"
	}
	if request.TradeType == "" {
		request.TradeType = JSAPI
	}
}

func NewWxPayService(appId, mchId, mchKey, notifyUrl, keyPath string) WxPayService {
	return newWxPayService(NewBaseV2Config(appId, mchId, mchKey, notifyUrl, keyPath))
}

func NewWxPayServiceFor(config *WxPayConfig) WxPayService {
	return newWxPayService(config)
}
