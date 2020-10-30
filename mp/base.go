package mp

import (
	"time"
	"wx-go/common"
)

type TicketType string
type ActionName string

const (
	JSAPI  TicketType = "jsapi"
	SDK    TicketType = "2"
	WxCard TicketType = "wx_card"
)

const (
	QrScene         = "QR_SCENE"
	QrStrScene      = "QR_STR_SCENE"
	QrLimitScene    = "QR_LIMIT_SCENE"
	QrLimitStrScene = "QR_LIMIT_STR_SCENE"
)

// 授权页ticket
type Ticket struct {
	common.Err
	Ticket    string    `json:"ticket"`
	ExpiresIn uint64    `json:"expires_in"`
	Time      time.Time `json:"time"`
	Type      string    `json:"type"`
}

// jspai signature.
type WxJsapiSignature struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}
