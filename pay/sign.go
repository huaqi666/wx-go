package pay

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

var Ignores = []string{"sign", "key", "xmlString", "xmlDoc", "couponList"}

func buildSign(params map[string]interface{}, sk string, ignoreParams ...string) string {
	var arr []string
	for k, _ := range params {
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
		case int:
			vv = strconv.Itoa(value.(int))
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

func buildSignFor(request interface{}, sk string, ignoreParams ...string) string {
	b, err := json.Marshal(request)
	if err != nil {
		return ""
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return ""
	}
	return buildSign(data, sk, ignoreParams...)
}
