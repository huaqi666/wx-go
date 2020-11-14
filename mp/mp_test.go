package mp

import (
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
