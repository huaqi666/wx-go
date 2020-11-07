package ma

type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type WxMaTemplateData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Color string `json:"color"`
}

type MiniProgram struct {
	AppId       string `json:"appid"`
	PagePath    string `json:"page_path"`
	UsePath     bool   `json:"use_path"`
	UsePagePath bool   `json:"use_page_path"`
}

type KfText struct {
	Content string `json:"content"`
}

type KfImage struct {
	MediaId string `json:"media_id"`
}

type KfLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumb_url"`
}

type KfMaPage struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type MsgTemplateInfo struct {
	ParameterList []*Parameter `json:"parameter_list"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
