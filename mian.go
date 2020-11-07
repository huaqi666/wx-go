package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {

	c := GetConfig()

	maTest(c.Ma)

	mpTest(c.Mp)

	payTest(c.Pay)
}

type Config struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
	MchId  string `json:"mch_id"`
	MchKey string `json:"mch_key"`
	Openid string `json:"openid"`
}

type WxConfig struct {
	Ma  Config `json:"ma"`
	Mp  Config `json:"mp"`
	Pay Config `json:"pay"`
}

func GetConfig() WxConfig {
	var c WxConfig
	f, err := os.Open("./config.json")
	if err != nil {
		f, err := os.Create("./config.json")
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
