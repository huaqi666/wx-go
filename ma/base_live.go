package ma

import "github.com/cliod/wx-go/common"

type RoomInfo struct {
	Name         string   `json:"name"`
	CoverImg     string   `json:"coverImg"`
	ShareImg     string   `json:"shareImg"`
	LiveStatus   uint64   `json:"liveStatus"`
	StartTime    uint64   `json:"startTime"`
	EndTime      uint64   `json:"endTime"`
	AnchorName   string   `json:"anchorName"`
	AnchorWechat string   `json:"anchorWechat"`
	AnchorImg    string   `json:"anchorImg"`
	TypeNum      uint64   `json:"type"`
	ScreenType   uint64   `json:"screenType"`
	CloseLike    uint64   `json:"closeLike"`
	CloseGoods   uint64   `json:"closeGoods"`
	CloseComment uint64   `json:"closeComment"`
	CloseReplay  uint64   `json:"closeReplay"`
	CloseShare   uint64   `json:"closeShare"`
	CloseKf      uint64   `json:"closeKf"`
	Goods        []*Goods `json:"goods"`
}

type RoomInfoRequest struct {
	RoomInfo
}

type RoomInfos struct {
	RoomInfo
	Roomid uint64 `json:"roomid"`
}

type Goods struct {
	GoodsId     uint64 `json:"goods_id"`
	CoverImgUrl string `json:"cover_img_url"`
	Url         string `json:"url"`
	PriceType   uint64 `json:"price_type"`
	Price       string `json:"price"`
	Price2      string `json:"price_2"`
	Name        string `json:"name"`
	// 1, 2：表示是为api添加商品，否则是在MP添加商品
	ThirdPartyTag string `json:"third_party_tag"`
}

type RoomInfoResult struct {
	common.Err
	RoomId uint64 `json:"roomId"`
	// "小程序直播" 小程序码, 当主播微信号没有在 “小程序直播“ 小程序实名认证 返回该字段
	QrcodeUrl string `json:"qrcode_url"`
}

type WxMaLiveResult struct {
	common.Err

	Total     uint64       `json:"total"`
	AuditId   uint64       `json:"audit_id"`
	GoodsId   uint64       `json:"goods_id"`
	Goods     []*Goods     `json:"goods"`
	RoomInfos []*RoomInfos `json:"room_info"`
}
