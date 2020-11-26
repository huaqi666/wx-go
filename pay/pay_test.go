package pay

import (
	"github.com/cliod/wx-go"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	wxConf = wx.GetConfig("./config.json")
	config = wxConf.Pay
)

func TestWxEntPayServiceImpl_EntPay(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.EntPay(&WxEntPayRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_GetPublicKey(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.GetPublicKey()
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_PayBank(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.PayBank(&EntPayBankRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_QueryEntPay(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_QueryEntPayBy(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPayBy("")
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_QueryEnterpriseRedPack(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_QueryPayBank(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_QueryPayBankBy(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_SendEnterpriseRedPack(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxEntPayServiceImpl_encryptRSA(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_CloseOrder(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_CloseOrderBy(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_Do(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_GetEntPayService(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_GetPayBaseUr(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_GetSandboxSignKey(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_GetWxPayConfig(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_ParseOrderNotifyResult(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_ParseRefundNotifyResult(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_ParseScanPayNotifyResult(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_Post(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_PostFor(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_PostKey(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_QueryOrder(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_QueryOrderBy(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_Refund(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_RefundQuery(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_RefundQueryV2(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_RefundV2(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_Sign(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	res, err := service.CloseOrder(&WxPayOrderCloseRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_UnifyOrder(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")
	ser := service.GetEntPayService()

	res, err := ser.QueryEntPay(&EntPayQueryRequest{})
	t.Log(res, err)
}

func TestWxPayV2ServiceImpl_UnifyPay(t *testing.T) {
	service := NewWxPayServiceBy(config.AppId, config.MchId, config.MchKey, "http://www.xxx.cn/notify", "")

	s := strconv.Itoa(time.Now().Nanosecond()) + strconv.Itoa(rand.Intn(999999))

	res, err := service.UnifyPay(&WxPayUnifiedOrderRequest{
		TotalFee:   100,
		Openid:     config.Openid,
		OutTradeNo: s,
		Body:       "测试数据",
	})
	t.Log(res, err)
}
