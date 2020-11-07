package mp

import "github.com/cliod/wx-go/common"

type WxMpQrcodeService interface {
	// 换取临时二维码ticket
	// 详情请见: https://mp.weixin.qq.com/wiki?action=doc&id=mp1443433542&t=0.9274944716856435
	// sceneStr 和 sceneId 二选一就行
	QrcodeCreateTmpTicket(actionName ActionName, sceneStr string, sceneId, expireSeconds int64) (*WxMpQrCodeTicket, error)
	// 换取永久二维码ticket
	// 详情请见: https://mp.weixin.qq.com/wiki?action=doc&id=mp1443433542&t=0.9274944716856435
	QrcodeCreateLastTicket(actionName ActionName, sceneStr string, sceneId, expireSeconds int64) (*WxMpQrCodeTicket, error)
}

// 二维码ticket
type WxMpQrCodeTicket struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"` // 如果为-1，说明是永久
	Url           string `json:"url"`
}

type WxMpQrcodeServiceImpl struct {
	service WxMpService
}

func newWxMpQrcodeService(service WxMpService) *WxMpQrcodeServiceImpl {
	return &WxMpQrcodeServiceImpl{
		service: service,
	}
}

func (r *WxMpQrcodeServiceImpl) QrcodeCreateTmpTicket(actionName ActionName, sceneStr string, sceneId, expireSeconds int64) (*WxMpQrCodeTicket, error) {
	if expireSeconds <= 0 {
		expireSeconds = 30
	}
	// 时间不能超过30天
	if expireSeconds > 2592000 {
		expireSeconds = 2592000
	}
	return r.getQrCodeTicket(actionName, sceneStr, sceneId, expireSeconds)
}

func (r *WxMpQrcodeServiceImpl) QrcodeCreateLastTicket(actionName ActionName, sceneStr string, sceneId, expireSeconds int64) (*WxMpQrCodeTicket, error) {
	return r.getQrCodeTicket(actionName, sceneStr, sceneId, expireSeconds)
}

func (r *WxMpQrcodeServiceImpl) getQrCodeTicket(actionName ActionName, sceneStr string, sceneId, expireSeconds int64) (*WxMpQrCodeTicket, error) {
	var (
		data       = map[string]interface{}{}
		scene      = map[string]interface{}{}
		actionInfo = map[string]interface{}{}
	)
	data["action_name"] = actionName
	if expireSeconds > 0 {
		data["expire_seconds"] = expireSeconds
	}

	if sceneStr != "" {
		scene["scene_str"] = sceneStr
	}
	if sceneId > 0 {
		scene["scene_id"] = sceneId
	}

	actionInfo["scene"] = scene
	data["action_info"] = actionInfo

	at, err := r.service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	var res WxMpQrCodeTicket
	err = r.service.PostFor(&res, common.MpQrcodeUrl, common.PostJsonContentType, data, at.AccessToken)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
