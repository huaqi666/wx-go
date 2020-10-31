package pay

import (
	"sort"
	"strings"
)

var Ignores = []string{"sign", "key", "xmlString", "xmlDoc", "couponList"}

func buildSign(params map[string]string, sk string, ignoreParams ...string) string {
	var arr []string
	for _, v := range params {
		arr = append(arr, v)
	}
	sort.Strings(arr)
	ign := strings.Join(Ignores, ",") + strings.Join(ignoreParams, ",")

	var sign string

	for _, key := range arr {
		shouldSign := false
		value := params[key]
		if key != "" && value != "" && !strings.Contains(ign, key) {
			shouldSign = true
		}
		if shouldSign {
			sign += key + "=" + value + "&"
		}
	}
	return sign + "key=" + sk
}
