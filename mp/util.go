package mp

import (
	"fmt"
	"github.com/cliod/wx-go/common/util"
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
