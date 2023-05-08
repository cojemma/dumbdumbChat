package live2drive

import (
	"dumbdumbChat/model"
	"fmt"
	"testing"
)

func TestListAllModel(t *testing.T) {
	live2dModelFile = "../static/live2d/"

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test list all model files",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(ListAllModel())
		})
	}
}

func TestSaveModelConfig(t *testing.T) {
	live2dConfigFile = "./live2dConfig.json"

	type args struct {
		config model.Live2dCharacterConfig
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "save model config",
			args: args{
				config: model.Live2dCharacterConfig{
					ModelName:     "shizuku",
					EmotionDriver: shizukuDriver,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveModelConfig(tt.args.config)
		})
	}
}
