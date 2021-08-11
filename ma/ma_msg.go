package ma

import "github.com/cliod/wx-go/common"

type WxMaMsgService interface {
	/* 发送客服消息
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/customerServiceMessage.send.html */
	SendKefuMsg(msg *WxMaKefuMessage) error
	/* 发送订阅消息
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html */
	SendSubscribeMsg(msg *WxMaSubscribeMessage) error
	/* 下发小程序和公众号统一的服务消息
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api/open-api/uniform-message/sendUniformMessage.html */
	SendUniformMsg(msg *WxMaUniformMessage) error
	/* 修改被分享的动态消息.
	   动态消息: https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share/updatable-message.html
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html */
	SetUpdatableMsg(msg *WxMaUpdatableMsg) error
	/* 创建被分享动态消息的 activity_id.
	   动态消息: https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share/updatable-message.html
	   文档地址：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html */
	CreateUpdatableMessageActivityId() ([]byte, error)
}

type WxMaMsgServiceImpl struct {
	service WxMaService
}

func newWxMaMsgService(service WxMaService) *WxMaMsgServiceImpl {
	return &WxMaMsgServiceImpl{
		service: service,
	}
}

func (m *WxMaMsgServiceImpl) SendSubscribeMsg(msg *WxMaSubscribeMessage) error {
	url := common.MaSubscribeMsgSendUrl
	return m.Send(url, msg)
}

func (m *WxMaMsgServiceImpl) SendUniformMsg(msg *WxMaUniformMessage) error {
	url := common.MaUniformMsgSendUrl
	return m.Send(url, msg)
}

func (m *WxMaMsgServiceImpl) SendKefuMsg(msg *WxMaKefuMessage) error {
	url := common.MaKefuMessageSendUrl
	return m.Send(url, msg)
}

func (m *WxMaMsgServiceImpl) SetUpdatableMsg(msg *WxMaUpdatableMsg) error {
	url := common.MaUpdatableMsgSendUrl
	return m.Send(url, msg)
}

func (m *WxMaMsgServiceImpl) CreateUpdatableMessageActivityId() ([]byte, error) {
	url := common.MaActivityIdCreateUrl
	return m.service.Get(url)
}

func (m *WxMaMsgServiceImpl) Send(url string, msg interface{}) error {
	var res common.WxCommonErr
	err := m.service.PostFor(&res, url, common.PostJsonContentType, msg)
	if err != nil {
		return err
	}
	return &res
}
