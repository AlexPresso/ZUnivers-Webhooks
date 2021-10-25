package discord

type WebhookFormData struct {
	Content     string          `json:"content"`
	Username    string          `json:"username"`
	AvatarURL   string          `json:"avatar_url"`
	TTS         bool            `json:"tts"`
	Embeds      []*DiscordEmbed `json:"embeds"`
	PayloadJSON string          `json:"payload_json"`
}

type DiscordEmbed struct {
	Title       string                 `json:"title"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	URL         string                 `json:"url"`
	Timestamp   string                 `json:"timestamp"`
	Color       uint32                 `json:"color"`
	Fields      []*DiscordEmbedField   `json:"fields"`
	Footer      *DiscordEmbedFooter    `json:"footer"`
	Author      *DiscordAuthor         `json:"author"`
	Thumbnail   *DiscordEmbedThumbnail `json:"thumbnail"`
}

type DiscordEmbedThumbnail struct {
	Url string `json:"url"`
}

type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordEmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type DiscordAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}
