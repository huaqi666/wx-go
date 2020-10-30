package pay

import "wx-go/common"

type WxPayService interface {
	common.WxService

	PostKey()

	GetPayBaseUr() string
}
