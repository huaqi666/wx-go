package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wx-go/ma"
)

func main() {

	c := GetConfig()
	appId := c.AppId
	secret := c.Secret

	service := ma.NewService(appId, secret)

	qc := service.GetWxMaQrcodeService()
	bytes, err := qc.CreateQrcode("/pages/index")
	if err == nil {
		err = qc.(*ma.WxMaQrCodeServiceImpl).BytesToFile("tmp.jpg", bytes)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

type Config struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
}

func GetConfig() Config {
	f, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}
