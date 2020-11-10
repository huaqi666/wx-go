package pay

import (
	"github.com/cliod/wx-go/common"
	"github.com/cliod/wx-go/common/util"
	"strconv"
	"time"
)

type WxEntPayService interface {
	/* 企业付款API.
	   企业付款业务是基于微信支付商户平台的资金管理能力，为了协助商户方便地实现企业向个人付款，针对部分有开发能力的商户，提供通过API完成企业付款的功能。
	   比如目前的保险行业向客户退保、给付、理赔。
	   企业付款将使用商户的可用余额，需确保可用余额充足。查看可用余额、充值、提现请登录商户平台“资金管理”https://pay.weixin.qq.com/进行操作。
	   注意：与商户微信支付收款资金并非同一账户，需要单独充值。
	   文档详见: https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2 */
	EntPay(*WxEntPayRequest) (*WxEntPayResult, error)
	/* 查询企业付款API.
	   用于商户的企业付款操作进行结果查询，返回付款操作详细结果。
	   文档详见:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3 */
	QueryEntPayBy(partnerTradeNo string) (*EntPayQueryResult, error)
	/* 查询企业付款API.
	   用于商户的企业付款操作进行结果查询，返回付款操作详细结果。
	   文档详见:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3 */
	QueryEntPay(request *EntPayQueryRequest) (*EntPayQueryResult, error)
	/* 获取RSA加密公钥API.
	   RSA算法使用说明（非对称加密算法，算法采用RSA/ECB/OAEPPadding模式）
	   1、 调用获取RSA公钥API获取RSA公钥，落地成本地文件，假设为public.pem
	   2、 确定public.pem文件的存放路径，同时修改代码中文件的输入路径，加载RSA公钥
	   3、 用标准的RSA加密库对敏感信息进行加密，选择RSA_PKCS1_OAEP_PADDING填充模式
	   （eg：Java的填充方式要选 " RSA/ECB/OAEPWITHSHA-1ANDMGF1PADDING"）
	   4、 得到进行rsa加密并转base64之后的密文
	   5、 将密文传给微信侧相应字段，如付款接口（enc_bank_no/enc_true_name）

	   接口默认输出PKCS#1格式的公钥，商户需根据自己开发的语言选择公钥格式
	   文档详见:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4 */
	GetPublicKey() (string, error)
	/* 企业付款到银行卡.
	   用于企业向微信用户银行卡付款
	   目前支持接口API的方式向指定微信用户的银行卡付款。
	   文档详见：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2 */
	PayBank(request *EntPayBankRequest) (*EntPayBankResult, error)
	/* 企业付款到银行卡查询.
	   用于对商户企业付款到银行卡操作进行结果查询，返回付款操作详细结果。
	   文档详见：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3 */
	QueryPayBankBy(partnerTradeNo string) (*EntPayBankQueryResult, error)
	/* 企业付款到银行卡查询.
	   用于对商户企业付款到银行卡操作进行结果查询，返回付款操作详细结果。
	   文档详见：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3 */
	QueryPayBank(request *EntPayBankQueryRequest) (*EntPayBankQueryResult, error)
	/* 企业发送微信红包给个人用户
	   文档地址：https://work.weixin.qq.com/api/doc
	*/
	SendEnterpriseRedPack(request *EntPayRedPackRequest) (*EntPayRedPackResult, error)
	/* 企业发送微信红包查询
	   文档地址：https://work.weixin.qq.com/api/doc */
	QueryEnterpriseRedPack(request *EntPayRedPackQueryRequest) (*EntPayRedPackQueryResult, error)
}

type WxEntPayServiceImpl struct {
	service WxPayService
}

func newWxEntPayService(service WxPayService) *WxEntPayServiceImpl {
	return &WxEntPayServiceImpl{
		service: service,
	}
}

func (w *WxEntPayServiceImpl) EntPay(request *WxEntPayRequest) (*WxEntPayResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayUrl

	var res WxEntPayResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEntPayBy(partnerTradeNo string) (*EntPayQueryResult, error) {
	return w.QueryEntPay(&EntPayQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryEntPay(request *EntPayQueryRequest) (*EntPayQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayQueryUrl

	var res EntPayQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) GetPublicKey() (string, error) {
	url := "https://fraud.mch.weixin.qq.com/risk/getpublickey"

	var res GetPublicKeyResult
	err := w.service.Do(url, &WxPayDefaultRequest{}, &res, true)
	return res.PubKey, err
}

func (w *WxEntPayServiceImpl) PayBank(request *EntPayBankRequest) (*EntPayBankResult, error) {
	var b []byte
	b, err := w.encryptRSA([]byte(request.EncBankNo))
	if err != nil {
		return nil, err
	}
	request.EncBankNo = string(b)
	b, err = w.encryptRSA([]byte(request.EncTrueName))
	if err != nil {
		return nil, err
	}
	request.EncTrueName = string(b)

	url := w.service.GetPayBaseUr() + common.EntPayBankUrl

	var res EntPayBankResult
	err = w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryPayBankBy(partnerTradeNo string) (*EntPayBankQueryResult, error) {
	return w.QueryPayBank(&EntPayBankQueryRequest{
		PartnerTradeNo: partnerTradeNo,
	})
}

func (w *WxEntPayServiceImpl) QueryPayBank(request *EntPayBankQueryRequest) (*EntPayBankQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntPayQueryBankUrl

	var res EntPayBankQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) SendEnterpriseRedPack(request *EntPayRedPackRequest) (*EntPayRedPackResult, error) {
	//企业微信签名,需要在请求签名之前
	request.NonceStr = strconv.Itoa(time.Now().Nanosecond())
	request.WorkWxSign = SignEnt(request.ActName, request.MchBillNo, request.MchId, request.NonceStr, request.ReOpenid, request.WxAppId,
		w.service.GetWxPayConfig().EntPayKey, request.TotalAmount, MD5)

	url := w.service.GetPayBaseUr() + common.EntSendEnterpriseRedPackUrl
	request.Sign = w.service.Sign(request, request.SignType)

	var res EntPayRedPackResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) QueryEnterpriseRedPack(request *EntPayRedPackQueryRequest) (*EntPayRedPackQueryResult, error) {
	url := w.service.GetPayBaseUr() + common.EntQueryEnterpriseRedPackUrl

	var res EntPayRedPackQueryResult
	err := w.service.Do(url, request, &res, true)
	return &res, err
}

func (w *WxEntPayServiceImpl) encryptRSA(data []byte) ([]byte, error) {
	key, err := w.GetPublicKey()
	if err != nil {
		return nil, err
	}
	return util.RsaEncrypt(data, []byte(key))
}
