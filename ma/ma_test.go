package ma

import (
	"github.com/cliod/wx-go/common"
	"reflect"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	appId := "<appId>"
	secret := "<secret>"

	at, err := GetAccessToken(appId, secret)
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("token是 %s ", at.AccessToken)
	}

	ws, err := CreateJsapiSignatureBy(appId, secret, "https://www.xxx.com")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("签名是 %s ", ws.Signature)
	}
}

func TestAddData(t *testing.T) {
	type args struct {
		msg *WxMaSubscribeMessage
		arr []*Data
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCheckAndGetUserInfo(t *testing.T) {
	type args struct {
		sessionKey    string
		rawData       string
		encryptedData string
		signature     string
		ivStr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckAndGetUserInfo(tt.args.sessionKey, tt.args.rawData, tt.args.encryptedData, tt.args.signature, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAndGetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckAndGetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSignature(t *testing.T) {
	type args struct {
		token     string
		timestamp string
		nonce     string
		signature string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSignature(tt.args.token, tt.args.timestamp, tt.args.nonce, tt.args.signature); got != tt.want {
				t.Errorf("CheckSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUserInfo(t *testing.T) {
	type args struct {
		sessionKey string
		rawData    string
		signature  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUserInfo(tt.args.sessionKey, tt.args.rawData, tt.args.signature); got != tt.want {
				t.Errorf("CheckUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateJsapiSignature(t *testing.T) {
	type args struct {
		url    string
		appId  string
		ticket string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.WxJsapiSignature
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJsapiSignature(tt.args.url, tt.args.appId, tt.args.ticket)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJsapiSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateJsapiSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateJsapiSignatureBy(t *testing.T) {
	type args struct {
		appId  string
		secret string
		url    string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.WxJsapiSignature
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJsapiSignatureBy(tt.args.appId, tt.args.secret, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJsapiSignatureBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateJsapiSignatureBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateWxaCodeUnlimited(t *testing.T) {
	type args struct {
		appId  string
		secret string
		scene  string
		page   string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateWxaCodeUnlimited(tt.args.appId, tt.args.secret, tt.args.scene, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWxaCodeUnlimited() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWxaCodeUnlimited() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultCodeLineColor(t *testing.T) {
	tests := []struct {
		name string
		want CodeLineColor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultCodeLineColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultCodeLineColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccessToken1(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.AccessToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccessToken(tt.args.appId, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJsapiTicket(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJsapiTicket(tt.args.appId, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJsapiTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJsapiTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPhoneNoInfo(t *testing.T) {
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *PhoneNumberInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPhoneNoInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPhoneNoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhoneNoInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShareInfo(t *testing.T) {
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *WxMaShareInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetShareInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShareInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShareInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTicket(t *testing.T) {
	type args struct {
		appId  string
		secret string
		t      common.TicketType
	}
	tests := []struct {
		name    string
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTicket(tt.args.appId, tt.args.secret, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsCode2SessionInfo(t *testing.T) {
	type args struct {
		appId  string
		secret string
		jsCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *JsCode2SessionResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsCode2SessionInfo(tt.args.appId, tt.args.secret, tt.args.jsCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsCode2SessionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsCode2SessionInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMaConfig(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WxMaConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMaConfig(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMaConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMaService(t *testing.T) {
	type args struct {
		config WxMaConfig
	}
	tests := []struct {
		name string
		args args
		want WxMaService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMaService(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMaService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMaServiceBy(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WxMaService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMaServiceBy(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMaServiceBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMaSubscribeMessage(t *testing.T) {
	type args struct {
		toUser     string
		templateId string
		page       string
		arr        []*Data
	}
	tests := []struct {
		name string
		args args
		want *WxMaSubscribeMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMaSubscribeMessage(tt.args.toUser, tt.args.templateId, tt.args.page, tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMaSubscribeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_ExpireTicket(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	type args struct {
		ticketType common.TicketType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
		})
	}
}

func TestWxMaConfigImpl_GetAccessToken(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   *common.AccessToken
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetAccessToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetAesKey(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetAesKey(); got != tt.want {
				t.Errorf("GetAesKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetAppID(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetAppID(); got != tt.want {
				t.Errorf("GetAppID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetMsgDataFormat(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetMsgDataFormat(); got != tt.want {
				t.Errorf("GetMsgDataFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetSecret(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetSecret(); got != tt.want {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetTicket(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	type args struct {
		ticketType common.TicketType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *common.Ticket
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetTicket(tt.args.ticketType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetToken(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetToken(); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_GetWxTicket(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	tests := []struct {
		name   string
		fields fields
		want   common.WxTicket
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.GetWxTicket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_IsTicketExpired(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	type args struct {
		ticketType common.TicketType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
			if got := c.IsTicketExpired(tt.args.ticketType); got != tt.want {
				t.Errorf("IsTicketExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaConfigImpl_SetAccessToken(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	type args struct {
		at *common.AccessToken
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
		})
	}
}

func TestWxMaConfigImpl_UpdateTicket(t *testing.T) {
	type fields struct {
		appId         string
		secret        string
		AccessToken   *common.AccessToken
		Token         string
		AesKey        string
		MsgDataFormat string
		ticket        common.WxTicket
	}
	type args struct {
		ticketType common.TicketType
		ticket     *common.Ticket
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxMaConfigImpl{
				appId:         tt.fields.appId,
				secret:        tt.fields.secret,
				AccessToken:   tt.fields.AccessToken,
				Token:         tt.fields.Token,
				AesKey:        tt.fields.AesKey,
				MsgDataFormat: tt.fields.MsgDataFormat,
				ticket:        tt.fields.ticket,
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_CreateJsapiSignature(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.WxJsapiSignature
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.CreateJsapiSignature(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJsapiSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateJsapiSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_ForceGetCardApiTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		forceRefresh bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.ForceGetCardApiTicket(tt.args.forceRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForceGetCardApiTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceGetCardApiTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_ForceGetJsapiTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		forceRefresh bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.ForceGetJsapiTicket(tt.args.forceRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForceGetJsapiTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceGetJsapiTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_ForceGetTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		ticketType   common.TicketType
		forceRefresh bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.ForceGetTicket(tt.args.ticketType, tt.args.forceRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForceGetTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceGetTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_GetCardApiTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.GetCardApiTicket()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardApiTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCardApiTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_GetJsapiTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.GetJsapiTicket()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJsapiTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJsapiTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_GetTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		ticketType common.TicketType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.GetTicket(tt.args.ticketType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaJsapiServiceImpl_getTicket(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		ticketType common.TicketType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *common.Ticket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaJsapiServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.getTicket(tt.args.ticketType)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_AddGoodsToRoom(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		roomId   uint64
		goodsIds []uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			if err := l.AddGoodsToRoom(tt.args.roomId, tt.args.goodsIds); (err != nil) != tt.wantErr {
				t.Errorf("AddGoodsToRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_CreateRoom(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		info *WxLiveCreateRoomRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxLiveCreateRoomResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.CreateRoom(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRoom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_EditRoom(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		info *WxLiveEditRoomRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxLiveEditRoomResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.EditRoom(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditRoom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_GetLiveInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		start int
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaLiveResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.GetLiveInfo(tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_GetLiveInfos(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*WxMaLiveRoomInfosResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.GetLiveInfos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveInfos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveInfos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_GetLiveReplay(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		action string
		roomId uint64
		start  int
		limit  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaLiveResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.GetLiveReplay(tt.args.action, tt.args.roomId, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveReplay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveReplay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaLiveServiceImpl_getLiveInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		start int
		limit int
		param map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WxMaLiveServiceImpl{
				service: tt.fields.service,
			}
			got, err := l.getLiveInfo(tt.args.start, tt.args.limit, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLiveInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLiveInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_CreateUpdatableMessageActivityId(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			got, err := m.CreateUpdatableMessageActivityId()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUpdatableMessageActivityId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUpdatableMessageActivityId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_Send(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		url string
		msg interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			if err := m.Send(tt.args.url, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_SendKefuMsg(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		msg *WxMaKefuMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			if err := m.SendKefuMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendKefuMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_SendSubscribeMsg(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		msg *WxMaSubscribeMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			if err := m.SendSubscribeMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendSubscribeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_SendUniformMsg(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		msg *WxMaUniformMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			if err := m.SendUniformMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendUniformMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaMsgServiceImpl_SetUpdatableMsg(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		msg *WxMaUpdatableMsg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaMsgServiceImpl{
				service: tt.fields.service,
			}
			if err := m.SetUpdatableMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SetUpdatableMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_BytesToFile(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		filename string
		bytes    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			if err := q.BytesToFile(tt.args.filename, tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("BytesToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateQrcode(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := q.CreateQrcode(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQrcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateQrcode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateQrcodeBytes(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		path  string
		width int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			gotBytes, err := q.CreateQrcodeBytes(tt.args.path, tt.args.width)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQrcodeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("CreateQrcodeBytes() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateWxaCode(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := q.CreateWxaCode(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWxaCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWxaCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateWxaCodeBytes(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		path          string
		width         int
		autoColor     bool
		lineColor     CodeLineColor
		isTransparent bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			gotBytes, err := q.CreateWxaCodeBytes(tt.args.path, tt.args.width, tt.args.autoColor, tt.args.lineColor, tt.args.isTransparent)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWxaCodeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("CreateWxaCodeBytes() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateWxaCodeUnlimited(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		scene string
		path  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := q.CreateWxaCodeUnlimited(tt.args.scene, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWxaCodeUnlimited() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWxaCodeUnlimited() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaQrCodeServiceImpl_CreateWxaCodeUnlimitedBytes(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		scene         string
		path          string
		width         int
		autoColor     bool
		lineColor     CodeLineColor
		isTransparent bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &WxMaQrCodeServiceImpl{
				service: tt.fields.service,
			}
			gotBytes, err := q.CreateWxaCodeUnlimitedBytes(tt.args.scene, tt.args.path, tt.args.width, tt.args.autoColor, tt.args.lineColor, tt.args.isTransparent)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWxaCodeUnlimitedBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("CreateWxaCodeUnlimitedBytes() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestWxMaServiceImpl_CheckSignature(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		timestamp string
		nonce     string
		signature string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.CheckSignature(tt.args.timestamp, tt.args.nonce, tt.args.signature); got != tt.want {
				t.Errorf("CheckSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetPaidUnionId(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		openid        string
		transactionId string
		mchId         string
		outTradeNo    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaUnionIdResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			got, err := s.GetPaidUnionId(tt.args.openid, tt.args.transactionId, tt.args.mchId, tt.args.outTradeNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaidUnionId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPaidUnionId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaConfig(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaJsapiService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaJsapiService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaJsapiService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaJsapiService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaMessageService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaMsgService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaMessageService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaMessageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaQrcodeService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaQrcodeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaQrcodeService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaQrcodeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaShareService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaShareService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaShareService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaShareService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaSubscribeService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaSubscribeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaSubscribeService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaSubscribeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_GetWxMaUserService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMaUserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			if got := s.GetWxMaUserService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMaUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_JsCode2SessionInfo(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		jsCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JsCode2SessionResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
			got, err := s.JsCode2SessionInfo(tt.args.jsCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsCode2SessionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsCode2SessionInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaConfig(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		config WxMaConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaJsapiService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		jsapiService WxMaJsapiService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaMsgService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		msgService WxMaMsgService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaQrcodeService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		qrcodeService WxMaQrcodeService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaShareService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		shareService WxMaShareService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaSubscribeService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		subscribeService WxMaSubscribeService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaServiceImpl_SetWxMaUserService(t *testing.T) {
	type fields struct {
		WxServiceImpl    common.WxServiceImpl
		config           WxMaConfig
		userService      WxMaUserService
		qrCodeService    WxMaQrcodeService
		subscribeService WxMaSubscribeService
		shareService     WxMaShareService
		msgService       WxMaMsgService
		liveService      WxMaLiveService
		jsapiService     WxMaJsapiService
	}
	type args struct {
		userService WxMaUserService
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaServiceImpl{
				WxServiceImpl:    tt.fields.WxServiceImpl,
				config:           tt.fields.config,
				userService:      tt.fields.userService,
				qrCodeService:    tt.fields.qrCodeService,
				subscribeService: tt.fields.subscribeService,
				shareService:     tt.fields.shareService,
				msgService:       tt.fields.msgService,
				liveService:      tt.fields.liveService,
				jsapiService:     tt.fields.jsapiService,
			}
		})
	}
}

func TestWxMaShareServiceImpl_GetShareInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaShareInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMaShareServiceImpl{
				service: tt.fields.service,
			}
			got, err := s.GetShareInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShareInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShareInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeMessage_AddData(t *testing.T) {
	type fields struct {
		ToUser           string
		TemplateId       string
		Page             string
		Data             []*Data
		MiniProgramState MiniProgramState
		Lang             MiniProgramLang
	}
	type args struct {
		data *Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WxMaSubscribeMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMaSubscribeMessage{
				ToUser:           tt.fields.ToUser,
				TemplateId:       tt.fields.TemplateId,
				Page:             tt.fields.Page,
				Data:             tt.fields.Data,
				MiniProgramState: tt.fields.MiniProgramState,
				Lang:             tt.fields.Lang,
			}
			if got := m.AddData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_AddTemplate(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		id         string
		sceneDesc  string
		keywordIds []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaAddTemplateResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			got, err := i.AddTemplate(tt.args.id, tt.args.sceneDesc, tt.args.keywordIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_DelTemplate(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		templateId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			if err := i.DelTemplate(tt.args.templateId); (err != nil) != tt.wantErr {
				t.Errorf("DelTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_GetCategory(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    *WxMaCategoryListResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			got, err := i.GetCategory()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_GetPubTemplateKeyWordsById(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaPubTemplateKeywordListResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			got, err := i.GetPubTemplateKeyWordsById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPubTemplateKeyWordsById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPubTemplateKeyWordsById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_GetPubTemplateTitleList(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		ids   []string
		start int
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMaPubTemplateTitleListResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			got, err := i.GetPubTemplateTitleList(tt.args.ids, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPubTemplateTitleList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPubTemplateTitleList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_GetTemplateList(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	tests := []struct {
		name    string
		fields  fields
		want    *WxMaTemplateListResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			got, err := i.GetTemplateList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTemplateList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTemplateList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaSubscribeServiceImpl_SendMsg(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		msg *WxMaSubscribeMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WxMaSubscribeServiceImpl{
				service: tt.fields.service,
			}
			if err := i.SendMsg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxMaUserServiceImpl_CheckUserInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		sessionKey string
		rawData    string
		signature  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WxMaUserServiceImpl{
				service: tt.fields.service,
			}
			if got := u.CheckUserInfo(tt.args.sessionKey, tt.args.rawData, tt.args.signature); got != tt.want {
				t.Errorf("CheckUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaUserServiceImpl_GetPhoneNoInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PhoneNumberInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WxMaUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := u.GetPhoneNoInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPhoneNoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPhoneNoInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaUserServiceImpl_GetSessionInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		jsCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JsCode2SessionResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WxMaUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := u.GetSessionInfo(tt.args.jsCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSessionInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaUserServiceImpl_GetUserInfo(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		sessionKey    string
		encryptedData string
		ivStr         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WxMaUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := u.GetUserInfo(tt.args.sessionKey, tt.args.encryptedData, tt.args.ivStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMaUserServiceImpl_SetUserStorage(t *testing.T) {
	type fields struct {
		service WxMaService
	}
	type args struct {
		kvMap      map[string]string
		sessionKey string
		openid     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &WxMaUserServiceImpl{
				service: tt.fields.service,
			}
			if err := u.SetUserStorage(tt.args.kvMap, tt.args.sessionKey, tt.args.openid); (err != nil) != tt.wantErr {
				t.Errorf("SetUserStorage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newWxMaConfig(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want *WxMaConfigImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaConfig(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaJsapiService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaJsapiServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaJsapiService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaJsapiService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaLiveService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaLiveServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaLiveService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaLiveService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaMsgService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaMsgServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaMsgService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaMsgService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaQrcodeService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaQrCodeServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaQrcodeService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaQrcodeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaService(t *testing.T) {
	type args struct {
		config WxMaConfig
	}
	tests := []struct {
		name string
		args args
		want *WxMaServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaService(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaShareService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaShareServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaShareService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaShareService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaSubscribeService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaSubscribeServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaSubscribeService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaSubscribeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMaUserService(t *testing.T) {
	type args struct {
		service WxMaService
	}
	tests := []struct {
		name string
		args args
		want *WxMaUserServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMaUserService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMaUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}
