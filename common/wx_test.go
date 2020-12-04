package common

import (
	"testing"
)

func TestCreateJsapiSignature(t *testing.T) {
	t.Log(CreateJsapiSignature("www.xxx.com", "", ""))
}

func TestServiceImpl_Get(t *testing.T) {
	t.Log(NewService().Get("http://127.0.0.1:80"))
}

func TestServiceImpl_GetFor(t *testing.T) {
	var v interface{}
	t.Log(NewService().GetFor(&v, "http://127.0.0.1:80/ping"))
	t.Log(v)
}

func TestServiceImpl_Post(t *testing.T) {
	t.Log(NewService().Post("http://127.0.0.1:80/post", PostJsonContentType, nil))
}

func TestServiceImpl_PostFor(t *testing.T) {
	var v interface{}
	t.Log(NewService().PostFor(&v, "http://127.0.0.1:80/post", PostJsonContentType, nil))
	t.Log(t)
}

func TestXmlServiceImpl_Get(t *testing.T) {
	t.Log(NewService().Get("http://127.0.0.1:80"))
}

func TestXmlServiceImpl_GetFor(t *testing.T) {
	var v interface{}
	t.Log(NewService().GetFor(&v, "http://127.0.0.1:80/ping"))
	t.Log(v)
}

func TestXmlServiceImpl_Post(t *testing.T) {
	t.Log(NewService().Post("http://127.0.0.1:80/post", PostJsonContentType, nil))
}

func TestXmlServiceImpl_PostFor(t *testing.T) {
	var v interface{}
	t.Log(NewService().PostFor(&v, "http://127.0.0.1:80/post", PostJsonContentType, nil))
	t.Log(t)
}
