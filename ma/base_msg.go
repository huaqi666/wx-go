package ma

type WxMaSubscribeMessage struct {
	ToUser           string           `json:"to_user"`
	TemplateId       string           `json:"template_id"`
	Page             string           `json:"page"`
	Data             []*Data          `json:"data"`
	MiniProgramState MiniProgramState `json:"miniprogram_state"`
	Lang             MiniProgramLang  `json:"lang"`
}

func (m *WxMaSubscribeMessage) AddData(data *Data) *WxMaSubscribeMessage {
	if m.Data == nil {
		m.Data = []*Data{data}
	} else {
		m.Data = append(m.Data, data)
	}
	return m
}

func NewWxMaSubscribeMessage(toUser, templateId, page string, arr ...*Data) *WxMaSubscribeMessage {
	return &WxMaSubscribeMessage{
		ToUser:           toUser,
		TemplateId:       templateId,
		Page:             page,
		Data:             arr,
		MiniProgramState: FORMAL,
		Lang:             ZhCn,
	}
}

func AddData(msg *WxMaSubscribeMessage, arr ...*Data) {
	for _, v := range arr {
		msg.AddData(v)
	}
}

type WxMaUniformMessage struct {
	IsMpTemplateMsg bool                `json:"is_mp_template_msg"`
	ToUser          string              `json:"to_user"`
	AppId           string              `json:"appid"`
	TemplateId      string              `json:"template_id"`
	Url             string              `json:"url"`
	Page            string              `json:"page"`
	FormId          string              `json:"form_id"`
	MiniProgram     MiniProgram         `json:"mini_program"`
	Data            []*WxMaTemplateData `json:"data"`
	EmphasisKeyword string              `json:"emphasis_keyword"`
}

type WxMaKefuMessage struct {
	ToUser  string   `json:"touser"`
	MsgType string   `json:"msgtype"`
	Text    *KfText  `json:"text"`
	Image   *KfImage `json:"image"`
	Link    *KfLink  `json:"link"`
	MaPage  KfMaPage `json:"miniprogrampage"`
}

type WxMaUpdatableMsg struct {
	ActivityId   string           `json:"activity_id"`
	TargetState  uint64           `json:"target_state"`
	TemplateInfo *MsgTemplateInfo `json:"template_info"`
}
