package model

type ChatCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index        int         `json:"index"`
		Message      ChatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type WSMessageType int

const (
	Log WSMessageType = iota
	Live2Drive
	Chat
	ChatVoice
)

type ModelDriver struct {
	Motion    string `json:"motion"`
	Expresion string `json:"expression"`
}

type WSResponse struct {
	MessageType WSMessageType `json:"messageType"`
	Data        interface{}   `json:"data"`
}

type ChatResp struct {
	Message ChatMessage `json:"message"`
	Emotion map[Emotion]int
}
