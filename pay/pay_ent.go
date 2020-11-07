package pay

import (
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

type WxEntPayService interface {
	EntPay(*WxEntPayRequest) (*WxEntPayResult, error)
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

func (w *WxEntPayServiceImpl) EntPay(request *WxEntPayRequest) (*WxEntPayResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl

	var res WxEntPayResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEntPayBy(partnerTradeNo string) (*EntPayQueryResult, error) {
	return w.QueryEntPay(&EntPayQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryEntPay(request *EntPayQueryRequest) (*EntPayQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayQueryUrl

	var res EntPayQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) GetPublicKey() (string, error) {
	url := "https://fraud.mch.weixin.qq.com/risk/getpublickey"

	var res GetPublicKeyResult
	err := w.service.Do(url, &WxPayDefaultRequest{}, &res, true)
	return res.PubKey, err
}

func (w *WxEntPayServiceImpl) PayBank(request *EntPayBankRequest) (*EntPayBankResult, error) {
	var b []byte
	b, err := w.encryptRSA([]byte(request.EncBankNo))
	if err != nil {
		return nil, err
	}
	request.EncBankNo = string(b)
	b, err = w.encryptRSA([]byte(request.EncTrueName))
	if err != nil {
		return nil, err
	}
	request.EncTrueName = string(b)

	url := w.service.GetPayBaseUr() + common.EntPayBankUrl

	var res EntPayBankResult
	err = w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryPayBankBy(partnerTradeNo string) (*EntPayBankQueryResult, error) {
	return w.QueryPayBank(&EntPayBankQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryPayBank(request *EntPayBankQueryRequest) (*EntPayBankQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayQueryBankUrl

	var res EntPayBankQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) SendEnterpriseRedPack(request *EntPayRedPackRequest) (*EntPayRedPackResult, error) {
	//企业微信签名,需要在请求签名之前
	request.NonceStr = strconv.Itoa(time.Now().Nanosecond())
	request.WorkWxSign = SignEnt(request.ActName, request.MchBillNo, request.MchId, request.NonceStr, request.ReOpenid, request.WxAppId,
		w.service.GetWxPayConfig().EntPayKey, request.TotalAmount, MD5)

	url := w.service.GetPayBaseUr() + common.EntSendEnterpriseRedPackUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayRedPackResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEnterpriseRedPack(request *EntPayRedPackQueryRequest) (*EntPayRedPackQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntQueryEnterpriseRedPackUrl

	var res EntPayRedPackQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) encryptRSA(data []byte) ([]byte, error) {
	key, err := w.GetPublicKey()
	if err != nil {
		return nil, err
	}
	return util.RsaEncrypt(data, []byte(key))
}
