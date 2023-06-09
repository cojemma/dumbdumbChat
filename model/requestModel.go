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

type TTSConfig struct {
	TTSLanguage  string `json:"ttsLanguage"`
	TTSGender    string `json:"ttsGender"`
	TTSVoiceName string `json:"ttsVoiceName"`
}

type HtmlRenderer struct {
	SetKey          SetKeyRequest
	Live2dModelList map[string]string
	ChatHistory     []ChatMessage
	EmotionDriver   EmotionDriver
	TTSConfig       TTSConfig
}
