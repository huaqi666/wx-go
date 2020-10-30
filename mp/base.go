package mp

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
