package mp

import "github.com/cliod/wx-go/common"

type WxMpMaterialService interface {
	// MaterialNewsBatchGet 分页获取图文素材列表
	//   详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738734&token=&lang=zh_CN
	MaterialNewsBatchGet(offset, count int) (*WxMpMaterialNewsBatchGetResult, error)
	// MaterialFileBatchGet 分页获取其他媒体素材列表
	//   详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738734&token=&lang=zh_CN
	MaterialFileBatchGet(materialType MaterialType, offset, count int) (*WxMpMaterialFileBatchGetResult, error)
}

type WxMpMaterialServiceImpl struct {
	service WxMpService
}

func newWxMpMaterialService(service WxMpService) WxMpMaterialService {
	return &WxMpMaterialServiceImpl{
		service: service,
	}
}

func (m *WxMpMaterialServiceImpl) MaterialNewsBatchGet(offset, count int) (*WxMpMaterialNewsBatchGetResult, error) {
	var data = map[string]interface{}{}
	data["type"] = NEWS
	data["offset"] = offset
	data["count"] = count

	var res WxMpMaterialNewsBatchGetResult
	err := m.service.PostFor(&res, common.MpMaterialBatchGetUrl, common.PostJsonContentType, data)
	return &res, err
}

func (m *WxMpMaterialServiceImpl) MaterialFileBatchGet(materialType MaterialType, offset, count int) (*WxMpMaterialFileBatchGetResult, error) {
	var data = map[string]interface{}{}
	data["type"] = materialType
	data["offset"] = offset
	data["count"] = count

	var res WxMpMaterialFileBatchGetResult
	err := m.service.PostFor(&res, common.MpMaterialBatchGetUrl, common.PostJsonContentType, data)
	return &res, err
}
