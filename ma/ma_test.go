package ma

import (
	"github.com/cliod/wx-go"
	"github.com/cliod/wx-go/common"
	"testing"
)

func TestGetAccessToken(t *testing.T) {

	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

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

func TestCheckAndGetUserInfo(t *testing.T) {
	got, err := CheckAndGetUserInfo("", "", "", "", "")
	t.Log(got, err)
}

func TestCheckSignature(t *testing.T) {
	got := CheckSignature("", "", "", "")
	t.Log(got)
}

func TestCheckUserInfo(t *testing.T) {
	b := CheckUserInfo("", "", "")
	t.Log(b)
}

func TestCreateJsapiSignature(t *testing.T) {
	got, err := CreateJsapiSignature("", "", "")
	t.Log(got, err)

}

func TestCreateJsapiSignatureBy(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	got, err := CreateJsapiSignatureOnce(appId, secret, "")
	t.Log(got, err)

}

func TestCreateWxaCodeUnlimited(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	got, err := CreateWxaCodeUnlimited(appId, secret, "", "pages/index")
	t.Log(got, err)
}

func TestGetJsapiTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	got, err := GetJsapiTicket(appId, secret)
	t.Log(got, err)
}

func TestGetPhoneNoInfo(t *testing.T) {
	got, err := GetPhoneNoInfo("", "", "")
	t.Log(got, err)

}

func TestGetShareInfo(t *testing.T) {
	got, err := GetShareInfo("", "", "")
	t.Log(got, err)

}

func TestGetTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	got, err := GetTicket(appId, secret, common.JSAPI)
	t.Log(got, err)

}

func TestJsCode2SessionInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	got, err := JsCode2SessionInfo(appId, secret, "")
	t.Log(got, err)
}

func TestWxMaJsapiServiceImpl_CreateJsapiSignature(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaJsapiService().CreateJsapiSignature("")
	t.Log(res, err)
}

func TestWxMaJsapiServiceImpl_GetTicket(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaJsapiService().GetTicket(common.JSAPI)
	t.Log(res, err)
}

func TestWxMaLiveServiceImpl_GetLiveInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaLiveService().GetLiveInfo(0, 1)
	t.Log(res, err)
}

func TestWxMaLiveServiceImpl_GetLiveInfos(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaLiveService().GetLiveInfos()
	t.Log(res, err)
}

func TestWxMaMsgServiceImpl_SendKefuMsg(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	err := service.GetWxMaMessageService().SendKefuMsg(&WxMaKefuMessage{})
	t.Log(err)
}

func TestWxMaMsgServiceImpl_SendSubscribeMsg(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	err := service.GetWxMaMessageService().SendSubscribeMsg(&WxMaSubscribeMessage{})
	t.Log(err)
}

func TestWxMaQrCodeServiceImpl_CreateQrcode(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaQrcodeService().CreateQrcode("")
	t.Log(res, err)
}

func TestWxMaQrCodeServiceImpl_CreateWxaCode(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaQrcodeService().CreateWxaCode("")
	t.Log(res, err)
}

func TestWxMaQrCodeServiceImpl_CreateWxaCodeUnlimited(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaQrcodeService().CreateWxaCodeUnlimited("", "pages/index")
	t.Log(res, err)
}

func TestWxMaServiceImpl_GetPaidUnionId(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetPaidUnionId("", "", "", "")
	t.Log(res, err)
}

func TestWxMaServiceImpl_JsCode2SessionInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.JsCode2SessionInfo("")
	t.Log(res, err)
}

func TestWxMaSubscribeServiceImpl_GetCategory(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaSubscribeService().GetCategory()
	t.Log(res, err)
}

func TestWxMaSubscribeServiceImpl_GetPubTemplateKeyWordsById(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaSubscribeService().GetPubTemplateKeyWordsById("")
	t.Log(res, err)
}

func TestWxMaSubscribeServiceImpl_GetPubTemplateTitleList(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaSubscribeService().GetPubTemplateTitleList([]string{"0"}, 0, 1)
	t.Log(res, err)
}

func TestWxMaSubscribeServiceImpl_GetTemplateList(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaSubscribeService().GetTemplateList()
	t.Log(res, err)
}

func TestWxMaUserServiceImpl_GetPhoneNoInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaUserService().GetPhoneInfo("", "", "")
	t.Log(res, err)
}

func TestWxMaUserServiceImpl_GetSessionInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaUserService().GetSessionInfo("")
	t.Log(res, err)
}

func TestWxMaUserServiceImpl_GetUserInfo(t *testing.T) {
	c := wx.GetConfig("./config.json")
	appId := c.Ma.AppId
	secret := c.Ma.Secret

	service := NewWxMaServiceBy(appId, secret)

	res, err := service.GetWxMaUserService().GetUserInfo("", "", "")
	t.Log(res, err)
}
