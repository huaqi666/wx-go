package pay

import "github.com/cliod/wx-go/common"

type WxEntPayService interface {
	EntPay(*EntPayRequest) (*EntPayResult, error)
	QueryEntPayBy(partnerTradeNo string) (*EntPayQueryResult, error)
	QueryEntPay(request *EntPayQueryRequest) (*EntPayQueryResult, error)

	GetPublicKey() (string, error)

	PayBank(request *EntPayBankRequest) (*EntPayBankResult, error)
	QueryPayBankBy(partnerTradeNo string) (*EntPayBankQueryResult, error)
	QueryPayBank(request *EntPayBankQueryRequest) (*EntPayBankQueryResult, error)

	SendEnterpriseRedPack(request *EntPayRedPackRequest) (*EntPayRedPackResult, error)
	QueryEnterpriseRedPack(request *EntPayRedPackQueryRequest) (*EntPayRedPackQueryResult, error)
}

type WxEntPayServiceImpl struct {
	service WxPayService
}

func newWxEntPayService(service WxPayService) *WxEntPayServiceImpl {
	return &WxEntPayServiceImpl{
		service: service,
	}
}

func (w *WxEntPayServiceImpl) EntPay(request *EntPayRequest) (*EntPayResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEntPayBy(partnerTradeNo string) (*EntPayQueryResult, error) {
	return w.QueryEntPay(&EntPayQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryEntPay(request *EntPayQueryRequest) (*EntPayQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayQueryUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayQueryResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}

func (w *WxEntPayServiceImpl) GetPublicKey() (string, error) {
	url := "https://fraud.mch.weixin.qq.com/risk/getpublickey"

	var request BaseWxPayRequest
	request.Sign = w.service.Sign(request, request.SignType)

	var res GetPublicKeyResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return res.PubKey, err
}

func (w *WxEntPayServiceImpl) PayBank(request *EntPayBankRequest) (*EntPayBankResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayBankResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryPayBankBy(partnerTradeNo string) (*EntPayBankQueryResult, error) {
	return w.QueryPayBank(&EntPayBankQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryPayBank(request *EntPayBankQueryRequest) (*EntPayBankQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayBankQueryResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}

func (w *WxEntPayServiceImpl) SendEnterpriseRedPack(request *EntPayRedPackRequest) (*EntPayRedPackResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayRedPackResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEnterpriseRedPack(request *EntPayRedPackQueryRequest) (*EntPayRedPackQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayRedPackQueryResult
	err := w.service.PostKeyFor(&res, url, common.PostXml, request)
	return &res, err
}
