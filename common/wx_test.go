package common

import (
	"reflect"
	"testing"
)

func TestCreateJsapiSignature(t *testing.T) {
	type args struct {
		url    string
		appId  string
		ticket string
	}
	tests := []struct {
		name    string
		args    args
		want    *WxJsapiSignature
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

func TestErrMsg_Error(t *testing.T) {
	type fields struct {
		Err Err
		Msg string
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
			e := &ErrMsg{
				Err: tt.fields.Err,
				Msg: tt.fields.Msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErr_Error(t *testing.T) {
	type fields struct {
		ErrCode uint64
		ErrMsg  string
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
			e := &Err{
				ErrCode: tt.fields.ErrCode,
				ErrMsg:  tt.fields.ErrMsg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorOf(t *testing.T) {
	type args struct {
		msg    string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
		want *ErrMsg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorOf(tt.args.msg, tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServiceFor(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceFor(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewXmlService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXmlService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXmlService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewXmlServiceFor(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXmlServiceFor(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXmlServiceFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_Get(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url  string
		args []interface{}
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
			s := &ServiceImpl{
				client: tt.fields.client,
			}
			got, err := s.Get(tt.args.url, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_GetFor(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		v    interface{}
		url  string
		args []interface{}
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
			s := &ServiceImpl{
				client: tt.fields.client,
			}
			if err := s.GetFor(tt.args.v, tt.args.url, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("GetFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceImpl_Post(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &ServiceImpl{
				client: tt.fields.client,
			}
			got, err := s.Post(tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_PostFor(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		v           interface{}
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &ServiceImpl{
				client: tt.fields.client,
			}
			if err := s.PostFor(tt.args.v, tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("PostFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxServiceImpl_ExpireAccessToken(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
		})
	}
}

func TestWxServiceImpl_ForceGetAccessToken(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		forceRefresh bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccessToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			got, err := s.ForceGetAccessToken(tt.args.forceRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForceGetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceGetAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_Get(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		url  string
		args []interface{}
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			got, err := s.Get(tt.args.url, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_GetAccessToken(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	tests := []struct {
		name    string
		fields  fields
		want    *AccessToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			got, err := s.GetAccessToken()
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

func TestWxServiceImpl_GetFor(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		v    interface{}
		url  string
		args []interface{}
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			if err := s.GetFor(tt.args.v, tt.args.url, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("GetFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxServiceImpl_GetWxConfig(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	tests := []struct {
		name   string
		fields fields
		want   WxConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			if got := s.GetWxConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWxConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_IsAccessTokenExpired(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			if got := s.IsAccessTokenExpired(); got != tt.want {
				t.Errorf("IsAccessTokenExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_Post(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			got, err := s.Post(tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_PostFor(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		v           interface{}
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			if err := s.PostFor(tt.args.v, tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("PostFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWxServiceImpl_SetHttpService(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		service Service
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
		})
	}
}

func TestWxServiceImpl_SetWxConfig(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		config WxConfig
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
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
		})
	}
}

func TestWxServiceImpl_attachAccessToken(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	type args struct {
		url  string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			if got := s.attachAccessToken(tt.args.url, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxServiceImpl_getAccessToken(t *testing.T) {
	type fields struct {
		config WxConfig
		http   Service
	}
	tests := []struct {
		name    string
		fields  fields
		want    *AccessToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WxServiceImpl{
				config: tt.fields.config,
				http:   tt.fields.http,
			}
			got, err := s.getAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxTicketImpl_ExpireTicket(t *testing.T) {
	type fields struct {
		JsapiTicket  *Ticket
		SdkTicket    *Ticket
		WxCardTicket *Ticket
	}
	type args struct {
		ticketType TicketType
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
			c := &WxTicketImpl{
				JsapiTicket:  tt.fields.JsapiTicket,
				SdkTicket:    tt.fields.SdkTicket,
				WxCardTicket: tt.fields.WxCardTicket,
			}
		})
	}
}

func TestWxTicketImpl_GetTicket(t *testing.T) {
	type fields struct {
		JsapiTicket  *Ticket
		SdkTicket    *Ticket
		WxCardTicket *Ticket
	}
	type args struct {
		ticketType TicketType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Ticket
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WxTicketImpl{
				JsapiTicket:  tt.fields.JsapiTicket,
				SdkTicket:    tt.fields.SdkTicket,
				WxCardTicket: tt.fields.WxCardTicket,
			}
			if got := c.GetTicket(tt.args.ticketType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxTicketImpl_IsTicketExpired(t *testing.T) {
	type fields struct {
		JsapiTicket  *Ticket
		SdkTicket    *Ticket
		WxCardTicket *Ticket
	}
	type args struct {
		ticketType TicketType
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
			c := &WxTicketImpl{
				JsapiTicket:  tt.fields.JsapiTicket,
				SdkTicket:    tt.fields.SdkTicket,
				WxCardTicket: tt.fields.WxCardTicket,
			}
			if got := c.IsTicketExpired(tt.args.ticketType); got != tt.want {
				t.Errorf("IsTicketExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxTicketImpl_UpdateTicket(t *testing.T) {
	type fields struct {
		JsapiTicket  *Ticket
		SdkTicket    *Ticket
		WxCardTicket *Ticket
	}
	type args struct {
		ticketType TicketType
		ticket     *Ticket
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
			c := &WxTicketImpl{
				JsapiTicket:  tt.fields.JsapiTicket,
				SdkTicket:    tt.fields.SdkTicket,
				WxCardTicket: tt.fields.WxCardTicket,
			}
		})
	}
}

func TestXmlServiceImpl_Get(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url  string
		args []interface{}
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
			s := &XmlServiceImpl{
				client: tt.fields.client,
			}
			got, err := s.Get(tt.args.url, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlServiceImpl_GetFor(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		v    interface{}
		url  string
		args []interface{}
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
			s := &XmlServiceImpl{
				client: tt.fields.client,
			}
			if err := s.GetFor(tt.args.v, tt.args.url, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("GetFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXmlServiceImpl_Post(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &XmlServiceImpl{
				client: tt.fields.client,
			}
			got, err := s.Post(tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlServiceImpl_PostFor(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		v           interface{}
		url         string
		contentType string
		data        interface{}
		args        []interface{}
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
			s := &XmlServiceImpl{
				client: tt.fields.client,
			}
			if err := s.PostFor(tt.args.v, tt.args.url, tt.args.contentType, tt.args.data, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("PostFor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
