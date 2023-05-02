package tts

import (
	"live2dViewer/model"
	"testing"
)

func TestTexttoSpeech(t *testing.T) {
	type args struct {
		text string
		emo  map[model.Emotion]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "tts test",
			args: args{
				text: "我正在練習職業叫做黑魔導士，這個角色擅長使用黑魔法攻擊敵人，希望能夠熟練掌握它的技能！",
				emo: map[model.Emotion]int{
					model.Happiness: 7,
					model.Surprise:  5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TexttoSpeech(tt.args.text, tt.args.emo)
		})
	}
}
