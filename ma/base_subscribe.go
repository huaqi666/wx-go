package ma

import "github.com/cliod/wx-go/common"

type WxMaPubTemplateTitleListResult struct {
	common.Err
	Count uint64          `json:"count"`
	Data  []*TemplateItem `json:"data"`
}

type WxMaPubTemplateKeywordListResult struct {
	common.Err
	Count uint64                `json:"count"`
	Data  []*PubTemplateKeyword `json:"data"`
}

type WxMaTemplateListResult struct {
	common.Err
	Data []*TemplateInfo `json:"data"`
}

type WxMaCategoryListResult struct {
	common.Err
	Data []*CategoryData `json:"data"`
}

type WxMaAddTemplateResult struct {
	common.Err
	PriTmplId string `json:"pri_tmpl_id"`
}

type TemplateItem struct {
	TypeNum    uint64 `json:"type"`
	Tid        uint64 `json:"tid"`
	CategoryId string `json:"category_id"`
	Title      string `json:"title"`
}

type PubTemplateKeyword struct {
	Kid     uint64 `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

type TemplateInfo struct {
	PriTmplId string `json:"pri_tmpl_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	TypeNum   uint64 `json:"type"`
}

type CategoryData struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
