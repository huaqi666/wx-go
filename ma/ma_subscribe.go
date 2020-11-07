package ma

import (
	"github.com/cliod/wx-go/common"
	"strconv"
	"strings"
)

type WxMaSubscribeService interface {
	// 携带access_token执行
	Do(url string, res error) error
	/* 获取帐号所属类目下的公共模板标题
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html */
	GetPubTemplateTitleList(ids []string, start, limit int) (*WxMaPubTemplateTitleListResult, error)
	/* 获取模板库某个模板标题下关键词库
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html */
	GetPubTemplateKeyWordsById(id string) (*WxMaPubTemplateKeywordListResult, error)
	/* 获取当前帐号下的个人模板列表
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html */
	GetTemplateList() (*WxMaTemplateListResult, error)
	/* 组合模板并添加至帐号下的个人模板库
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html */
	AddTemplate(id, sceneDesc string, keywordIds []string) (*WxMaAddTemplateResult, error)
	/* 删除帐号下的某个模板
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html */
	DelTemplate(templateId string) error
	/* 获取小程序账号的类目
	   详情请见: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html */
	GetCategory() (*WxMaCategoryListResult, error)
}

type WxMaSubscribeServiceImpl struct {
	service WxMaService
}

func newWxMaSubscribeService(service WxMaService) *WxMaSubscribeServiceImpl {
	return &WxMaSubscribeServiceImpl{
		service: service,
	}
}

func (i *WxMaSubscribeServiceImpl) GetPubTemplateTitleList(ids []string, start, limit int) (*WxMaPubTemplateTitleListResult, error) {
	param := map[string]string{}
	param["ids"] = strings.Join(ids, ",")
	param["start"] = strconv.Itoa(start)
	param["limit"] = strconv.Itoa(limit)
	url := common.MaGetPubTemplateTitleListUrl
	for k, v := range param {
		url += "&" + k + "=" + v
	}

	var res WxMaPubTemplateTitleListResult
	err := i.Do(url, &res)
	return &res, err
}

func (i *WxMaSubscribeServiceImpl) GetPubTemplateKeyWordsById(id string) (*WxMaPubTemplateKeywordListResult, error) {
	url := common.MaGetPubTemplateKeyWordsByIdUrl + "&tid=" + id

	var res WxMaPubTemplateKeywordListResult
	err := i.Do(url, &res)
	return &res, err
}

func (i *WxMaSubscribeServiceImpl) GetTemplateList() (*WxMaTemplateListResult, error) {
	var res WxMaTemplateListResult
	err := i.Do(common.MaTemplateListUrl, &res)
	return &res, err
}

func (i *WxMaSubscribeServiceImpl) AddTemplate(id, sceneDesc string, keywordIds []string) (*WxMaAddTemplateResult, error) {
	url := common.MaTemplateAddUrl

	param := map[string]interface{}{
		"tid":       id,
		"kidList":   keywordIds,
		"sceneDesc": sceneDesc,
	}

	at, err := i.service.GetAccessToken()
	if err != nil {
		return nil, err
	}

	var res WxMaAddTemplateResult
	err = i.service.PostFor(&res, url, common.PostJsonContentType, param, at.AccessToken)
	return &res, err
}

func (i *WxMaSubscribeServiceImpl) DelTemplate(templateId string) error {
	url := common.MaTemplateDelUrl

	param := map[string]string{
		"priTmplId": templateId,
	}

	at, err := i.service.GetAccessToken()
	if err != nil {
		return err
	}

	var res common.Err
	err = i.service.PostFor(&res, url, common.PostJsonContentType, param, at.AccessToken)
	if err != nil {
		return err
	}
	return &res
}

func (i *WxMaSubscribeServiceImpl) GetCategory() (*WxMaCategoryListResult, error) {
	var res WxMaCategoryListResult
	err := i.Do(common.MaGetCategoryUrl, &res)
	return &res, err
}

func (i *WxMaSubscribeServiceImpl) Do(url string, res error) error {
	return i.service.Do(url, res)
}
