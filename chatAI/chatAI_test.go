package chatAI

import (
	"live2dViewer/model"
	"testing"
)

func TestSendToChat(t *testing.T) {
	type args struct {
		chatMessage []model.ChatMessage
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				[]model.ChatMessage{
					{
						Role:    "system",
						Content: `請扮演一名遊戲實況主，並且正在直播一款名為FF14的mmorpg。並且在最後要以json格式輸出你的情緒，包含happiness, sadness, anger, fear, surprise, disgust六個基本維度，量級為1~10。`,
					},
					{
						Role:    "user",
						Content: "安安，今天打算直播什麼遊戲啊?",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendToChatAI(tt.args.chatMessage)
		})
	}
}
