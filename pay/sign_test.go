package pay

import (
	"testing"
)

func TestSignFor(t *testing.T) {

	t.Log(Sign(BaseWxPayRequest{
		AppId: "123",
		MchId: "123",
	}, MD5, ""))
}

func TestBuildSignStr(t *testing.T) {
	t.Log(buildSignStr(BaseWxPayRequest{
		AppId: "123",
		MchId: "123",
	}, "123"))
}
