package pay

import (
	"encoding/json"
	"encoding/xml"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WxPayService interface {
	// 执行操作
	Do(url string, request WxPayRequest, res WxPayResult, useKey bool) error
	// 执行Post请求
	Post(url, contentType string, data interface{}, args ...interface{}) ([]byte, error)
	// Post 执行Post请求并将结果转成对象
	PostFor(v interface{}, url, contentType string, data interface{}, args ...interface{}) error
	// Post 携带证书执行Post请求
	PostKey(url string, contentType string, data interface{}, args ...interface{}) ([]byte, error)
	// Post 携带证书执行Post请求并将结果转成对象
	PostKeyFor(v interface{}, url, contentType string, data interface{}, args ...interface{}) error

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

	/* 微信支付-申请退款.
	   详见 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
	   1、交易时间超过一年的订单无法提交退款
	   2、微信支付退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。申请退款总金额不能超过订单金额。 一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号
	   3、请求频率限制：150qps，即每秒钟正常的申请退款请求次数不超过150次
	       错误或无效请求频率限制：6qps，即每秒钟异常或错误的退款申请请求不超过6次
	   4、每个支付订单的部分退款次数不能超过50次
	   5、如果同一个用户有多笔退款，建议分不同批次进行退款，避免并发退款导致退款失败 */
	Refund(*WxPayRefundRequest) (*WxPayRefundResult, error)
	RefundV2(*WxPayRefundRequest) (*WxPayRefundResult, error)
	/* 微信支付-查询退款（适合于需要自定义子商户号和子商户appid的情形）.
	   应用场景：
	     提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，
	     银行卡支付的退款3个工作日后重新查询退款状态。
	   详见 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5 */
	RefundQuery(*WxPayRefundQueryRequest) (*WxPayRefundQueryResult, error)
	/* 微信支付-查询退款API（支持单品）.
	   应用场景：提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态。
	   注意：
	   1、本接口支持查询单品优惠相关退款信息，且仅支持按微信退款单号或商户退款单号查询，若继续调用老查询退款接口，
	      请见https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5
	   2、请求频率限制：300qps，即每秒钟正常的退款查询请求次数不超过300次
	   3、错误或无效请求频率限制：6qps，即每秒钟异常或错误的退款查询请求不超过6次 */
	RefundQueryV2(*WxPayRefundQueryRequest) (*WxPayRefundQueryResult, error)
	/* 解析支付结果通知.
	   详见https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7 */
	ParseOrderNotifyResult(xmlData string, signType SignType) (*WxPayOrderNotifyResult, error)
	/* 解析退款结果通知
	   详见https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=9 */
	ParseRefundNotifyResult(xmlData string) (*WxPayRefundNotifyResult, error)
	/* 解析扫码支付回调通知
	   详见https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=6_4 */
	ParseScanPayNotifyResult(xmlData string) (*WxScanPayNotifyResult, error)

	// 获取配置
	GetWxPayConfig() *WxPayConfig
	// 设置配置
	SetWxPayConfig(*WxPayConfig)

	// 获取提现接口
	GetEntPayService() WxEntPayService

	// 获取沙河环境
	GetSandboxSignKey(*WxPayDefaultRequest) (*WxPaySandboxSignKeyResult, error)

	// 签名
	Sign(params interface{}, st SignType, ignoreParams ...string) string
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
	if contentType == "" {
		contentType = common.PostXmlContentType
	}
	return p.http.Post(url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostFor(v interface{}, url string, contentType string, data interface{}, args ...interface{}) error {
	return p.http.PostFor(v, url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostKey(url, contentType string, data interface{}, args ...interface{}) ([]byte, error) {
	var (
		cli *http.Client
		err error
	)
	c := p.GetWxPayConfig()
	if c.KeyPath != "" {
		certData, err := ioutil.ReadFile(c.KeyPath)
		if err != nil {
			return nil, err
		}
		cli, err = util.NewTLSClientByPkc12(certData, c.MchId)
	} else {
		cli, err = util.NewTLSClient(c.PrivateCertPath, c.PrivateKeyPath)
	}
	if err != nil {
		return nil, err
	}
	ser := common.NewXmlServiceFor(cli)
	if contentType == "" {
		contentType = common.PostXmlContentType
	}
	return ser.Post(url, contentType, data, args...)
}

func (p *WxPayV2ServiceImpl) PostKeyFor(v interface{}, url, contentType string, data interface{}, args ...interface{}) error {
	b, err := p.PostKey(url, contentType, data, args...)
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, v)
}

func (p *WxPayV2ServiceImpl) GetPayBaseUr() string {
	url := p.GetWxPayConfig().GetPayBaseUrl()
	if p.GetWxPayConfig().UseSandboxEnv {
		url += "/sandboxnew"
	}
	return url
}

func (p *WxPayV2ServiceImpl) UnifyPay(request *WxPayUnifiedOrderRequest) ([]byte, error) {
	if request == nil {
		return nil, common.ErrorOf("参数为空")
	}
	v, err := p.UnifyOrder(request)
	if err != nil {
		return nil, err
	}
	prepayId := v.PrepayId
	if prepayId == "" {
		return nil, v
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
		sign := p.Sign(configMap, request.SignType)
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
		sign := p.Sign(configMap, request.SignType)
		configMap["sign"] = sign
		return json.Marshal(configMap)
	default:
		return nil, common.ErrorOf("该交易类型暂不支持")
	}
}

func (p *WxPayV2ServiceImpl) UnifyOrder(request *WxPayUnifiedOrderRequest) (*WxPayUnifiedOrderResult, error) {
	if request == nil {
		return nil, common.ErrorOf("参数为空")
	}

	c := p.GetWxPayConfig()
	if request.NotifyUrl == "" && c.NotifyUrl == "" {
		return nil, common.ErrorOf("参数为空")
	}

	if c.UseSandboxEnv {
		re, err := p.GetSandboxSignKey(&WxPayDefaultRequest{
			BaseWxPayRequest: request.BaseWxPayRequest,
		})
		if err != nil {
			return nil, err
		}
		request.Sign = re.SandboxSignkey
	} else {
		request.Sign = p.sign(request)
	}

	url := p.GetPayBaseUr() + common.PayUnifiedOrder

	var res WxPayUnifiedOrderResult
	err := p.Do(url, request, &res, false)

	return &res, err
}

func (p *WxPayV2ServiceImpl) CloseOrderBy(outTradeNo string) (*WxPayOrderCloseResult, error) {
	if outTradeNo == "" {
		return nil, common.ErrorOf("outTradeNo不能为空")
	}
	return p.CloseOrder(&WxPayOrderCloseRequest{
		OutTradeNo: outTradeNo,
	})
}

func (p *WxPayV2ServiceImpl) CloseOrder(request *WxPayOrderCloseRequest) (*WxPayOrderCloseResult, error) {
	if request == nil || request.OutTradeNo == "" {
		return nil, common.ErrorOf("outTradeNo不能为空")
	}

	url := p.GetPayBaseUr() + common.PayCloseOrder
	var res WxPayOrderCloseResult

	err := p.Do(url, request, &res, false)
	return &res, err
}

func (p *WxPayV2ServiceImpl) QueryOrderBy(outTradeNo, transactionId string) (*WxPayOrderQueryResult, error) {
	if outTradeNo != "" || transactionId != "" {
		return nil, common.ErrorOf("参数为空")
	}
	return p.QueryOrder(&WxPayOrderQueryRequest{
		OutTradeNo:    outTradeNo,
		TransactionId: transactionId,
	})
}

func (p *WxPayV2ServiceImpl) QueryOrder(request *WxPayOrderQueryRequest) (*WxPayOrderQueryResult, error) {
	url := p.GetPayBaseUr() + common.PayQueryOrder

	var res WxPayOrderQueryResult
	err := p.Do(url, request, &res, false)
	return &res, err
}

func (p *WxPayV2ServiceImpl) Refund(request *WxPayRefundRequest) (*WxPayRefundResult, error) {
	url := p.GetPayBaseUr() + common.PayRefundUrl
	if p.GetWxPayConfig().UseSandboxEnv {
		url = p.GetWxPayConfig().GetPayBaseUrl() + common.PayRefundSandboxUrl
	}

	var res WxPayRefundResult
	err := p.Do(url, request, &res, true)
	return &res, err
}

func (p *WxPayV2ServiceImpl) RefundV2(request *WxPayRefundRequest) (*WxPayRefundResult, error) {
	url := p.GetPayBaseUr() + common.PayRefundUrlV2
	if p.GetWxPayConfig().UseSandboxEnv {
		url = p.GetWxPayConfig().GetPayBaseUrl() + common.PayRefundSandboxUrlV2
	}

	var res WxPayRefundResult
	err := p.Do(url, request, &res, true)
	return &res, err
}

func (p *WxPayV2ServiceImpl) RefundQuery(request *WxPayRefundQueryRequest) (*WxPayRefundQueryResult, error) {
	url := p.GetPayBaseUr() + common.PayQueryRefundUrl

	var res WxPayRefundQueryResult
	err := p.Do(url, request, &res, false)
	return &res, err
}

func (p *WxPayV2ServiceImpl) RefundQueryV2(request *WxPayRefundQueryRequest) (*WxPayRefundQueryResult, error) {
	url := p.GetPayBaseUr() + common.PayQueryRefundUrlV2

	var res WxPayRefundQueryResult
	err := p.Do(url, request, &res, false)
	return &res, err
}

func (p *WxPayV2ServiceImpl) ParseOrderNotifyResult(xmlData string, signType SignType) (*WxPayOrderNotifyResult, error) {

	var res WxPayOrderNotifyResult
	err := xml.Unmarshal([]byte(xmlData), &res)
	if err != nil {
		return nil, err
	}
	if signType == "" {
		if res.SignType != "" {
			signType = res.SignType
		} else if p.GetWxPayConfig().SignType != "" {
			signType = p.GetWxPayConfig().SignType
		}
	}
	err = res.CheckResult(p, signType, false)
	return &res, err
}

func (p *WxPayV2ServiceImpl) ParseRefundNotifyResult(xmlData string) (*WxPayRefundNotifyResult, error) {

	var res WxPayRefundNotifyResult
	err := xml.Unmarshal([]byte(xmlData), &res)
	if err != nil {
		return nil, err
	}

	res.MchKey = p.GetWxPayConfig().MchKey
	res.Compose()

	return &res, err
}

func (p *WxPayV2ServiceImpl) ParseScanPayNotifyResult(xmlData string) (*WxScanPayNotifyResult, error) {

	var res WxScanPayNotifyResult
	err := xml.Unmarshal([]byte(xmlData), &res)
	if err != nil {
		return nil, err
	}

	err = res.CheckResult(p, p.GetWxPayConfig().SignType, false)

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

func (p *WxPayV2ServiceImpl) GetSandboxSignKey(request *WxPayDefaultRequest) (*WxPaySandboxSignKeyResult, error) {
	url := common.PayGetSandboxSignKey
	var res WxPaySandboxSignKeyResult
	err := p.Do(url, request, &res, false)
	return &res, err
}

// 微信支付签名算法(详见:https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3).
func (p *WxPayV2ServiceImpl) Sign(params interface{}, st SignType, ignoreParams ...string) string {
	sk := p.GetWxPayConfig().MchKey
	if st == "" {
		st = MD5
	}
	sign := Sign(params, st, sk, ignoreParams...)
	return strings.ToUpper(sign)
}

// 操作
func (p *WxPayV2ServiceImpl) Do(url string, request WxPayRequest, res WxPayResult, useKey bool) error {
	request.CheckAndSign(p.GetWxPayConfig())

	var err error
	if useKey {
		err = p.PostKeyFor(&res, url, common.PostXmlContentType, request)
	} else {
		err = p.PostFor(&res, url, common.PostXmlContentType, request)
	}

	if err == nil {
		res.Compose()
		err = res.CheckResult(p, request.GetSignType(), true)
	}
	return err
}

// 微信支付签名算法(详见:https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3).
func (p *WxPayV2ServiceImpl) sign(params interface{}, ignoreParams ...string) string {
	st := p.GetWxPayConfig().SignType
	return p.Sign(params, st, ignoreParams...)
}

func NewWxPayServiceBy(appId, mchId, mchKey, notifyUrl, keyPath string) WxPayService {
	return newWxPayService(newWxPayV2Config(appId, mchId, mchKey, notifyUrl, keyPath))
}

func NewWxPayService(config *WxPayConfig) WxPayService {
	if config == nil {
		config = new(WxPayConfig)
	}
	return newWxPayService(config)
}
