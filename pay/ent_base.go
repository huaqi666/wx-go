package pay

type EntPayQueryRequest struct {
	BaseWxPayRequest

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
}

type EntPayQueryResult struct {
}

type EntPayBankRequest struct {
	BaseWxPayRequest
}

type EntPayBankResult struct {
	BaseWxPayResult
}

type EntPayBankQueryRequest struct {
	BaseWxPayRequest

	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
}

type EntPayBankQueryResult struct {
	BaseWxPayResult
}

type EntPayRedPackRequest struct {
	BaseWxPayRequest
}

type EntPayRedPackResult struct {
}

type EntPayRedPackQueryRequest struct {
	BaseWxPayRequest
}

type EntPayRedPackQueryResult struct {
}

type GetPublicKeyResult struct {
	BaseWxPayResult

	PubKey string `json:"pub_key" xml:"pub_key"`
}
