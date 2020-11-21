package mp

import "github.com/cliod/wx-go/common"

type ActionName string

const (
	QrScene         = "QR_SCENE"
	QrStrScene      = "QR_STR_SCENE"
	QrLimitScene    = "QR_LIMIT_SCENE"
	QrLimitStrScene = "QR_LIMIT_STR_SCENE"
)

type MaterialType string

const (
	NEWS  MaterialType = "news"
	VOICE MaterialType = "voice"
	IMAGE MaterialType = "image"
	VIDEO MaterialType = "video"
)

// 二维码ticket
type WxMpQrCodeTicket struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"` // 如果为-1，说明是永久
	Url           string `json:"url"`
}

type WxMpMaterialNewsBatchGetResult struct {
	common.Err

	TotalCount uint64                            `json:"total_count"`
	ItemCount  uint64                            `json:"item_count"`
	Items      []*WxMaterialNewsBatchGetNewsItem `json:"items"`
}

type WxMpMaterialFileBatchGetResult struct {
	common.Err

	TotalCount uint64                            `json:"total_count"`
	ItemCount  uint64                            `json:"item_count"`
	Item       []*WxMaterialFileBatchGetNewsItem `json:"items"`
}
