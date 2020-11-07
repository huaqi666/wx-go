package main

import (
	"testing"
)

func TestAll(t *testing.T) {

	c := WxConfig{
		Ma: Config{
			AppId:  "",
			Secret: "",
			Openid: "",
		},
		Mp: Config{
			AppId:  "",
			Secret: "",
			Openid: "",
		},
		Pay: Config{
			AppId:  "",
			MchId:  "",
			MchKey: "",
			Openid: "",
		},
	}

	maTest(c.Ma)

	mpTest(c.Mp)

	payTest(c.Pay)
}
