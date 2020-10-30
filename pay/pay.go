package pay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"wx-go/common"
	"wx-go/common/util"
)

type WxPayService interface {
	common.WxService

	PostKey(url, contentType string, data interface{}, certPath, keyPath string, args ...interface{}) ([]byte, error)

	GetPayBaseUr() string
}

type WxPayServiceImpl struct {
	common.WxServiceImpl
}

func (p *WxPayServiceImpl) PostKey(url, contentType string, data interface{}, certPath, keyPath string, args ...interface{}) ([]byte, error) {
	uri := fmt.Sprintf(url, args...)
	body, err := json.Marshal(data)

	cli, err := util.NewTLSClient(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	res, err := cli.Post(uri, contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (p *WxPayServiceImpl) GetPayBaseUr() string {
	return ""
}
