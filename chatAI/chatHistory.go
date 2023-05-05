package chatAI

import (
	"dumbdumbChat/model"
	"encoding/json"
	"os"
)

var chatHistoryFile = "./chatAI/chathistory.json"

func recordChatHistory(chatMessages []model.ChatMessage) {
	file, _ := os.Create(chatHistoryFile)
	defer file.Close()

	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.Encode(chatMessages)
}

func GetChatHistory() []model.ChatMessage {
	file, _ := os.Open(chatHistoryFile)
	defer file.Close()

	chatHistory := []model.ChatMessage{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&chatHistory)

	return chatHistory
}
