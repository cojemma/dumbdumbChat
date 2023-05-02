package chatAI

import (
	"bytes"
	"dumbdumbChat/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	chatUrl   = "https://api.openai.com/v1/chat/completions"
	ChatModel = "gpt-3.5-turbo"
	RoleUser  = "user"
)

var openAIKey = ""

func SetOpenAIKey(key string) {
	openAIKey = key
}

func ChatAI(message string) model.ChatResp {
	chatHistory := GetChatHistory()
	userMessage := model.ChatMessage{
		Role:    RoleUser,
		Content: message,
	}

	chatReq := append(chatHistory, userMessage)

	fmt.Println(chatReq)
	chatAIResp := sendToChatAI(chatReq)
	chatAIMessage := chatAIResp.Choices[0].Message

	if len(chatAIResp.Choices) == 0 {
		return model.ChatResp{}
	}

	if chatAIResp.Usage.TotalTokens >= 4000 {
		chatHistory = append(chatHistory[:5], userMessage, chatAIMessage)
	} else {
		chatHistory = append(chatHistory, userMessage, chatAIMessage)
	}
	recordChatHistory(chatHistory)

	emotion := emotionAnalysis(chatHistory)
	chat := model.ChatResp{
		Message: chatAIMessage,
		Emotion: emotion,
	}

	return chat
}

func sendToChatAI(chatMessage []model.ChatMessage) model.ChatCompletion {
	chatReq := model.ChatAPIRequest{
		Model:     ChatModel,
		Messages:  chatMessage,
		MaxTokens: 200,
	}
	data, _ := json.Marshal(chatReq)

	req, err := http.NewRequest("POST", chatUrl, bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+openAIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Body)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	chatResp := model.ChatCompletion{}
	json.Unmarshal(body, &chatResp)

	return chatResp
}
