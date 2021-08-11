package ma

import "github.com/cliod/wx-go/common"

// 订阅的公共模板
type WxMaPubTemplateTitleListResult struct {
	common.WxCommonErr
	Count uint64          `json:"count"`
	Data  []*TemplateItem `json:"data"`
}

// 模板标题下关键词库结果
type WxMaPubTemplateKeywordListResult struct {
	common.WxCommonErr
	Count uint64                `json:"count"`
	Data  []*PubTemplateKeyword `json:"data"`
}

// 个人模板结果
type WxMaTemplateListResult struct {
	common.WxCommonErr
	Data []*TemplateInfo `json:"data"`
}

// 类目
type WxMaCategoryListResult struct {
	common.WxCommonErr
	Data []*CategoryData `json:"data"`
}

type WxMaAddTemplateResult struct {
	common.WxCommonErr
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
