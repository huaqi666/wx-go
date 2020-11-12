package mp

type ActionName string

const (
	QrScene         = "QR_SCENE"
	QrStrScene      = "QR_STR_SCENE"
	QrLimitScene    = "QR_LIMIT_SCENE"
	QrLimitStrScene = "QR_LIMIT_STR_SCENE"
)

// 二维码ticket
type WxMpQrCodeTicket struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"` // 如果为-1，说明是永久
	Url           string `json:"url"`
}
