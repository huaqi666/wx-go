package pay

type BaseWxPayRequest struct {
	AppId      string `json:"appid"`
	MchId      string `json:"mch_id"`
	SubAppId   string `json:"sub_app_id"`
	SubMchId   string `json:"sub_mch_id"`
	NonceStr   string `json:"nonce_str"`
	Sign       string `json:"sign"`
	SignType   string `json:"sign_type"`
	WorkWxSign string `json:"work_wx_sign"`
}

type WxPayUnifiedOrderRequest struct {
	BaseWxPayRequest
	// 接口版本号.
	// 是否必填：单品优惠必填
	// 类型：String(32) 示例值：1.0
	// 描述：单品优惠新增字段，接口版本号，区分原接口，默认填写1.0。
	// 入参新增version后，则支付通知接口也将返回单品优惠信息字段promotion_detail，请确保支付通知的签名验证能通过。
	// 更多信息，详见文档：https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_102&index=2
	Version        string `json:"version"`
	DeviceInfo     string `json:"device_info"`
	Body           string `json:"body"`
	Detail         string `json:"detail"`
	Attach         string `json:"attach"`
	OutTradeNo     string `json:"out_trade_no"`
	FeeType        string `json:"fee_type"`
	TotalFee       uint64 `json:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	TimeStart      string `json:"time_start"`
	TimeExpire     string `json:"time_expire"`
	GoodsTag       string `json:"goods_tag"`
	NotifyUrl      string `json:"notify_url"`
	TradeType      string `json:"trade_type"`
	ProductId      string `json:"product_id"`
	LimitPay       string `json:"limit_pay"`
	Openid         string `json:"openid"`
	SubOpenid      string `json:"sub_openid"`
	Receipt        string `json:"receipt"`
	SceneInfo      string `json:"scene_info"`
	Fingerprint    string `json:"fingerprint"`
	ProfitSharing  string `json:"profit_sharing"`
}
