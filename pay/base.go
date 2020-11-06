package pay

import (
	"encoding/xml"
	"fmt"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strings"
)

type (
	TradeType string
	SignType  string
)

const (
	JSAPI  TradeType = "JSAPI"
	NATIVE TradeType = "NATIVE"
	APP    TradeType = "APP"
	H5     TradeType = "MWEB"
	MICRO  TradeType = "MICROPAY"
)

const (
	HmacSha256 SignType = "HMAC-SHA256"
	MD5        SignType = "MD5"
)

// 基础请求对象
type BaseWxPayRequest struct {
	XMLName xml.Name `xml:"xml" json:"-"`

	// 公众号或者小程序appId
	AppId string `json:"appid" xml:"appid"`
	// 商户号.
	MchId string `json:"mch_id" xml:"mch_id"`
	// 服务商模式下的子商户公众账号ID.
	SubAppId string `json:"sub_app_id,omitempty" xml:"sub_app_id,omitempty"`
	// 服务商模式下的子商户号.
	SubMchId string `json:"sub_mch_id,omitempty" xml:"sub_mch_id,omitempty"`
	// 随机字符串.不长于32位。推荐随机数生成算法
	NonceStr string `json:"nonce_str" xml:"nonce_str"`
	// 签名.
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 签名类型.
	SignType SignType `json:"sign_type" xml:"sign_type,omitempty"`
	// 企业微信签名
	WorkWxSign string `json:"work_wx_sign,omitempty" xml:"work_wx_sign,omitempty"`
}

// 基础响应对象
type BaseWxPayResult struct {
	common.Err
	ResultCode string `json:"result_code" xml:"result_code"`
	RetMsg     string `json:"retmsg" xml:"retmsg"`

	ReturnCode string `json:"return_code" xml:"return_code"`
	ReturnMsg  string `json:"return_msg" xml:"return_msg"`

	ErrCode    string `json:"err_code" xml:"err_code"`
	ErrCodeDes string `json:"err_code_des" xml:"err_code_des"`

	AppId    string `json:"appid" xml:"appid"`
	MchId    string `json:"mch_id" xml:"mch_id"`
	SubAppId string `json:"sub_app_id" xml:"sub_app_id"`
	SubMchId string `json:"sub_mch_id" xml:"sub_mch_id"`
	NonceStr string `json:"nonce_str" xml:"nonce_str"`
	Sign     string `json:"sign" xml:"sign"`
}

func (r *BaseWxPayRequest) CheckAndSign(c *WxPayConfig) {
	if !r.IsIgnoreAppId() {
		if r.AppId == "" {
			r.AppId = c.AppId
		}
	}
	if r.MchId == "" {
		r.MchId = c.MchId
	}
	if !r.IsIgnoreSubAppId() {
		if r.SubAppId == "" {
			r.SubAppId = c.SubAppId
		}
	}
	if !r.IsIgnoreSubMchId() {
		if r.SubMchId == "" {
			r.SubMchId = c.SubMchId
		}
	}
	if r.SignType == "" {
		if c.SignType == "" {
			r.SignType = MD5
		} else {
			r.SignType = c.SignType
		}
	}
	if r.NeedNonceStr() && r.NonceStr == "" {
		r.NonceStr = util.RandSeq(32)
	}
	r.Sign = SignFor(r, r.SignType, c.MchKey)
}

func (r *BaseWxPayRequest) IsIgnoreAppId() bool {
	return false
}

func (r *BaseWxPayRequest) IsIgnoreSubAppId() bool {
	return false
}

func (r *BaseWxPayRequest) IsIgnoreSubMchId() bool {
	return false
}

func (r *BaseWxPayRequest) NeedNonceStr() bool {
	return true
}

func (r *BaseWxPayResult) CheckResult(service WxPayService, signType SignType, checkSuccess bool) (string, bool) {
	data := r.ToMap()

	if r.Sign != "" && r.Sign != service.Sign(data, signType) {
		msg := fmt.Sprintf("校验结果签名失败，参数：%s", data)
		return msg, false
	}

	if checkSuccess {
		r1, r2 := TrimToUpper(r.ReturnCode), TrimToUpper(r.ResultCode)
		if !(r1 == common.Success || r1 == "" || r2 == common.Success || r2 == "") {
			var errorMsg string
			if r1 != "" {
				errorMsg += "返回代码：" + r1 + "\n"
			}
			if msg := TrimToUpper(r.ReturnMsg); msg != "" {
				errorMsg += "返回信息：" + msg + "\n"
			}
			if r2 != "" {
				errorMsg += "结果代码：" + r2 + "\n"
			}
			if msg := TrimToUpper(r.ErrCode); msg != "" {
				errorMsg += "错误代码：" + msg + "\n"
			}
			if msg := TrimToUpper(r.ErrCodeDes); msg != "" {
				errorMsg += "错误详情：" + msg + "\n"
			}
			if msg := TrimToUpper(r.ErrMsg); msg != "" {
				errorMsg += "错误信息：" + msg + "\n"
			}
			return errorMsg, false
		}
	}

	return "", true
}

func (r *BaseWxPayResult) ToMap() map[string]interface{} {
	return toMap(r)
}

func TrimToUpper(str string) string {
	return strings.ToUpper(strings.Trim(str, ""))
}
