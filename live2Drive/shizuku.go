package live2drive

import (
	"fmt"
	"live2dViewer/model"
)

type shizukuExpression = string

const (
	shizukuExpressionSmile    shizukuExpression = "f01"
	shizukuExpressionUnhappy  shizukuExpression = "f02"
	shizukuExpressionAnger    shizukuExpression = "f03"
	shizukuExpressionSurprise shizukuExpression = "f04"
)

var shizukuExpressionMap = map[model.Emotion]shizukuExpression{
	model.Happiness: shizukuExpressionSurprise,
	model.Sadness:   shizukuExpressionUnhappy,
	model.Anger:     shizukuExpressionAnger,
	model.Fear:      shizukuExpressionUnhappy,
	model.Surprise:  shizukuExpressionSurprise,
	model.Disgust:   shizukuExpressionUnhappy,
}

type shizukuMotion = string

const (
	shizukuMotionIdle      shizukuMotion = "idle"
	shizukuMotionTapBody   shizukuMotion = "tap_body"
	shizukuMotionPinchIn   shizukuMotion = "pinch_in"
	shizukuMotionPinchOut  shizukuMotion = "pinch_out"
	shizukuMotionShake     shizukuMotion = "shake"
	shizukuMotionFlickHead shizukuMotion = "flick_head"
)

var shizukuMotionMap = map[model.Emotion]shizukuMotion{
	model.Happiness: shizukuMotionTapBody,
	model.Sadness:   shizukuMotionShake,
	model.Anger:     shizukuMotionPinchOut,
	model.Fear:      shizukuMotionPinchIn,
	model.Surprise:  shizukuMotionPinchIn,
	model.Disgust:   shizukuMotionPinchOut,
}

type Shizuku struct {
	lastStatus map[model.Emotion]int
}

func (char *Shizuku) GetDriverbyEmotion(emotion map[model.Emotion]int) model.ModelDriver {
	topEmos := model.GetTopEmotion(emotion, 1)
	char.lastStatus = emotion
	var topEmo model.Emotion
	if len(topEmos) > 0 {
		topEmo = topEmos[0]
	}

	driver := model.ModelDriver{
		Motion:    shizukuMotionMap[topEmo],
		Expresion: shizukuExpressionMap[topEmo],
	}

	fmt.Println(emotion, topEmo)
	if emotion[topEmo] <= 5 {
		driver.Motion = shizukuMotionIdle
	}

	return driver
}
