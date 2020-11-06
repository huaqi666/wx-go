package pay

import (
	"github.com/cliod/wx-go/common/util"
	"sort"
	"strconv"
	"strings"
)

var Ignores = []string{"sign", "key", "xmlString", "xmlDoc", "couponList"}

func SignFor(params interface{}, st SignType, sk string, ignoreParams ...string) string {
	signStr := buildSignStr(params, sk, ignoreParams...)
	var sign string
	switch st {
	case HmacSha256:
		sign = util.HmacSha256(signStr, sk)
	case MD5:
		sign = util.Md5(signStr)
	}
	return sign
}

func buildSignStr(request interface{}, sk string, ignoreParams ...string) string {

	var params map[string]interface{}
	params, ok := request.(map[string]interface{})
	if !ok {
		params = util.ToMap(request)
	}
	if params == nil {
		return ""
	}

	var arr []string
	for k := range params {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	ign := strings.Join(Ignores, ",") + strings.Join(ignoreParams, ",")

	var sign string

	for _, key := range arr {
		shouldSign := false
		value := params[key]
		var vv string
		switch value.(type) {
		case string:
			vv = value.(string)
		case float64:
			vv = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		case int64:
			vv = strconv.FormatInt(value.(int64), 10)
		case SignType:
			vv = value.(string)
		}
		if key != "" && vv != "" && !strings.Contains(ign, key) {
			shouldSign = true
		}
		if shouldSign {
			sign += key + "=" + vv + "&"
		}
	}
	return sign + "key=" + sk
}
