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

type ReqInfo struct {
	TransactionId       string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo          string `json:"out_trade_no" xml:"out_trade_no"`
	RefundId            string `json:"refund_id" xml:"refund_id"`
	OutRefundNo         string `json:"out_refund_no" xml:"out_refund_no"`
	TotalFee            uint64 `json:"total_fee" xml:"total_fee"`
	SettlementTotalFee  uint64 `json:"settlement_total_fee" xml:"settlement_total_fee"`
	RefundFee           uint64 `json:"refund_fee" xml:"refund_fee"`
	SettlementRefundFee uint64 `json:"settlement_refund_fee" xml:"settlement_refund_fee"`
	RefundStatus        string `json:"refund_status" xml:"refund_status"`
	SuccessTime         string `json:"success_time" xml:"success_time"`
	RefundRecvAccout    string `json:"refund_recv_accout" xml:"refund_recv_accout"`
	RefundAccount       string `json:"refund_account" xml:"refund_account"`
	RefundRequestSource string `json:"refund_request_source" xml:"refund_request_source"`
}
