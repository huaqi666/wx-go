package common

// common
const (
	ApiDefaultHostUrl  = "https://api.weixin.qq.com"
	MpDefaultHostUrl   = "https://mp.weixin.qq.com"
	OpenDefaultHostUrl = "https://open.weixin.qq.com"

	AccessTokenUrl = ApiDefaultHostUrl + "/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	ApiSuffix = "?access_token=%s"

	Success = "SUCCESS"
	Fail    = "FAIL"
)

// ma
const (
	MaSessionInfoUrl = ApiDefaultHostUrl + "/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	MaQrcodeUrl          = ApiDefaultHostUrl + "/cgi-bin/wxaapp/createwxaqrcode" + ApiSuffix
	MaQrWxaCodeUrl       = ApiDefaultHostUrl + "/wxa/getwxacode" + ApiSuffix
	MaQrCodeUnlimitedUrl = ApiDefaultHostUrl + "/wxa/getwxacodeunlimit" + ApiSuffix
	MaGetPaidUnionIdUrl  = ApiDefaultHostUrl + "/wxa/getpaidunionid" + ApiSuffix

	MaSetUserStorage = ApiDefaultHostUrl + "/wxa/set_user_storage?appid=%s&signature=%s&openid=%s&sig_method=%s"

	MaGetPubTemplateTitleListUrl    = ApiDefaultHostUrl + "/wxaapi/newtmpl/getpubtemplatetitles" + ApiSuffix
	MaGetPubTemplateKeyWordsByIdUrl = ApiDefaultHostUrl + "/wxaapi/newtmpl/getpubtemplatekeywords" + ApiSuffix
	MaTemplateListUrl               = ApiDefaultHostUrl + "/wxaapi/newtmpl/gettemplate" + ApiSuffix
	MaGetCategoryUrl                = ApiDefaultHostUrl + "/wxaapi/newtmpl/getcategory" + ApiSuffix
	MaTemplateDelUrl                = ApiDefaultHostUrl + "/wxaapi/newtmpl/deltemplate" + ApiSuffix
	MaTemplateAddUrl                = ApiDefaultHostUrl + "/wxaapi/newtmpl/addtemplate" + ApiSuffix

	MaSubscribeMsgSendUrl = ApiDefaultHostUrl + "/cgi-bin/message/subscribe/send" + ApiSuffix
	MaUniformMsgSendUrl   = ApiDefaultHostUrl + "/cgi-bin/message/wxopen/template/uniform_send" + ApiSuffix
	MaKefuMessageSendUrl  = ApiDefaultHostUrl + "/cgi-bin/message/custom/send" + ApiSuffix
	MaUpdatableMsgSendUrl = ApiDefaultHostUrl + "/cgi-bin/message/wxopen/updatablemsg/send" + ApiSuffix
	MaActivityIdCreateUrl = ApiDefaultHostUrl + "/cgi-bin/message/wxopen/activityid/create" + ApiSuffix

	MaCreateRoom  = ApiDefaultHostUrl + "/wxaapi/broadcast/room/create" + ApiSuffix
	MaEditRoom    = ApiDefaultHostUrl + "/wxaapi/broadcast/room/editroom" + ApiSuffix
	MaGetLiveInfo = ApiDefaultHostUrl + "/wxa/business/getliveinfo" + ApiSuffix
	MaAddGoods    = ApiDefaultHostUrl + "/wxaapi/broadcast/room/addgoods" + ApiSuffix
)

// mp
const (
	MpGetTicketUrl = ApiDefaultHostUrl + "/cgi-bin/ticket/getticket" + ApiSuffix + "&type=%s"

	MpQrcodeUrl = ApiDefaultHostUrl + "/cgi-bin/qrcode/create" + ApiSuffix

	MpUserUpdateRemarkUrl = ApiDefaultHostUrl + "/cgi-bin/user/info/updateremark" + ApiSuffix
	MpUserInfoUrl         = ApiDefaultHostUrl + "/cgi-bin/user/info" + ApiSuffix + "&openid=%s&lang=%s"
	MpUserGetUrl          = ApiDefaultHostUrl + "/cgi-bin/user/get" + ApiSuffix + "&next_openid="
	MpUserInfoBatchGetUrl = ApiDefaultHostUrl + "/cgi-bin/user/info/batchget" + ApiSuffix
	MpUserChangeOpenidUrl = ApiDefaultHostUrl + "/cgi-bin/changeopenid"

	MpMaterialBatchgetUrl = ApiDefaultHostUrl + "/cgi-bin/material/batchget_material" + ApiSuffix
)

// pay
const (
	PayDefaultPayBaseUrl = "https://api.mch.weixin.qq.com"
	PayUnifiedOrder      = "/pay/unifiedorder"
	PayCloseOrder        = "/pay/closeorder"
	PayQueryOrder        = "/pay/orderquery"

	PayGetSandboxSignKey = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

	PayRefundUrl          = "/secapi/pay/refund"
	PayRefundUrlV2        = "/secapi/pay/refundv2"
	PayRefundSandboxUrl   = "/pay/refund"
	PayRefundSandboxUrlV2 = "/pay/refundv2"
	PayQueryRefundUrl     = "/pay/refundquery"
	PayQueryRefundUrlV2   = "/pay/refundqueryv2"

	EntPayUrl          = "/mmpaymkttransfers/promotion/transfers"
	EntPayQueryUrl     = "/mmpaymkttransfers/gettransferinfo"
	EntPayBankUrl      = "/mmpaysptrans/pay_bank"
	EntPayQueryBankUrl = "/mmpaysptrans/query_bank"

	EntSendEnterpriseRedPackUrl  = "/mmpaymkttransfers/sendworkwxredpack"
	EntQueryEnterpriseRedPackUrl = "/mmpaymkttransfers/queryworkwxredpack"
)

// http header
const (
	PostXmlContentType  = "application/xml; charset=utf-8"
	PostJsonContentType = "application/json; charset=utf-8"
)
