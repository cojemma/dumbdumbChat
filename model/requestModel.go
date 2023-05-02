package model

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatAPIRequest struct {
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type SetKeyRequest struct {
	ChatKey   string `json:"chatKey"`
	TTSKey    string `json:"ttsKey"`
	TTSRegion string `json:"ttsRegion"`
}
