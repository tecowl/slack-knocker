package slackknocker

type PayloadBase struct {
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

type Payload struct {
	Text string `json:"text"`
	PayloadBase
}
