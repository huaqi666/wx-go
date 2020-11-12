package mp

import (
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
)

func CheckSignature(token, timestamp, nonce, signature string) bool {
	return util.CheckSignature(token, timestamp, nonce, signature)
}

func CreateJsapiSignature(url, appId, ticket string) (*common.WxJsapiSignature, error) {
	return common.CreateJsapiSignature(url, appId, ticket)
}
