package chatAI

import (
	"dumbdumbChat/model"
	"fmt"
	"testing"
)

func Test_getChatHistory(t *testing.T) {
	chatHistoryFile = "./chathistory.json"

	tests := []struct {
		name string
		want []model.ChatMessage
	}{
		// TODO: Add test cases.
		{
			name: "chatHistory test",
			want: []model.ChatMessage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(GetChatHistory())
		})
	}
}
