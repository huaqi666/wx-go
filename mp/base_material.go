package mp

// 永久图文消息素材
type WxMaterialNewsBatchGetNewsItem struct {
	MediaId    string           `json:"media_id"`
	UpdateTime string           `json:"update_time"`
	Content    WxMpMaterialNews `json:"content"`
}

type WxMpMaterialNews struct {
	CreateTime string             `json:"create_time"`
	UpdateTime string             `json:"update_time"`
	Articles   []*WxMpNewsArticle `json:"news_item"`
}

type WxMpNewsArticle struct {
	ThumbMediaId       string `json:"thumb_media_id"`
	ThumbUrl           string `json:"thumb_url"`
	Author             string `json:"author"`
	Title              string `json:"title"`
	ContentSourceUrl   string `json:"content_source_url"`
	Content            string `json:"content"`
	Digest             string `json:"digest"`
	ShowCoverPic       bool   `json:"show_cover_pic"`
	Url                string `json:"url"`
	NeedOpenComment    bool   `json:"need_open_comment"`
	OnlyFansCanComment bool   `json:"only_fans_can_comment"`
}

// 永久媒体消息素材
type WxMaterialFileBatchGetNewsItem struct {
	MediaId    string `json:"media_id"`
	UpdateTime string `json:"update_time"`
	Name       string `json:"name"`
	Url        string `json:"url"`
}
