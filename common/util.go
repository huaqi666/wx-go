package common

import (
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

func CreateJsapiSignature(url, appId, ticket string) (*WxJsapiSignature, error) {
	timestamp := strconv.Itoa(time.Now().Second())
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
