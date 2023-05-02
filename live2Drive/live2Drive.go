package live2drive

import (
	"dumbdumbChat/model"
)

type Live2Driver interface {
	GetDriverbyEmotion(map[model.Emotion]int) model.ModelDriver
}
