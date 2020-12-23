package ma

import "github.com/cliod/wx-go/common"

// 直播间信息
type WxLiveCreateRoomRequest struct {
	RoomInfoBase
	/* 主播副号微信号
	   如果未实名认证，需要先前往“小程序直播”小程序进行实名验证
	   小程序二维码链接：https://res.wx.qq.com/op_res/BbVNeczA1XudfjVqCVoKgfuWe7e3aUhokktRVOqf_F0IqS6kYR--atCpVNUUC3zr */
	SubAnchorWechat string `json:"subAnchorWechat"`
	// 直播间类型 【1: 推流，0：手机直播】
	TypeNum uint64 `json:"type"`
	// 横屏、竖屏 【1：横屏，0：竖屏】（横屏：视频宽高比为16:9、4:3、1.85:1 ；竖屏：视频宽高比为9:16、2:3）
	ScreenType uint64 `json:"screenType"`
}

type WxLiveCreateRoomResult struct {
	common.Err
	// 直播间id
	RoomId uint64 `json:"roomId"`
	// "小程序直播" 小程序码, 当主播微信号没有在 “小程序直播“ 小程序实名认证 返回该字段
	QrcodeUrl string `json:"qrcode_url"`
}

type WxLiveEditRoomRequest struct {
	RoomInfoBase
	// 直播间id
	Id uint64 `json:"id"`
}

// 编辑结果
type WxLiveEditRoomResult struct {
	common.Err
}

// 删除结果
type WxLiveDeleteRoomResult struct {
	common.Err
}

// 删除结果
type WxLiveGetPushUrlResult struct {
	common.Err

	// 地址
	PushAddr string `json:"pushAddr"`
}

// 删除结果
type WxLiveGetSharedUrlResult struct {
	common.Err

	// 地址
	CdnUrl string `json:"cdnUrl"`
}

// 直播信息-直播间操作返回结果
type WxMaLiveResult struct {
	common.Err

	Total   uint64 `json:"total"`
	AuditId uint64 `json:"auditId"`
	GoodsId uint64 `json:"goodsId"`
	// 商品列表
	Goods []*WxMaLiveGoodsResult `json:"goods"`
	// 直播间列表
	RoomInfos []*WxMaLiveRoomInfosResult `json:"room_info"`
	// 获取回放源视频列表
	LiveReplay []*WxMaLiveReplayResult `json:"live_replay"`
}

// 直播间小助手用户信息
type WxMaLiveAssistantInfo struct {
	Timestamp uint64 `json:"timestamp"` // 修改时间
	HeadImg   string `json:"headimg"`   // 头像
	Username  string `json:"username"`  // 用户微信号
	Nickname  string `json:"nickname"`  // 用户昵称
	Alias     string `json:"alias"`     // 微信号
	Openid    string `json:"openid"`    // openid
}

// 直播间小助手用户信息
type WxMaAssistantResult struct {
	common.Err

	Count    uint64           `json:"count"`
	MaxCount uint64           `json:"maxCount"`
	List     []*WxMaAssistant `json:"list"`
}

type WxMaAssistant struct {
	Timestamp uint64 `json:"timestamp"` // 修改时间
	HeadImg   string `json:"headimg"`   // 头像
	Nickname  string `json:"nickname"`  // 用户昵称
	Alias     string `json:"alias"`     // 微信号
	Openid    string `json:"openid"`    // openid
}
