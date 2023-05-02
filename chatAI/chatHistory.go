package chatAI

import (
	"encoding/json"
	"live2dViewer/model"
	"os"
)

func recordChatHistory(chatMessages []model.ChatMessage) {
	file, _ := os.Create("./chatAI/chathistory.json")
	defer file.Close()

	file.Seek(0, 0) // 將檔案指標移至開頭
	encoder := json.NewEncoder(file)
	encoder.Encode(chatMessages)
}

func GetChatHistory() []model.ChatMessage {
	file, _ := os.Open("./chatAI/chathistory.json")
	defer file.Close()

	chatHistory := []model.ChatMessage{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&chatHistory)

	return chatHistory
}
