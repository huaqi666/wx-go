package wx

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 测试使用
type Config struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
	MchId  string `json:"mch_id"`
	MchKey string `json:"mch_key"`
	Openid string `json:"openid"`
}

// 测试使用
type TestConfig struct {
	Ma  Config `json:"ma"`
	Mp  Config `json:"mp"`
	Pay Config `json:"pay"`
}

// 测试使用
func GetConfig(filename string) TestConfig {
	var c TestConfig
	f, err := os.Open(filename)
	if err != nil {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		_, _ = f.Write(b)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}
