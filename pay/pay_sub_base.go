package pay

type WxPayOrderCoupon struct {
	CouponType string `json:"coupon_type" xml:"coupon_type"`
	CouponId   string `json:"coupon_id" xml:"coupon_id"`
	CouponFee  uint64 `json:"coupon_fee" xml:"coupon_fee"`
}

type WxPayRefundCoupon struct {
	CouponRefundId  string `json:"coupon_refund_id" xml:"coupon_refund_id"`
	CouponType      string `json:"coupon_type" xml:"coupon_type"`
	CouponRefundFee uint64 `json:"coupon_refund_fee" xml:"coupon_refund_fee"`
}

type WxPayRefundPromotionDetail struct {
	PromotionId  string `json:"promotion_id"`
	Scope        string `json:"scope"`
	TypeStr      string `json:"type_str"`
	RefundAmount string `json:"refund_amount"`
	GoodsDetails struct {
		GoodsId        string `json:"goods_id"`
		RefundAmount   uint64 `json:"refund_amount"`
		RefundQuantity uint64 `json:"refund_quantity"`
		Price          uint64 `json:"price"`
	} `json:"goods_details"`
}

type WxPayRefundRecord struct {
	OutRefundNo         string               `json:"out_refund_no"`
	RefundId            string               `json:"refund_id"`
	RefundChannel       string               `json:"refund_channel"`
	RefundFee           uint64               `json:"refund_fee"`
	SettlementRefundFee uint64               `json:"settlement_refund_fee"`
	RefundAccount       string               `json:"refund_account"`
	CouponRefundFee     uint64               `json:"coupon_refund_fee"`
	CouponRefundCount   uint64               `json:"coupon_refund_count"`
	RefundCoupons       []*WxPayRefundCoupon `json:"refund_coupons"`
	RefundStatus        string               `json:"refund_status"`
	RefundRecvAccount   string               `json:"refund_recv_account"`
	RefundSuccessTime   string               `json:"refund_success_time"`
}
