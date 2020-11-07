package mp

import "github.com/cliod/wx-go/common"

type WxMpUserService interface {
	// 设置用户备注名
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140838&token=&lang=zh_CN
	UserUpdateRemark(openid, remark string) error

	// 获取用户基本信息（语言为默认的zh_CN 简体）
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839&token=&lang=zh_CN
	GetUserInfo(openid string) (*WxMpUser, error)
	// 获取用户基本信息指定语言
	GetUserInfoBy(openid, lang string) (*WxMpUser, error)

	// 获取用户基本信息列表
	// 开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条。
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839&token=&lang=zh_CN
	GetUserInfoList(openidArr ...string) ([]*WxMpUser, error)
	// 获取用户基本信息列表指定语言
	GetUserInfoListBy([]*WxMpUserQueryParam) ([]*WxMpUser, error)

	// 获取用户列表
	// 公众号可通过本接口来获取帐号的关注者列表，
	// 关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。
	// 一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求。
	// 详情请见: http://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140840&token=&lang=zh_CN
	GetUserList(nextOpenid string) (*WxMpUserList, error)
	// 微信公众号主体变更迁移用户 openid
	// 详情请见: http://kf.qq.com/faq/170221aUnmmU170221eUZJNf.html, http://kf.qq.com/faq/1901177NrqMr190117nqYJze.html
	// fromAppId: 原公众号的
	// 一次最多100个
	ChangeOpenid(fromAppId string, openidArr ...string) ([]*WxMpChangeOpenid, error)
}

// 微信用户信息
type WxMpUser struct {
	Subscribe bool   `json:"subscribe"`
	OpenId    string `json:"open_id"`
	Nickname  string `json:"nickname"`
	// 性别描述信息：男、女、未知等.
	SexDesc string `json:"sex_desc"`
	// 性别表示：1，2等数字.
	Sex           string `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	HeadImgUrl    string `json:"head_img_url"`
	SubscribeTime string `json:"subscribe_time"`
	// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&announce_id=11513156443eZYea&version=&lang=zh_CN
	// 只有在将公众号绑定到微信开放平台帐号后，才会出现该字段。
	// 另外，在用户未关注公众号时，将不返回用户unionID信息。
	// 已关注的用户，开发者可使用“获取用户基本信息接口”获取unionID；
	// 未关注用户，开发者可使用“微信授权登录接口”并将scope参数设置为snsapi_userinfo，获取用户unionID
	UnionId string `json:"union_id"`
	Remark  string `json:"remark"`
	GroupId string `json:"group_id"`
	TagIds  string `json:"tag_ids"`

	// 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）.
	Privileges string `json:"privileges"`
	// subscribe_scene 返回用户关注的渠道来源.
	// ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	SubscribeScene string `json:"subscribe_scene"`
	// 二维码扫码场景（开发者自定义）.
	QrScene string `json:"qr_scene"`
	// 二维码扫码场景描述（开发者自定义）.
	QrSceneStr string `json:"qr_scene_str"`
}

type WxMpUserQueryParam struct {
	Openid string `json:"openid"`
	Lang   string `json:"lang"`
}

func GetWxMpUserQueryParam(openidArr ...string) []*WxMpUserQueryParam {
	var arr []*WxMpUserQueryParam
	for _, openid := range openidArr {
		arr = append(arr, &WxMpUserQueryParam{
			Openid: openid,
		})
	}
	return arr
}

type WxMpUserList struct {
	Total      uint64   `json:"total"`
	Count      int      `json:"count"`
	NextOpenid string   `json:"next_openid"`
	OpenidArr  []string `json:"openids"`
}

type WxMpChangeOpenid struct {
	OriOpenid string `json:"ori_openid"`
	NewOpenid string `json:"new_openid"`
	ErrMsg    string `json:"err_msg"`
}

type WxMpUserServiceImpl struct {
	service WxMpService
}

func newWxMpUserService(service WxMpService) *WxMpUserServiceImpl {
	return &WxMpUserServiceImpl{
		service: service,
	}
}

func (r *WxMpUserServiceImpl) UserUpdateRemark(openid, remark string) error {

	data := map[string]interface{}{
		"openid": openid,
		"remark": remark,
	}
	_, err := r.service.Post(common.MpUserUpdateRemarkUrl, "", data)
	return err
}

func (r *WxMpUserServiceImpl) GetUserInfo(openid string) (*WxMpUser, error) {

	return r.GetUserInfoBy(openid, "")
}

func (r *WxMpUserServiceImpl) GetUserInfoBy(openid, lang string) (*WxMpUser, error) {

	if lang == "" {
		lang = "zh_CN"
	}

	var data WxMpUser
	err := r.service.GetFor(&data, common.MpUserInfoUrl, openid, lang)
	return &data, err
}

func (r *WxMpUserServiceImpl) GetUserInfoList(openidArr ...string) ([]*WxMpUser, error) {

	return r.GetUserInfoListBy(GetWxMpUserQueryParam(openidArr...))
}

func (r *WxMpUserServiceImpl) GetUserInfoListBy(arr []*WxMpUserQueryParam) ([]*WxMpUser, error) {

	data := map[string][]*WxMpUserQueryParam{
		"user_list": arr,
	}

	var res []*WxMpUser
	err := r.service.PostFor(&res, common.MpUserInfoBatchGetUrl, common.PostJsonContentType, data)
	return res, err
}

func (r *WxMpUserServiceImpl) GetUserList(nextOpenid string) (*WxMpUserList, error) {

	var data WxMpUserList
	err := r.service.GetFor(&data, common.MpUserGetUrl, nextOpenid)
	return &data, err
}

func (r *WxMpUserServiceImpl) ChangeOpenid(fromAppId string, openidArr ...string) ([]*WxMpChangeOpenid, error) {

	data := map[string]interface{}{
		"from_appid":  fromAppId,
		"openid_list": openidArr,
	}

	var res []*WxMpChangeOpenid
	err := r.service.PostFor(&res, common.MpUserChangeOpenidUrl, common.PostJsonContentType, data)
	return res, err
}
