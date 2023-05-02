package live2drive

import (
	"live2dViewer/model"
)

type Live2Driver interface {
	GetDriverbyEmotion(map[model.Emotion]int) model.ModelDriver
}
