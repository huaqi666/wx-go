package ma

type Config interface {
	GetAppID() string
	GetSecret() string
	GetAccessToken() *AccessToken
	SetAccessToken(at *AccessToken)
}

type ConfigImpl struct {
	appId       string
	secret      string
	AccessToken *AccessToken

	Token         string
	AesKey        string
	MsgDataFormat string
}

func newConfig(appId, secret string) Config {
	return &ConfigImpl{
		appId:  appId,
		secret: secret,
	}
}

func (c *ConfigImpl) GetAppID() string {
	return c.appId
}

func (c *ConfigImpl) GetSecret() string {
	return c.secret
}

func (c *ConfigImpl) GetAccessToken() *AccessToken {
	return c.AccessToken
}

func (c *ConfigImpl) SetAccessToken(at *AccessToken) {
	c.AccessToken = at
}
