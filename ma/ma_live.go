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
	CreateRoom(info *RoomInfoRequest) (*RoomInfoResult, error)
	// 获取直播房间列表.（分页）
	GetLiveInfos() ([]*RoomInfos, error)
	// 获取所有直播间信息（没有分页直接获取全部）
	GetLiveInfo(start, limit int) (*WxMaLiveResult, error)
	// 获取直播房间回放数据信息.
	GetLiveReplay(action string, roomId uint64, start, limit int) (*WxMaLiveResult, error)
	/* 直播间导入商品
	   调用接口往指定直播间导入已入库的商品
	   调用频率, 调用额度：10000次/一天 */
	AddGoodsToRoom(roomId uint64, goodsIds []uint64) error
}

type WxMaLiveServiceImpl struct {
	service WxMaService
}

func newWxMaLiveService(service WxMaService) *WxMaLiveServiceImpl {
	return &WxMaLiveServiceImpl{
		service: service,
	}
}

func (l *WxMaLiveServiceImpl) CreateRoom(info *RoomInfoRequest) (*RoomInfoResult, error) {
	url := common.MaCreateRoom

	var res RoomInfoResult
	err := l.service.PostFor(&res, url, common.PostJsonContentType, info)
	return &res, err
}

func (l *WxMaLiveServiceImpl) GetLiveInfos() ([]*RoomInfos, error) {
	var arr []*RoomInfos
	start := 0
	limit := 80
	total := uint64(0)
	var info WxMaLiveResult
	for uint64(len(arr)) <= total {
		if total != 0 && total <= uint64(start) {
			break
		}
		lf, err := l.GetLiveInfo(start, limit)
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

func (l *WxMaLiveServiceImpl) GetLiveInfo(start, limit int) (*WxMaLiveResult, error) {
	var res *WxMaLiveResult
	b, err := l.getLiveInfo(start, limit, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}

func (l *WxMaLiveServiceImpl) GetLiveReplay(action string, roomId uint64, start, limit int) (*WxMaLiveResult, error) {
	var res *WxMaLiveResult

	param := map[string]interface{}{
		"action":  action,
		"room_id": roomId,
	}
	b, err := l.getLiveInfo(start, limit, param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}

func (l *WxMaLiveServiceImpl) AddGoodsToRoom(roomId uint64, goodsIds []uint64) error {

	param := map[string]interface{}{
		"ids":    goodsIds,
		"roomId": roomId,
	}

	url := common.MaAddGoods
	var res common.Err
	err := l.service.PostFor(&res, url, common.PostJsonContentType, param)
	if err != nil {
		return err
	}
	return &res
}

// start 起始房间，0表示从第1个房间开始拉取
// limit 每次拉取的房间数量，建议100以内
func (l *WxMaLiveServiceImpl) getLiveInfo(start, limit int, param map[string]interface{}) ([]byte, error) {
	url := common.MaGetLiveInfo

	if param == nil {
		param = map[string]interface{}{
			"start": start,
			"limit": limit,
		}
	}

	return l.service.Post(url, common.PostJsonContentType, param)
}
