package ma

const (
	AccessTokenUrl     = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	SessionInfoUrl     = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	QrCodeUrl          = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s"
	QrWxaCodeUrl       = "https://api.weixin.qq.com/wxa/getwxacode?access_token=%s"
	QrCodeUnlimitedUrl = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)
