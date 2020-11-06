package pay

import "encoding/xml"

type EntPayQueryRequest struct {
	BaseWxPayRequest

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
}

type EntPayQueryResult struct {
	BaseWxPayResult

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	DetailId       string `json:"detail_id" xml:"detail_id"`
	Status         string `json:"status" xml:"status"`
	Reason         string `json:"reason" xml:"reason"`
	Openid         string `json:"openid" xml:"openid"`
	TransferName   string `json:"transfer_name" xml:"transfer_name"`
	PaymentAmount  uint64 `json:"payment_amount" xml:"payment_amount"`
	TransferTime   string `json:"transfer_time" xml:"transfer_time"`
	PaymentTime    string `json:"payment_time" xml:"payment_time"`
	Desc           string `json:"desc" xml:"desc"`
}

type EntPayBankRequest struct {
	BaseWxPayRequest

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	EncBankNo      string `json:"enc_bank_no" xml:"enc_bank_no"`
	EncTrueName    string `json:"enc_true_name" xml:"enc_true_name"`
	BankCode       string `json:"bank_code" xml:"bank_code"`
	Amount         uint64 `json:"amount" xml:"amount"`
	Description    string `json:"description" xml:"description"`
}

type EntPayBankResult struct {
	BaseWxPayResult

	Amount         uint64 `json:"amount" xml:"amount"`
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	PaymentNo      string `json:"payment_no" xml:"payment_no"`
	CmmsAmount     uint64 `json:"cmms_amount" xml:"cmms_amount"`
}

type EntPayBankQueryRequest struct {
	EntPayQueryRequest

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
}

func (r EntPayQueryRequest) IsIgnoreAppId() bool {
	return true
}

type EntPayBankQueryResult struct {
	BaseWxPayResult

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	PaymentNo      string `json:"payment_no" xml:"payment_no"`
	BankNoMd5      string `json:"bank_no_md_5" xml:"bank_no_md_5"`
	TrueNameMd5    string `json:"true_name_md_5" xml:"true_name_md_5"`
	Amount         uint64 `json:"amount" xml:"amount"`
	Status         string `json:"status" xml:"status"`
	CmmsAmount     uint64 `json:"cmms_amount" xml:"cmms_amount"`
	CreateTime     string `json:"create_time" xml:"create_time"`
	PaySuccessTime string `json:"pay_success_time" xml:"pay_success_time"`
	FailReason     string `json:"fail_reason" xml:"fail_reason"`
}

type EntPayRedPackRequest struct {
	BaseWxPayRequest

	MchBillNo           string `json:"mch_billno" xml:"mch_billno"`
	WxAppId             string `json:"wx_appid" xml:"wx_appid"`
	SenderName          string `json:"sender_name" xml:"sender_name"`
	AgentId             string `json:"agentid" xml:"agentid"`
	SenderHeaderMediaId string `json:"sender_header_media_id" xml:"sender_header_media_id"`
	ReOpenid            string `json:"re_openid" xml:"re_openid"`
	TotalAmount         uint64 `json:"total_amount" xml:"total_amount"`
	Wishing             string `json:"wishing" xml:"wishing"`
	ActName             string `json:"act_name" xml:"act_name"`
	Remark              string `json:"remark" xml:"remark"`
	SceneId             string `json:"scene_id" xml:"scene_id"`
}

func (r EntPayRedPackRequest) IsIgnoreAppId() bool {
	return true
}

func (r EntPayRedPackRequest) IsIgnoreSubAppId() bool {
	return true
}

func (r EntPayRedPackRequest) IsIgnoreSubMchId() bool {
	return true
}

func (r EntPayRedPackRequest) IsWxWorkSign() bool {
	return true
}

type EntPayRedPackResult struct {
	BaseWxPayResult

	MchBillNo           string `json:"mch_billno" xml:"mch_billno"`
	MchId               string `json:"mch_id" xml:"mch_id"`
	WxAppId             string `json:"wx_appid" xml:"wx_appid"`
	ReOpenid            string `json:"re_openid" xml:"re_openid"`
	TotalAmount         uint64 `json:"total_amount" xml:"total_amount"`
	SendListId          string `json:"sendListid" xml:"sendListid"`
	SenderName          string `json:"sender_name" xml:"sender_name"`
	SenderHeaderMediaId string `json:"sender_header_media_id" xml:"sender_header_media_id"`
}

type EntPayRedPackQueryRequest struct {
	BaseWxPayRequest

	MchBillNo string `json:"mch_billno" xml:"mch_billno"`
}

type EntPayRedPackQueryResult struct {
	BaseWxPayResult

	MchBillNo           string `json:"mch_billno" xml:"mch_billno"`
	DetailId            string `json:"detail_id" xml:"detail_id"`
	Status              string `json:"status" xml:"status"`
	TotalAmount         uint64 `json:"total_amount" xml:"total_amount"`
	Reason              string `json:"reason" xml:"reason"`
	SendTime            string `json:"send_time" xml:"send_time"`
	RefundTime          string `json:"refund_time" xml:"refund_time"`
	RefundAmount        uint64 `json:"refund_amount" xml:"refund_amount"`
	Wishing             string `json:"wishing" xml:"wishing"`
	Remark              string `json:"remark" xml:"remark"`
	ActName             string `json:"act_name" xml:"act_name"`
	Openid              string `json:"openid" xml:"openid"`
	Amount              uint64 `json:"amount" xml:"amount"`
	RcvTime             uint64 `json:"rcv_time" xml:"rcv_time"`
	SenderName          uint64 `json:"sender_name" xml:"sender_name"`
	SenderHeaderMediaId uint64 `json:"sender_header_media_id" xml:"sender_header_media_id"`
}

type GetPublicKeyResult struct {
	BaseWxPayResult

	PubKey string `json:"pub_key" xml:"pub_key"`
}

// 提现请求对象
type WxEntPayRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	BaseWxPayRequest

	AppId          string `json:"mch_appid" xml:"mch_appid"`
	MchId          string `json:"mchid" xml:"mchid"`
	DeviceInfo     string `json:"device_info" xml:"device_info"`
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	Openid         string `json:"openid" xml:"openid"`
	CheckName      string `json:"check_name" xml:"check_name"`
	ReUserName     string `json:"re_user_name" xml:"re_user_name"`
	Amount         uint64 `json:"amount" xml:"amount"`
	Description    string `json:"description" xml:"description"`
	SpbillCreateIp string `json:"spbill_create_ip" xml:"spbill_create_ip"`
}

func (r WxEntPayRequest) IgnoredParamsForSign() []string {
	return []string{"sign_type"}
}

// 提现响应对象
type WxEntPayResult struct {
	BaseWxPayResult

	MchId          string `json:"mchid" xml:"mchid"`
	AppId          string `json:"mch_appid" xml:"mch_appid"`
	DeviceInfo     string `json:"device_info" xml:"device_info"`
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	PaymentNo      string `json:"payment_no" xml:"payment_no"`
	PaymentTime    string `json:"payment_time" xml:"payment_time"`
}
