package common

import (
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

func CheckSignature(token, timestamp, nonce, signature string) bool {
	return util.CheckSignature(token, timestamp, nonce, signature)
}

func CreateJsapiSignature(url, appId, ticket string) (*WxJsapiSignature, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	randomStr := util.RandSeq(16)
	arr := []string{"jsapi_ticket=" + ticket, "noncestr=" + randomStr, "timestamp=" + timestamp, "url=" + url}
	signature, err := util.GenWithAmple(arr)
	if err != nil {
		return nil, err
	}
	return &WxJsapiSignature{
		AppId:     appId,
		Timestamp: timestamp,
		NonceStr:  randomStr,
		Url:       url,
		Signature: signature,
	}, nil
}
