package live2drive

import (
	"dumbdumbChat/model"
	"fmt"
)

type Live2Driver interface {
	GetDriverbyEmotion(map[model.Emotion]int) model.ModelDriver
	SetEmotionDriver(model.EmotionDriver)
}

type SelfLoadModel struct {
	emotionDriver model.EmotionDriver
	lastStatus    map[model.Emotion]int
}

func (m *SelfLoadModel) GetDriverbyEmotion(emotion map[model.Emotion]int) model.ModelDriver {
	topEmos := model.GetTopEmotion(emotion, 1)
	m.lastStatus = emotion
	var topEmo model.Emotion
	if len(topEmos) > 0 {
		topEmo = topEmos[0]
	}

	driver := m.emotionDriver[topEmo]

	fmt.Println(emotion, topEmo)

	return driver
}

func (m *SelfLoadModel) SetEmotionDriver(emotionDriver model.EmotionDriver) {
	copyEotionDriver := model.EmotionDriver{}
	for emo, driver := range emotionDriver {
		copyEotionDriver[emo] = driver
	}

	m.emotionDriver = copyEotionDriver
}
