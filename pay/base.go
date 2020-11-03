package pay

import (
	"encoding/xml"
	"github.com/cliod/wx-go/common"
)

type (
	TradeType string
	SignType  string
)

const (
	JSAPI  TradeType = "JSAPI"
	NATIVE TradeType = "NATIVE"
	APP    TradeType = "APP"
	H5     TradeType = "MWEB"
	MICRO  TradeType = "MICROPAY"
)

const (
	HmacSha256 SignType = "HMAC_SHA256"
	MD5        SignType = "MD5"
)

type BaseWxPayResult struct {
	common.Err
	RetMsg string `json:"retmsg" xml:"retmsg"`

	ReturnCode string `json:"return_code" xml:"return_code"`
	ReturnMsg  string `json:"return_msg" xml:"return_msg"`

	ResultCode string `json:"result_code" xml:"result_code"`
	ErrCode    string `json:"err_code" xml:"err_code"`
	ErrCodeDes string `json:"err_code_des" xml:"err_code_des"`
	AppId      string `json:"appid" xml:"appid"`
	MchId      string `json:"mch_id" xml:"mch_id"`
	SubAppId   string `json:"sub_app_id" xml:"sub_app_id"`
	SubMchId   string `json:"sub_mch_id" xml:"sub_mch_id"`
	NonceStr   string `json:"nonce_str" xml:"nonce_str"`
	Sign       string `json:"sign" xml:"sign"`
}

type BaseWxPayRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// 公众号或者小程序appId
	AppId string `json:"appid" xml:"appid"`
	// 商户号.
	MchId string `json:"mch_id" xml:"mch_id"`
	// 服务商模式下的子商户公众账号ID.
	SubAppId string `json:"sub_app_id,omitempty" xml:"sub_app_id,omitempty"`
	// 服务商模式下的子商户号.
	SubMchId string `json:"sub_mch_id,omitempty" xml:"sub_mch_id,omitempty"`
	// 随机字符串.不长于32位。推荐随机数生成算法
	NonceStr string `json:"nonce_str" xml:"nonce_str"`
	// 签名.
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 签名类型.
	SignType SignType `json:"sign_type" xml:"sign_type,omitempty"`
	// 企业微信签名
	WorkWxSign string `json:"work_wx_sign,omitempty" xml:"work_wx_sign,omitempty"`
}

type WxPayUnifiedOrderRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest
	// 必填 ...

	/* 总金额.
	   订单总金额，单位为分，详见支付金额 */
	TotalFee uint64 `json:"total_fee" xml:"total_fee"`
	/* 通知地址.
	   接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。 */
	NotifyUrl string `json:"notify_url" xml:"notify_url"`
	/* 用户标识.
	   trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。
	   openid如何获取，可参考【获取openid】。
	   企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换 */
	Openid string `json:"openid" xml:"openid"`
	// 商品描述.
	Body string `json:"body" xml:"body"`
	/* 商户订单号.
	   商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号 */
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`

	// 选填 ...

	/* 附加数据.
	   附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据 */
	Attach string `json:"attach,omitempty" xml:"attach,omitempty"`
	/* 终端IP.
	   APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。 */
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" xml:"spbill_create_ip,omitempty"`
	/* 接口版本号.
	   是否必填：单品优惠必填
	   类型：String(32) 示例值：1.0
	   描述：单品优惠新增字段，接口版本号，区分原接口，默认填写1.0。
	   入参新增version后，则支付通知接口也将返回单品优惠信息字段promotion_detail，请确保支付通知的签名验证能通过。
	   更多信息，详见文档：https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_102&index=2 */
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	/* 设备号.
	   类型：String(32)
	   终端设备号(门店号或收银设备Id)，注意：PC网页或公众号内支付请传"WEB" */
	DeviceInfo string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 商品详情.
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	/* 货币类型.
	   符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型 */
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	/* 交易起始时间.
	   订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则 */
	TimeStart string `json:"time_start,omitempty" xml:"time_start,omitempty"`
	/* 交易结束时间.
	   订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
	   注意：最短失效时间间隔必须大于5分钟 */
	TimeExpire string `json:"time_expire,omitempty" xml:"time_expire,omitempty"`
	GoodsTag   string `json:"goods_tag,omitempty" xml:"goods_tag,omitempty"`
	/* 交易类型.
	   取值如下：
	   JSAPI--公众号支付、
	   NATIVE --原生扫码支付、
	   APP--app支付，统一下单接口trade_type的传参可参考这里 */
	TradeType     TradeType `json:"trade_type" xml:"trade_type"`
	ProductId     string    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	LimitPay      string    `json:"limit_pay,omitempty" xml:"limit_pay,omitempty"`
	SubOpenid     string    `json:"sub_openid,omitempty" xml:"sub_openid,omitempty"`
	Receipt       string    `json:"receipt,omitempty" xml:"receipt,omitempty"`
	SceneInfo     string    `json:"scene_info,omitempty" xml:"scene_info,omitempty"`
	Fingerprint   string    `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	ProfitSharing string    `json:"profit_sharing,omitempty" xml:"profit_sharing,omitempty"`
}

type WxPayUnifiedOrderResult struct {
	BaseWxPayResult
	// 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	PrepayId string `json:"prepay_id" xml:"prepay_id"`
	// 交易类型，取值为：JSAPI，NATIVE，APP等
	TradeType TradeType `json:"trade_type" xml:"trade_type"`
	// mweb_url 支付跳转链接
	MwebUrl string `json:"mweb_url" xml:"mweb_url"`
	// trade_type为NATIVE时有返回，用于生成二维码，展示给用户进行扫码支付
	CodeURL string `json:"code_url" xml:"code_url"`
}

type WxPayOrderCloseResult struct {
	BaseWxPayResult
	// 业务结果描述
	ResultMsg string `json:"result_msg" xml:"result_msg"`
}

type WxPayOrderQueryResult struct {
	BaseWxPayResult

	PromotionDetail    string    `json:"promotion_detail" xml:"promotion_detail"`
	DeviceInfo         string    `json:"device_info" xml:"device_info"`
	Openid             string    `json:"openid" xml:"openid"`
	IsSubscribe        string    `json:"is_subscribe" xml:"is_subscribe"`
	SubOpenid          string    `json:"sub_openid" xml:"sub_openid"`
	IsSubscribeSub     string    `json:"is_subscribe_sub" xml:"is_subscribe_sub"`
	TradeType          string    `json:"trade_type" xml:"trade_type"`
	TradeState         string    `json:"trade_state" xml:"trade_state"`
	BankType           string    `json:"bank_type" xml:"bank_type"`
	Detail             string    `json:"detail" xml:"detail"`
	TotalFee           int       `json:"total_fee" xml:"total_fee"`
	FeeType            string    `json:"fee_type" xml:"fee_type"`
	SettlementTotalFee int       `json:"settlement_total_fee" xml:"settlement_total_fee"`
	CashFee            int       `json:"cash_fee" xml:"cash_fee"`
	CashFeeType        string    `json:"cash_fee_type" xml:"cash_fee_type"`
	CouponFee          int       `json:"coupon_fee" xml:"coupon_fee"`
	CouponCount        int       `json:"coupon_count" xml:"coupon_count"`
	Coupons            []*Coupon `json:"coupons" xml:"coupons"`
	TransactionId      string    `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo         string    `json:"out_trade_no" xml:"out_trade_no"`
	Attach             string    `json:"attach" xml:"attach"`
	TimeEnd            string    `json:"time_end" xml:"time_end"`
	TradeStateDesc     string    `json:"trade_state_desc" xml:"trade_state_desc"`
}

type Coupon struct {
	CouponType string `json:"coupon_type" xml:"coupon_type"`
	CouponId   string `json:"coupon_id" xml:"coupon_id"`
	CouponFee  int    `json:"coupon_fee" xml:"coupon_fee"`
}

func ComposeCoupons(value string) []*Coupon {
	//todo
	return nil
}

type WxPayOrderCloseRequest struct {
	BaseWxPayRequest

	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
}

type WxPayOrderQueryRequest struct {
	BaseWxPayRequest

	Version       string `json:"version" xml:"version"`
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no" xml:"out_trade_no"`
}

type EntPayRequest struct {
}

type EntPayResult struct {
}

type WxPaySandboxSignKeyResult struct {
	BaseWxPayResult
	SandboxSignkey string `json:"sandbox_signkey" xml:"sandbox_signkey"`
}
