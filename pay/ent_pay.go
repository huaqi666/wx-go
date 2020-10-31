package pay

type WxEntPayService interface {
	EntPay(*EntPayRequest) *EntPayResult
}

type WxEntPayServiceImpl struct {
	service WxPayService
}

func newWxEntPayService(service WxPayService) *WxEntPayServiceImpl {
	return &WxEntPayServiceImpl{
		service: service,
	}
}

func (w *WxEntPayServiceImpl) EntPay(*EntPayRequest) *EntPayResult {
	return nil
}
