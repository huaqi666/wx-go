package main

import (
	"fmt"
	"wx-go/ma"
)

func main() {

	appId := ""
	secret := ""

	service := ma.NewService(appId, secret)

	qc := service.GetQrCodeService()
	bytes, err := qc.CreateQrcode("/pages/index")
	if err == nil {
		err = qc.(*ma.QrCodeServiceImpl).BytesToFile("tmp.jpg", bytes)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
