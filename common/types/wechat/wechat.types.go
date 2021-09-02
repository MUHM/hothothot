package wechatTypes

type WechatPic struct {
	MediaId string `json:"media_id"`
}

type WechatMsg struct {
	Content string `json:"content"`
}

type WechatJsonData struct {
	ToUser                 string    `json:"touser"`
	AgentId                string    `json:"agentid"`
	MsgType                string    `json:"msgtype"`
	DuplicateCheckInterval int       `json:"duplicate_check_interval"`
	Text                   WechatMsg `json:"text"`
	Image                  WechatPic `json:"image"`
}
