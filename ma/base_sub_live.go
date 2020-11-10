package ma

// 直播间基础信息
type RoomInfoBase struct {
	// 直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	Name string `json:"name"`
	/* 背景图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取
	   直播间背景图，图片规则：建议像素1080*1920，大小不超过2M
	   请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；*/
	CoverImg string `json:"coverImg"`
	// 直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	StartTime uint64 `json:"startTime"`
	// 直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	EndTime uint64 `json:"endTime"`

	// 主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorName string `json:"anchorName"`
	// 主播微信号，如果未实名认证，需要先前往“小程序直播”小程序进行实名验证
	// 小程序二维码链接：https://res.wx.qq.com/op_res/BbVNeczA1XudfjVqCVoKgfuWe7e3aUhokktRVOqf_F0IqS6kYR--atCpVNUUC3zr
	AnchorWechat string `json:"anchorWechat"`

	/* 分享图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取
	   直播间分享图，图片规则：建议像素800*640，大小不超过1M；
	   请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；*/
	ShareImg string `json:"shareImg"`

	// 是否关闭点赞 【0：开启，1：关闭】（若关闭，直播开始后不允许开启）
	CloseLike uint64 `json:"closeLike"`
	// 是否关闭货架 【0：开启，1：关闭】（若关闭，直播开始后不允许开启）
	CloseGoods uint64 `json:"closeGoods"`
	// 是否关闭评论 【0：开启，1：关闭】（若关闭，直播开始后不允许开启）
	CloseComment uint64 `json:"closeComment"`
	// 是否关闭回放 【0：开启，1：关闭】默认关闭回放
	CloseReplay uint64 `json:"closeReplay"`
	// 是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	CloseShare uint64 `json:"closeShare"`
	// 是否关闭客服 【0：开启，1：关闭】 默认关闭客服
	CloseKf uint64 `json:"closeKf"`
	// 是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	IsFeedsPublic uint64 `json:"isFeedsPublic"`
	/* 购物直播频道封面图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，
	   请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html;
	   购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB；*/
	FeedsImg string `json:"feedsImg"`
}

type WxMaLiveGoodsResult struct {
	GoodsId       uint64 `json:"goods_id"`
	CoverImgUrl   string `json:"cover_img_url"`
	CoverImg      string `json:"cover_img"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	PriceType     uint64 `json:"price_type"`
	AuditStatus   uint64 `json:"audit_status"`
	Price         string `json:"price"`
	Price2        string `json:"price2"`
	ThirdPartyTag string `json:"third_party_tag"`
}

type WxMaLiveRoomInfosResult struct {
	Name         string                 `json:"name"`
	RoomId       uint64                 `json:"roomid"`
	CoverImg     string                 `json:"cover_img"`
	ShareImg     string                 `json:"share_img"`
	LiveStatus   string                 `json:"live_status"`
	StartTime    string                 `json:"start_time"`
	EndTime      string                 `json:"end_time"`
	AnchorName   string                 `json:"anchor_name"`
	AnchorWechat string                 `json:"anchor_wechat"`
	AnchorImg    string                 `json:"anchor_img"`
	TypeNum      uint64                 `json:"type"`
	ScreenType   uint64                 `json:"screen_type"`
	CloseLike    uint64                 `json:"close_like"`
	CloseGoods   uint64                 `json:"close_goods"`
	CloseComment uint64                 `json:"close_comment"`
	Goods        []*WxMaLiveGoodsResult `json:"goods"`
}

type WxMaLiveReplayResult struct {
	ExpireTime string `json:"expire_time"`
	CreateTime string `json:"create_time"`
	MediaUrl   string `json:"media_url"`
}
