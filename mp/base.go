package mp

import "github.com/cliod/wx-go/common"

type ActionName string

func (a ActionName) String() string {
	return string(a)
}

const (
	QrScene         ActionName = "QR_SCENE"
	QrStrScene      ActionName = "QR_STR_SCENE"
	QrLimitScene    ActionName = "QR_LIMIT_SCENE"
	QrLimitStrScene ActionName = "QR_LIMIT_STR_SCENE"
)

type MaterialType string

func (m MaterialType) String() string {
	return string(m)
}

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
	Items      []*WxMaterialNewsBatchGetNewsItem `json:"item"`
}

type WxMpMaterialFileBatchGetResult struct {
	common.Err

	TotalCount uint64                            `json:"total_count"`
	ItemCount  uint64                            `json:"item_count"`
	Items      []*WxMaterialFileBatchGetNewsItem `json:"item"`
}
