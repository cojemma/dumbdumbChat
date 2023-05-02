package chatAI

import (
	"dumbdumbChat/model"
	"testing"
)

func Test_getChatHistory(t *testing.T) {
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
			GetChatHistory()
		})
	}
}
