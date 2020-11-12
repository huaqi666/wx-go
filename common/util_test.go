package common

import (
	"testing"
)

func TestCreateJsapiSignature(t *testing.T) {
	t.Log(CreateJsapiSignature("www.xxx.com", "", ""))
}
