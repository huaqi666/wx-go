package mp

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

func TestGetWxMpUserQueryParam(t *testing.T) {
	type args struct {
		openidArr []string
	}
	tests := []struct {
		name string
		args args
		want []*WxMpUserQueryParam
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWxMpUserQueryParam(tt.args.openidArr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMpUserQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMpConfig(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WxMpConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMpConfig(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMpConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMpService(t *testing.T) {
	type args struct {
		config WxMpConfig
	}
	tests := []struct {
		name string
		args args
		want WxMpService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMpService(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMpService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWxMpServiceBy(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WxMpService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWxMpServiceBy(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWxMpServiceBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_ExpireTicket(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
		})
	}
}

func TestWxMpConfigImpl_GetAccessToken(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetAccessToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetAesKey(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetAesKey(); got != tt.want {
				t.Errorf("GetAesKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetAppID(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetAppID(); got != tt.want {
				t.Errorf("GetAppID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetSecret(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetSecret(); got != tt.want {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetTicket(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetTicket(tt.args.ticketType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetToken(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetToken(); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_GetWxTicket(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.GetWxTicket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_IsTicketExpired(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
			if got := c.IsTicketExpired(tt.args.ticketType); got != tt.want {
				t.Errorf("IsTicketExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpConfigImpl_SetAccessToken(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
		})
	}
}

func TestWxMpConfigImpl_UpdateTicket(t *testing.T) {
	type fields struct {
		appId       string
		secret      string
		AccessToken *common.AccessToken
		Token       string
		AesKey      string
		ticket      common.WxTicket
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
			c := &WxMpConfigImpl{
				appId:       tt.fields.appId,
				secret:      tt.fields.secret,
				AccessToken: tt.fields.AccessToken,
				Token:       tt.fields.Token,
				AesKey:      tt.fields.AesKey,
				ticket:      tt.fields.ticket,
			}
		})
	}
}

func TestWxMpMaterialServiceImpl_MaterialFileBatchgetUrl(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		materialType MaterialType
		offset       int
		count        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpMaterialFileBatchGetResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMpMaterialServiceImpl{
				service: tt.fields.service,
			}
			got, err := m.MaterialFileBatchgetUrl(tt.args.materialType, tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaterialFileBatchgetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaterialFileBatchgetUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpMaterialServiceImpl_MaterialNewsBatchGet(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		offset int
		count  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpMaterialNewsBatchGetResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &WxMpMaterialServiceImpl{
				service: tt.fields.service,
			}
			got, err := m.MaterialNewsBatchGet(tt.args.offset, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaterialNewsBatchGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaterialNewsBatchGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpQrcodeServiceImpl_QrcodeCreateLastTicket(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		actionName    ActionName
		sceneStr      string
		sceneId       int64
		expireSeconds int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpQrCodeTicket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpQrcodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.QrcodeCreateLastTicket(tt.args.actionName, tt.args.sceneStr, tt.args.sceneId, tt.args.expireSeconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("QrcodeCreateLastTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QrcodeCreateLastTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpQrcodeServiceImpl_QrcodeCreateTmpTicket(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		actionName    ActionName
		sceneStr      string
		sceneId       int64
		expireSeconds int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpQrCodeTicket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpQrcodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.QrcodeCreateTmpTicket(tt.args.actionName, tt.args.sceneStr, tt.args.sceneId, tt.args.expireSeconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("QrcodeCreateTmpTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QrcodeCreateTmpTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpQrcodeServiceImpl_getQrCodeTicket(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		actionName    ActionName
		sceneStr      string
		sceneId       int64
		expireSeconds int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpQrCodeTicket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpQrcodeServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.getQrCodeTicket(tt.args.actionName, tt.args.sceneStr, tt.args.sceneId, tt.args.expireSeconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("getQrCodeTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQrCodeTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_CheckSignature(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
			if got := s.CheckSignature(tt.args.timestamp, tt.args.nonce, tt.args.signature); got != tt.want {
				t.Errorf("CheckSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_CreateJsapiSignature(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpServiceImpl_ForceGetJsapiTicket(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpServiceImpl_ForceGetTicket(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpServiceImpl_GetJsapiTicket(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpServiceImpl_GetTicket(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpServiceImpl_GetWxMpConfig(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMpConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
			if got := s.GetWxMpConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMpConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_GetWxMpMaterialService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMpMaterialService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
			if got := s.GetWxMpMaterialService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMpMaterialService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_GetWxMpQrcodeService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMpQrcodeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
			if got := s.GetWxMpQrcodeService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMpQrcodeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_GetWxMpUserService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	tests := []struct {
		name   string
		fields fields
		want   WxMpUserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
			if got := s.GetWxMpUserService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxMpUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpServiceImpl_SetWxMpConfig(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	type args struct {
		config WxMpConfig
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
		})
	}
}

func TestWxMpServiceImpl_SetWxMpMaterialService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	type args struct {
		service WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
		})
	}
}

func TestWxMpServiceImpl_SetWxMpQrcodeService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	type args struct {
		service WxMpQrcodeService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
		})
	}
}

func TestWxMpServiceImpl_SetWxMpUserService(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
	}
	type args struct {
		service WxMpUserService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
			}
		})
	}
}

func TestWxMpServiceImpl_getTicket(t *testing.T) {
	type fields struct {
		WxServiceImpl   common.WxServiceImpl
		config          WxMpConfig
		userService     WxMpUserService
		qrcodeService   WxMpQrcodeService
		materialService WxMpMaterialService
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
			s := &WxMpServiceImpl{
				WxServiceImpl:   tt.fields.WxServiceImpl,
				config:          tt.fields.config,
				userService:     tt.fields.userService,
				qrcodeService:   tt.fields.qrcodeService,
				materialService: tt.fields.materialService,
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

func TestWxMpUserServiceImpl_ChangeOpenid(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		fromAppId string
		openidArr []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WxMpChangeOpenid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.ChangeOpenid(tt.args.fromAppId, tt.args.openidArr...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeOpenid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeOpenid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpUserServiceImpl_GetUserInfo(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		openid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.GetUserInfo(tt.args.openid)
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

func TestWxMpUserServiceImpl_GetUserInfoBy(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		openid string
		lang   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.GetUserInfoBy(tt.args.openid, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfoBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfoBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpUserServiceImpl_GetUserInfoList(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		openidArr []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WxMpUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.GetUserInfoList(tt.args.openidArr...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfoList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpUserServiceImpl_GetUserInfoListBy(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		arr []*WxMpUserQueryParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*WxMpUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.GetUserInfoListBy(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfoListBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfoListBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpUserServiceImpl_GetUserList(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		nextOpenid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WxMpUserList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			got, err := r.GetUserList(tt.args.nextOpenid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxMpUserServiceImpl_UserUpdateRemark(t *testing.T) {
	type fields struct {
		service WxMpService
	}
	type args struct {
		openid string
		remark string
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
			r := &WxMpUserServiceImpl{
				service: tt.fields.service,
			}
			if err := r.UserUpdateRemark(tt.args.openid, tt.args.remark); (err != nil) != tt.wantErr {
				t.Errorf("UserUpdateRemark() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newWxMpConfig(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WxMpConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMpConfig(tt.args.appId, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMpConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMpMaterialService(t *testing.T) {
	type args struct {
		service WxMpService
	}
	tests := []struct {
		name string
		args args
		want WxMpMaterialService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMpMaterialService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMpMaterialService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMpQrcodeService(t *testing.T) {
	type args struct {
		service WxMpService
	}
	tests := []struct {
		name string
		args args
		want *WxMpQrcodeServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMpQrcodeService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMpQrcodeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMpService(t *testing.T) {
	type args struct {
		config WxMpConfig
	}
	tests := []struct {
		name string
		args args
		want *WxMpServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMpService(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMpService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWxMpUserService(t *testing.T) {
	type args struct {
		service WxMpService
	}
	tests := []struct {
		name string
		args args
		want *WxMpUserServiceImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWxMpUserService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWxMpUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}
