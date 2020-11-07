package pay

import (
	"encoding/xml"
	"fmt"
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
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

type WxPayRequest interface {
	CheckAndSign(c *WxPayConfig)
	IsIgnoreAppId() bool
	IsIgnoreSubAppId() bool
	IsIgnoreSubMchId() bool
	NeedNonceStr() bool

	GetSignType() SignType
	IgnoredParamsForSign() []string
}

type WxPayResult interface {
	CheckResult(service WxPayService, signType SignType, checkSuccess bool) error
	Compose()
}

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
	r.Sign = Sign(r, r.SignType, c.MchKey, r.IgnoredParamsForSign()...)
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

func (r *BaseWxPayRequest) GetSignType() SignType {
	return r.SignType
}

func (r *BaseWxPayRequest) IgnoredParamsForSign() []string {
	return []string{}
}

// 基础响应对象
type BaseWxPayResult struct {
	XMLName xml.Name `xml:"xml" json:"-"`

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
	// 以时间戳为随机字符串，可以不设置.
	NonceStr string `json:"nonce_str" xml:"nonce_str"`
	Sign     string `json:"sign" xml:"sign"`

	Content []byte `xml:",innerxml"`
}

func (r *BaseWxPayResult) Error() string {
	return fmt.Sprintf("返回代码：%s，返回信息：%s，结果代码：%s，结果信息：%s，错误代码：%s，错误详情：%s，错误信息：%s",
		r.ReturnCode, r.ReturnMsg, r.ResultCode, r.RetMsg, r.ErrCode, r.ErrCodeDes, r.ErrMsg)
}

func (r *BaseWxPayResult) CheckResult(service WxPayService, signType SignType, checkSuccess bool) error {
	data := util.ToMap(r)

	if r.Sign != "" && r.Sign != service.Sign(data, signType) {
		return common.ErrorOf("校验结果签名失败，参数：%s", data)
	}

	if checkSuccess {
		r1, r2 := util.TrimToUpper(r.ReturnCode), util.TrimToUpper(r.ResultCode)
		if !(r1 == common.Success || r1 == "" || r2 == common.Success || r2 == "") {
			return r
		}
	}

	return nil
}

func (r *BaseWxPayResult) Compose() {
}
