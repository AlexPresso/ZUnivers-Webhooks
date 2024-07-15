package discord

type WebhookFormData struct {
	Content     string   `json:"content"`
	Username    string   `json:"username"`
	AvatarURL   string   `json:"avatar_url"`
	TTS         bool     `json:"tts"`
	Embeds      *[]Embed `json:"embeds"`
	PayloadJSON string   `json:"payload_json"`
}

type Embed struct {
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	Timestamp   string        `json:"timestamp"`
	Color       uint32        `json:"color"`
	Fields      []*EmbedField `json:"fields"`
	Footer      *EmbedFooter  `json:"footer"`
	Author      *Author       `json:"author"`
	Thumbnail   *EmbedMedia   `json:"thumbnail"`
	Image       *EmbedMedia   `json:"image"`
	Role        string
}

type EmbedMedia struct {
	Url string `json:"url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type Author struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}
