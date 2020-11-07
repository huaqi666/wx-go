package pay

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"strings"
)

// 默认参数，与基础参数一致
type WxPayDefaultRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest
}

func (r *WxPayDefaultRequest) IsIgnoreAppId() bool {
	return true
}

// 统一下单请求对象
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

func (r *WxPayUnifiedOrderRequest) CheckAndSign(c *WxPayConfig) {

	r.BaseWxPayRequest.CheckAndSign(c)

	if r.NotifyUrl == "" {
		r.NotifyUrl = c.NotifyUrl
	}
	if r.SpbillCreateIp == "" {
		r.SpbillCreateIp = "127.0.0.1"
	}
	if r.TradeType == "" {
		r.TradeType = JSAPI
	}
}

// 统一下单响应对象
type WxPayUnifiedOrderResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
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

// 订单查询请求对象
type WxPayOrderQueryRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest

	Version       string `json:"version" xml:"version"`
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no" xml:"out_trade_no"`
}

// 订单查询响应对象
type WxPayOrderQueryResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	PromotionDetail    string `json:"promotion_detail" xml:"promotion_detail"`
	DeviceInfo         string `json:"device_info" xml:"device_info"`
	Openid             string `json:"openid" xml:"openid"`
	IsSubscribe        string `json:"is_subscribe" xml:"is_subscribe"`
	SubOpenid          string `json:"sub_openid" xml:"sub_openid"`
	IsSubscribeSub     string `json:"is_subscribe_sub" xml:"is_subscribe_sub"`
	TradeType          string `json:"trade_type" xml:"trade_type"`
	TradeState         string `json:"trade_state" xml:"trade_state"`
	BankType           string `json:"bank_type" xml:"bank_type"`
	Detail             string `json:"detail" xml:"detail"`
	TotalFee           uint64 `json:"total_fee" xml:"total_fee"`
	FeeType            string `json:"fee_type" xml:"fee_type"`
	SettlementTotalFee uint64 `json:"settlement_total_fee" xml:"settlement_total_fee"`
	CashFee            uint64 `json:"cash_fee" xml:"cash_fee"`
	CashFeeType        string `json:"cash_fee_type" xml:"cash_fee_type"`
	CouponFee          uint64 `json:"coupon_fee" xml:"coupon_fee"`
	CouponCount        uint64 `json:"coupon_count" xml:"coupon_count"`
	TransactionId      string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo         string `json:"out_trade_no" xml:"out_trade_no"`
	Attach             string `json:"attach" xml:"attach"`
	TimeEnd            string `json:"time_end" xml:"time_end"`
	TradeStateDesc     string `json:"trade_state_desc" xml:"trade_state_desc"`

	Coupons []*WxPayOrderCoupon `json:"coupons" xml:"-"`
}

func (r WxPayOrderQueryResult) Compose() {
	if r.CouponCount > 0 {
		r.Coupons = []*WxPayOrderCoupon{}
		doc := etree.NewDocument()
		err := doc.ReadFromBytes(r.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := uint64(0); i < r.CouponCount; i++ {
			fee, _ := strconv.ParseUint(doc.SelectAttrValue("coupon_fee_"+strconv.FormatUint(i, 10), ""), 10, 64)
			r.Coupons = append(r.Coupons, &WxPayOrderCoupon{
				CouponType: doc.SelectAttrValue("coupon_type_"+strconv.FormatUint(i, 10), ""),
				CouponId:   doc.SelectAttrValue("coupon_id_"+strconv.FormatUint(i, 10), ""),
				CouponFee:  fee,
			})
		}
	}
}

// 订单关闭请求对象
type WxPayOrderCloseRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest

	// 订单支付编号
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
}

// 订单关闭响应对象
type WxPayOrderCloseResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	// 业务结果描述
	ResultMsg string `json:"result_msg" xml:"result_msg"`
}

// 退款请求对象
type WxPayRefundRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest

	// 必填 ...
	// transaction_id 和 out_trade_no 二选一
	TransactionId string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	// transaction_id 和 out_trade_no 二选一
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no" xml:"out_refund_no"`
	// 订单总金额，单位为分，只能为整数
	TotalFee uint64 `json:"total_fee" xml:"total_fee"`
	// 退款金额，最大订单总金额，单位为分，只能为整数
	RefundFee uint64 `json:"refund_fee" xml:"refund_fee"`

	// 选填 ...
	DeviceInfo    string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	RefundFeeType string `json:"refund_fee_type,omitempty" xml:"refund_fee_type,omitempty"`
	// 操作员帐号, 默认为商户号
	OpUserId      string `json:"op_user_id,omitempty" xml:"op_user_id,omitempty"`
	RefundAccount string `json:"refund_account,omitempty" xml:"refund_account,omitempty"`
	RefundDesc    string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
	NotifyUrl     string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	Detail        string `json:"detail,omitempty" xml:"detail,omitempty"`
}

func (r *WxPayRefundRequest) CheckAndSign(c *WxPayConfig) {
	r.BaseWxPayRequest.CheckAndSign(c)

	if r.OpUserId == "" {
		r.OpUserId = c.MchId
	}
}

// 退款响应对象
type WxPayRefundResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	// transaction_id 和 out_trade_no 二选一
	TransactionId string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	// transaction_id 和 out_trade_no 二选一
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo string `json:"out_refund_no" xml:"out_refund_no"`
	// 微信退款单号.
	RefundId string `json:"refund_id" xml:"refund_id"`
	// 退款金额，最大订单总金额，单位为分，只能为整数
	RefundFee uint64 `json:"refund_fee" xml:"refund_fee"`
	// 订单总金额，单位为分，只能为整数
	TotalFee uint64 `json:"total_fee" xml:"total_fee"`

	SettlementRefundFee   uint64 `json:"settlement_refund_fee" xml:"settlement_refund_fee"`
	SettlementTotalFee    uint64 `json:"settlement_total_fee" xml:"settlement_total_fee"`
	FeeType               string `json:"fee_type" xml:"fee_type"`
	CashFee               uint64 `json:"cash_fee" xml:"cash_fee"`
	CashFeeType           string `json:"cash_fee_type" xml:"cash_fee_type"`
	CashRefundFee         uint64 `json:"cash_refund_fee" xml:"cash_refund_fee"`
	CouponRefundCount     uint64 `json:"coupon_refund_count" xml:"coupon_refund_count"`
	CouponRefundFee       uint64 `json:"coupon_refund_fee" xml:"coupon_refund_fee"`
	PromotionDetailString string `json:"promotion_detail_string" xml:"promotion_detail_string"`

	RefundCoupons    []*WxPayRefundCoupon          `json:"refund_coupons" xml:"-"`
	PromotionDetails []*WxPayRefundPromotionDetail `json:"promotion_details" xml:"-"`
}

func (r *WxPayRefundResult) Compose() {
	if r.CouponRefundCount > 0 {
		r.RefundCoupons = []*WxPayRefundCoupon{}
		doc := etree.NewDocument()
		err := doc.ReadFromBytes(r.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := uint64(0); i < r.CouponRefundCount; i++ {
			istr := strconv.FormatUint(i, 10)
			fee, _ := strconv.ParseUint(doc.SelectAttrValue("coupon_refund_fee_"+istr, ""), 10, 64)
			r.RefundCoupons = append(r.RefundCoupons, &WxPayRefundCoupon{
				CouponType:      doc.SelectAttrValue("coupon_type_"+istr, ""),
				CouponRefundId:  doc.SelectAttrValue("coupon_refund_id_"+istr, ""),
				CouponRefundFee: fee,
			})
		}
	}

	if r.PromotionDetailString != "" {
		_ = json.Unmarshal([]byte(r.PromotionDetailString), &r.PromotionDetails)
	}
}

type WxPayRefundQueryRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest

	// 必填 ...
	// 以下四选一
	// 微信订单号
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
	// 微信订单号
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
	// 商户退款单号
	OutRefundNo string `json:"out_refund_no" xml:"out_refund_no"`
	// 商户退款单号
	RefundId string `json:"refund_id" xml:"refund_id"`
	// 订单总金额，单位为分，只能为整数
	TotalFee uint64 `json:"total_fee" xml:"total_fee"`

	// 选填 ...
	// 设备号
	DeviceInfo string `json:"device_info" xml:"device_info"`
}

type WxPayRefundQueryResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	// 微信订单号
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
	// 微信订单号
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
	// 订单总金额，单位为分，只能为整数
	TotalFee uint64 `json:"total_fee" xml:"total_fee"`

	// 设备号
	DeviceInfo            string `json:"device_info" xml:"device_info"`
	SettlementTotalFee    uint64 `json:"settlement_total_fee" xml:"settlement_total_fee"`
	FeeType               string `json:"fee_type" xml:"fee_type"`
	CashFee               uint64 `json:"cash_fee" xml:"cash_fee"`
	RefundCount           uint64 `json:"refund_count" xml:"refund_count"`
	PromotionDetailString string `json:"promotion_detail_string" xml:"promotion_detail_string"`

	RefundRecords    []*WxPayRefundRecord          `json:"refund_records" xml:"-"`
	PromotionDetails []*WxPayRefundPromotionDetail `json:"promotion_details" xml:"-"`
}

func (r *WxPayRefundQueryResult) Compose() {
	if r.RefundCount > 0 {
		r.RefundRecords = []*WxPayRefundRecord{}
		doc := etree.NewDocument()
		err := doc.ReadFromBytes(r.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := uint64(0); i < r.RefundCount; i++ {
			istr := strconv.FormatUint(i, 10)

			var rr WxPayRefundRecord
			rr.OutRefundNo = doc.SelectAttrValue("out_refund_no_"+istr, "")
			rr.RefundId = doc.SelectAttrValue("refund_id_"+istr, "")
			rr.RefundChannel = doc.SelectAttrValue("refund_channel_"+istr, "")
			var ii uint64
			ii, _ = strconv.ParseUint(doc.SelectAttrValue("refund_fee_"+istr, ""), 10, 64)
			rr.RefundFee = ii
			ii, _ = strconv.ParseUint(doc.SelectAttrValue("settlement_refund_fee_"+istr, ""), 10, 64)
			rr.SettlementRefundFee = ii
			ii, _ = strconv.ParseUint(doc.SelectAttrValue("coupon_refund_fee_"+istr, ""), 10, 64)
			rr.CouponRefundFee = ii
			ii, _ = strconv.ParseUint(doc.SelectAttrValue("coupon_refund_count_"+istr, ""), 10, 64)
			rr.CouponRefundCount = ii
			rr.RefundStatus = doc.SelectAttrValue("refund_status_"+istr, "")
			rr.RefundRecvAccount = doc.SelectAttrValue("refund_recv_accout_"+istr, "")
			rr.RefundSuccessTime = doc.SelectAttrValue("refund_success_time_"+istr, "")

			var rc []*WxPayRefundCoupon
			for j := uint64(0); j < rr.CouponRefundCount; j++ {
				jstr := strconv.FormatUint(j, 10)
				fee, _ := strconv.ParseUint(doc.SelectAttrValue("coupon_refund_fee_"+istr+"_"+jstr, ""), 10, 64)
				rc = append(rc, &WxPayRefundCoupon{
					CouponType:      doc.SelectAttrValue("coupon_type_"+istr+"_"+jstr, ""),
					CouponRefundId:  doc.SelectAttrValue("coupon_refund_id_"+istr+"_"+jstr, ""),
					CouponRefundFee: fee,
				})
			}

			rr.RefundCoupons = rc
			r.RefundRecords = append(r.RefundRecords, &rr)
		}
	}

	if r.PromotionDetailString != "" {
		_ = json.Unmarshal([]byte(r.PromotionDetailString), &r.PromotionDetails)
	}
}

type WxPayOrderNotifyResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	PromotionDetail    string   `json:"promotion_detail" xml:"promotion_detail"`
	DeviceInfo         string   `json:"device_info" xml:"device_info"`
	Openid             string   `json:"openid" xml:"openid"`
	IsSubscribe        string   `json:"is_subscribe" xml:"is_subscribe"`
	SubOpenid          string   `json:"sub_openid" xml:"sub_openid"`
	SubIsSubscribe     string   `json:"sub_is_subscribe" xml:"sub_is_subscribe"`
	TradeType          string   `json:"trade_type" xml:"trade_type"`
	BankType           string   `json:"bank_type" xml:"bank_type"`
	TotalFee           string   `json:"total_fee" xml:"total_fee"`
	SettlementTotalFee string   `json:"settlement_total_fee" xml:"settlement_total_fee"`
	FeeType            string   `json:"fee_type" xml:"fee_type"`
	CashFee            string   `json:"cash_fee" xml:"cash_fee"`
	CashFeeType        string   `json:"cash_fee_type" xml:"cash_fee_type"`
	CouponFee          string   `json:"coupon_fee" xml:"coupon_fee"`
	CouponCount        uint64   `json:"coupon_count" xml:"coupon_count"`
	TransactionId      string   `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo         string   `json:"out_trade_no" xml:"out_trade_no"`
	Attach             string   `json:"attach" xml:"attach"`
	TimeEnd            string   `json:"time_end" xml:"time_end"`
	Version            string   `json:"version" xml:"version"`
	RateValue          string   `json:"rate_value" xml:"rate_value"`
	SignType           SignType `json:"sign_type" xml:"sign_type"`

	Coupons []*WxPayOrderCoupon `json:"coupons" xml:"-"`
}

func (r WxPayOrderNotifyResult) CheckResult(service WxPayService, signType SignType, checkSuccess bool) error {
	if r.ReturnCode == common.Success && r.Sign == "" {
		return common.ErrorOf("伪造的通知!")
	}
	return r.BaseWxPayResult.CheckResult(service, signType, checkSuccess)
}

func (r *WxPayOrderNotifyResult) Compose() {
	if r.CouponCount > 0 {
		r.Coupons = []*WxPayOrderCoupon{}
		doc := etree.NewDocument()
		err := doc.ReadFromBytes(r.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := uint64(0); i < r.CouponCount; i++ {
			fee, _ := strconv.ParseUint(doc.SelectAttrValue("coupon_fee_"+strconv.FormatUint(i, 10), ""), 10, 64)
			r.Coupons = append(r.Coupons, &WxPayOrderCoupon{
				CouponType: doc.SelectAttrValue("coupon_type_"+strconv.FormatUint(i, 10), ""),
				CouponId:   doc.SelectAttrValue("coupon_id_"+strconv.FormatUint(i, 10), ""),
				CouponFee:  fee,
			})
		}
	}
}

type WxPayRefundNotifyResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	ReqInfoString    string  `json:"-" xml:"req_info"`
	ReqInfo          ReqInfo `json:"req_info" xml:"-"`
	DecryptedReqInfo string  `json:"-" xml:"-"`

	MchKey string `json:"-" xml:"-"`
}

func (r *WxPayRefundNotifyResult) Compose() {
	if r.ReturnCode == common.Fail {
		return
	}
	// 1.对加密串A做base64解码，得到加密串B
	data := base64.StdEncoding.EncodeToString([]byte(r.ReqInfoString))
	// 2.对商户key做md5，得到32位小写key*
	key := strings.ToLower(util.Md5(r.MchKey))
	// 3.用key*对加密串B做AES-256-ECB解密
	b := util.AesDecryptECB([]byte(data), []byte(key))
	r.DecryptedReqInfo = string(b)
	_ = xml.Unmarshal(b, &r.ReqInfo)
}

type WxScanPayNotifyResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult

	Openid      string `json:"openid" xml:"openid"`
	IsSubscribe string `json:"is_subscribe" xml:"is_subscribe"`
	ProductId   string `json:"product_id" xml:"product_id"`
}

// 沙河请求签名结果
type WxPaySandboxSignKeyResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayResult
	SandboxSignkey string `json:"sandbox_signkey" xml:"sandbox_signkey"`
}
