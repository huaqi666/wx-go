package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wx-go/ma"
	"wx-go/mp"
)

func main() {

	c := GetConfig()

	maTest(c.Ma)

	mpTest(c.Mp)

	payTest(c.Pay)
}

func maTest(c Config) {
	service := ma.NewWxMaService(c.AppId, c.Secret)

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

	uc := service.GetWxMaUserService()

	res, err := uc.GetSessionInfo("<js_code>")
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err.Error())
	}
}

func mpTest(c Config) {
	service := mp.NewWxMpService(c.AppId, c.Secret)

	da, err := service.CreateJsapiSignature("https://www.baidu.com")
	if err == nil {
		fmt.Println(da)
	} else {
		fmt.Println(err.Error())
	}

	qc := service.GetWxMpQrcodeService()

	bytes, err := qc.QrcodeCreateTmpTicket(mp.QrScene, "/pages/index", 0, 30)
	if err == nil {
		fmt.Println(bytes)
	} else {
		fmt.Println(err.Error())
	}

	uc := service.GetWxMpUserService()

	res, err := uc.GetUserInfo("<open_id>")
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err.Error())
	}
}

func payTest(c Config) {

}

type Config struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
}

type WxConfig struct {
	Ma  Config `json:"ma"`
	Mp  Config `json:"mp"`
	Pay Config `json:"pay"`
}

func GetConfig() WxConfig {
	f, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var c WxConfig
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}
