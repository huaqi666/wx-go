package mp

import "wx-go/common"

type WxMpConfig interface {
	common.WxConfig

	GetToken() string
	GetAesKey() string
}

type WxMpConfigImpl struct {
	appId       string
	secret      string
	AccessToken *common.AccessToken

	Token  string
	AesKey string
}

func newWxMpConfig(appId, secret string) WxMpConfig {
	return &WxMpConfigImpl{
		appId:  appId,
		secret: secret,
	}
}

func (c *WxMpConfigImpl) GetAppID() string {
	return c.appId
}

func (c *WxMpConfigImpl) GetSecret() string {
	return c.secret
}

func (c *WxMpConfigImpl) GetAccessToken() *common.AccessToken {
	return c.AccessToken
}

func (c *WxMpConfigImpl) SetAccessToken(at *common.AccessToken) {
	c.AccessToken = at
}

func (c *WxMpConfigImpl) GetToken() string {
	return c.Token
}

func (c *WxMpConfigImpl) GetAesKey() string {
	return c.AesKey
}
