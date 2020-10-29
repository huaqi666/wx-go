package ma

import "wx-go/common"

type WxMaConfig interface {
	common.WxConfig
}

type WxMaConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token         string
	AesKey        string
	MsgDataFormat string
}

func newWxMaConfig(appId, secret string) *WxMaConfigImpl {
	return &WxMaConfigImpl{
		appId:  appId,
		secret: secret,
	}
}

func (c *WxMaConfigImpl) GetAppID() string {
	return c.appId
}

func (c *WxMaConfigImpl) GetSecret() string {
	return c.secret
}

func (c *WxMaConfigImpl) GetAccessToken() *common.AccessToken {
	return c.AccessToken
}

func (c *WxMaConfigImpl) SetAccessToken(at *common.AccessToken) {
	c.AccessToken = at
}
