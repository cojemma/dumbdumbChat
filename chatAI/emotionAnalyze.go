package chatAI

import (
	"encoding/json"
	"live2dViewer/model"
	"os"
)

func emotionAnalysis(messages []model.ChatMessage) map[model.Emotion]int {
	if len(messages) > 5 {
		messages = messages[len(messages)-4:]
	}

	analyzePrompts := getAnalyzePrompt()
	analyzePrompts = append(messages, analyzePrompts...)
	aiResp := sendToChatAI(analyzePrompts)
	analyzeResult := aiResp.Choices[0].Message.Content

	emotion := model.ConvertEmotionStruct(model.EmotionStatus{})

	json.Unmarshal([]byte(analyzeResult), &emotion)

	return emotion
}

func getAnalyzePrompt() []model.ChatMessage {
	file, _ := os.Open("./chatAI/emotionanalyze.json")
	defer file.Close()

	analyzePrompt := []model.ChatMessage{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&analyzePrompt)

	return analyzePrompt
}
