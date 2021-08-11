package ma

import (
	"encoding/json"
	"github.com/cliod/wx-go/common"
	"time"
)

type WxMaLiveService interface {
	/* 创建直播间
	   调用此接口创建直播间，创建成功后将在直播间列表展示，调用额度：10000次/一天
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/studio-api.html#1 */
	CreateRoom(*WxLiveCreateRoomRequest) (*WxLiveCreateRoomResult, error)
	/* 编辑直播间
	   调用此接口创建直播间，创建成功后将在直播间列表展示，调用额度：10000次/一天
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/studio-api.html#6 */
	EditRoom(*WxLiveEditRoomRequest) (*WxLiveEditRoomResult, error)
	/* 删除直播间
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/studio-api.html#5 */
	DeleteRoom(roomId uint64) (*WxLiveDeleteRoomResult, error)
	// 获取直播间推流地址
	// 文档地址：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/studio-api.html#7
	GetPushUrl(roomId uint64) (*WxLiveGetPushUrlResult, error)
	// 获取直播间分享二维码
	// 文档地址：https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/studio-api.html#8
	GetSharedCode(roomId uint64, params string) (*WxLiveGetSharedUrlResult, error)
	// 获取直播房间列表.（分页）
	GetLiveInfos() ([]*WxMaLiveRoomInfosResult, error)
	// 获取所有直播间信息（没有分页直接获取全部）
	GetLiveInfo(start, limit int) (*WxMaLiveResult, error)
	// 获取直播房间回放数据信息.
	GetLiveReplay(action string, roomId uint64, start, limit int) (*WxMaLiveResult, error)
	/* 直播间导入商品
	   调用接口往指定直播间导入已入库的商品
	   调用频率, 调用额度：10000次/一天 */
	AddGoodsToRoom(roomId uint64, goodsIds []uint64) error
	// 添加管理直播间小助手
	AddAssistant(roomId uint64, users []WxMaLiveAssistantInfo) error
	// 修改直播间小助手昵称
	ModifyAssistant(roomId uint64, username, nickname string) error
	// 查询直播间小助手
	GetAssistantList(roomId uint64) (*WxMaAssistantResult, error)
	// 删除直播间小助手
	RemoveAssistant(roomId uint64, username string) error
}

type WxMaLiveServiceImpl struct {
	service WxMaService
}

func newWxMaLiveService(service WxMaService) *WxMaLiveServiceImpl {
	return &WxMaLiveServiceImpl{
		service: service,
	}
}

func (live *WxMaLiveServiceImpl) CreateRoom(request *WxLiveCreateRoomRequest) (*WxLiveCreateRoomResult, error) {
	url := common.MaCreateRoom

	var res WxLiveCreateRoomResult
	err := live.service.PostFor(&res, url, common.PostJsonContentType, request)
	return &res, err
}

func (live *WxMaLiveServiceImpl) EditRoom(request *WxLiveEditRoomRequest) (*WxLiveEditRoomResult, error) {
	url := common.MaEditRoom

	var res WxLiveEditRoomResult
	err := live.service.PostFor(&res, url, common.PostJsonContentType, request)
	return &res, err
}

func (live *WxMaLiveServiceImpl) DeleteRoom(roomId uint64) (*WxLiveDeleteRoomResult, error) {
	url := common.MaDeleteRoom

	request := make(map[string]interface{})
	request["id"] = roomId

	var res WxLiveDeleteRoomResult
	err := live.service.PostFor(&res, url, common.PostJsonContentType, request)
	return &res, err
}

func (live *WxMaLiveServiceImpl) GetPushUrl(roomId uint64) (*WxLiveGetPushUrlResult, error) {
	url := common.MaGetPushUrl

	var res WxLiveGetPushUrlResult
	err := live.service.GetFor(&res, url, roomId)
	return &res, err
}

func (live *WxMaLiveServiceImpl) GetSharedCode(roomId uint64, params string) (*WxLiveGetSharedUrlResult, error) {
	url := common.MaGetSharedCode

	var res WxLiveGetSharedUrlResult
	err := live.service.GetFor(&res, url, roomId, params)
	return &res, err
}

func (live *WxMaLiveServiceImpl) GetLiveInfos() ([]*WxMaLiveRoomInfosResult, error) {
	var arr []*WxMaLiveRoomInfosResult
	start := 0
	limit := 80
	total := uint64(0)
	var info WxMaLiveResult
	for uint64(len(arr)) <= total {
		if total != 0 && total <= uint64(start) {
			break
		}
		lf, err := live.GetLiveInfo(start, limit)
		if err != nil {
			continue
		}
		if lf == nil {
			return nil, err
		}
		info = *lf
		arr = append(arr, info.RoomInfos...)
		total = info.Total
		start = len(arr)

		time.Sleep(100 * time.Millisecond)
	}
	return arr, nil
}

func (live *WxMaLiveServiceImpl) GetLiveInfo(start, limit int) (*WxMaLiveResult, error) {
	var res *WxMaLiveResult
	b, err := live.getLiveInfo(start, limit, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}

func (live *WxMaLiveServiceImpl) GetLiveReplay(action string, roomId uint64, start, limit int) (*WxMaLiveResult, error) {
	var res *WxMaLiveResult

	param := map[string]interface{}{
		"action":  action,
		"room_id": roomId,
	}
	b, err := live.getLiveInfo(start, limit, param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}

func (live *WxMaLiveServiceImpl) AddGoodsToRoom(roomId uint64, goodsIds []uint64) error {

	param := map[string]interface{}{
		"ids":    goodsIds,
		"roomId": roomId,
	}

	url := common.MaAddGoods
	var res common.WxCommonErr
	err := live.service.PostFor(&res, url, common.PostJsonContentType, param)
	if err != nil {
		return err
	}
	return &res
}

// start 起始房间，0表示从第1个房间开始拉取
// limit 每次拉取的房间数量，建议100以内
func (live *WxMaLiveServiceImpl) getLiveInfo(start, limit int, param map[string]interface{}) ([]byte, error) {
	url := common.MaGetLiveInfo

	if param == nil {
		param = map[string]interface{}{
			"start": start,
			"limit": limit,
		}
	}

	return live.service.Post(url, common.PostJsonContentType, param)
}

func (live *WxMaLiveServiceImpl) AddAssistant(roomId uint64, users []WxMaLiveAssistantInfo) error {

	param := map[string]interface{}{
		"users":  users,
		"roomId": roomId,
	}

	url := common.MaAddGoods
	var res common.WxCommonErr
	err := live.service.PostFor(&res, url, common.PostJsonContentType, param)
	if err != nil {
		return err
	}
	return &res
}

func (live *WxMaLiveServiceImpl) ModifyAssistant(roomId uint64, username, nickname string) error {

	param := map[string]interface{}{
		"username": username,
		"nickname": nickname,
		"roomId":   roomId,
	}

	url := common.MaAddGoods
	var res common.WxCommonErr
	err := live.service.PostFor(&res, url, common.PostJsonContentType, param)
	if err != nil {
		return err
	}
	return &res
}

func (live *WxMaLiveServiceImpl) GetAssistantList(roomId uint64) (*WxMaAssistantResult, error) {
	param := map[string]interface{}{
		"roomId": roomId,
	}

	url := common.MaAddGoods
	var res WxMaAssistantResult
	err := live.service.PostFor(&res, url, common.PostJsonContentType, param)
	return &res, err
}

func (live *WxMaLiveServiceImpl) RemoveAssistant(roomId uint64, username string) error {

	param := map[string]interface{}{
		"username": username,
		"roomId":   roomId,
	}

	url := common.MaAddGoods
	var res common.WxCommonErr
	err := live.service.PostFor(&res, url, common.PostJsonContentType, param)
	if err != nil {
		return err
	}
	return &res
}
