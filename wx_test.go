package wx

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := GetConfig("./config.json")
	fmt.Println(conf.Pay.MchId)
}
