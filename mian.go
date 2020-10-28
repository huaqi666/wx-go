package main

import (
	"fmt"
	"wx-go/ma"
)

func main() {

	appId := ""
	secret := ""

	service := ma.NewService(appId, secret)
	at, err := service.GetAccessToken()
	if err == nil {
		fmt.Println(at.AccessToken)
	} else {
		fmt.Println(err.Error())
	}

	f, err := service.GetQrCodeService().CreateQrcode("/pages/index")
	if err != nil {
		fmt.Println(f)
	}
}
