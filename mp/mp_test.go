package mp

import (
	"github.com/cliod/wx-go"
	"github.com/cliod/wx-go/common"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Mp.AppId
	secret := c.Mp.Secret

	at, err := GetAccessToken(appId, secret)
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("token是 %s ", at.AccessToken)
	}

	ws, err := CreateJsapiSignatureOnce(appId, secret, "https://www.xxx.com")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("签名是 %s ", ws.Signature)
	}
}

func TestGetWxMpUserQueryParam(t *testing.T) {
	c := wx.GetConfig("./config.json")

	got := NewWxMpUserQuery(c.Mp.Openid)
	t.Log(got)
}

func TestWxMpMaterialServiceImpl_MaterialFileBatchgetUrl(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpMaterialService().MaterialNewsBatchGet(0, 1)
	t.Log(got, err)
}

func TestWxMpMaterialServiceImpl_MaterialNewsBatchGet(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpMaterialService().MaterialFileBatchGet(VIDEO, 0, 1)
	t.Log(got, err)
}

func TestWxMpQrcodeServiceImpl_QrcodeCreateLastTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpQrcodeService().QrcodeCreateLastTicket(QrLimitStrScene, "", 0, 0)
	t.Log(got, err)
}

func TestWxMpQrcodeServiceImpl_QrcodeCreateTmpTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpQrcodeService().QrcodeCreateTmpTicket(QrStrScene, "", 0, 0)
	t.Log(got, err)
}

func TestWxMpServiceImpl_GetTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetTicket(common.JSAPI)
	t.Log(got, err)
}

func TestWxMpUserServiceImpl_GetUserInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpUserService().GetUserInfo(c.Mp.Openid)
	t.Log(got, err)
}

func TestWxMpUserServiceImpl_GetUserInfoList(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	got, err := service.GetWxMpUserService().GetUserInfos("")
	t.Log(got, err)
}

func TestWxMpUserServiceImpl_GetUserList(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMpServiceWith(appId, secret)

	res, err := service.GetWxMpUserService().GetUserList("")
	t.Log(res, err)
}
