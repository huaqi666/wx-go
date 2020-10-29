package common

// 配置的接口
type WxConfig interface {
	GetAppID() string
	GetSecret() string
	GetAccessToken() *AccessToken
	SetAccessToken(at *AccessToken)
}
