package mp

import (
	"fmt"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

func CheckSignature(token, timestamp, nonce, signature string) bool {
	arr := []string{token, timestamp, nonce}
	gen, err := util.Gen(arr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return gen == signature
}

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
