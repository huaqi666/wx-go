package main

import (
	"encoding/json"
	"fmt"
	"github.com/cliod/wx-go/ma"
	"github.com/cliod/wx-go/mp"
	"github.com/cliod/wx-go/pay"
	"math/rand"
	"strconv"
	"time"
)

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
		r, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(r))
		}
	} else {
		fmt.Println(err.Error())
	}
}

func mpTest(c Config) {
	service := mp.NewWxMpService(c.AppId, c.Secret)

	da, err := service.CreateJsapiSignature("https://www.baidu.com")
	if err == nil {
		r, err := json.Marshal(da)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(r))
		}
	} else {
		fmt.Println(err.Error())
	}

	qc := service.GetWxMpQrcodeService()

	bytes, err := qc.QrcodeCreateTmpTicket(mp.QrScene, "/pages/index", 0, 30)
	if err == nil {
		r, err := json.Marshal(bytes)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(r))
		}
	} else {
		fmt.Println(err.Error())
	}

	uc := service.GetWxMpUserService()

	res, err := uc.GetUserInfo(c.Openid)
	if err == nil {
		r, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(r))
		}
	} else {
		fmt.Println(err.Error())
	}
}

func payTest(c Config) {

	conf := pay.NewBaseV2Config(c.AppId, c.MchId, c.MchKey, "https://www.baidu.com/notify", "")
	//conf.UseSandboxEnv = true
	service := pay.NewWxPayServiceFor(conf)

	s := strconv.Itoa(time.Now().Nanosecond()) + strconv.Itoa(rand.Intn(999999))

	fmt.Println(s)

	res, err := service.UnifyPay(&pay.WxPayUnifiedOrderRequest{
		TotalFee:   100,
		Openid:     c.Openid,
		OutTradeNo: s,
		Body:       "测试数据",
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(res))
	}

	d, err := service.CloseOrderBy(s)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		res, err = json.Marshal(d)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(res))
		}
	}
}
