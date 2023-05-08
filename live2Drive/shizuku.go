package live2drive

import (
	"dumbdumbChat/model"
	"fmt"
)

type shizukuExpression = string

const (
	shizukuExpressionSmile    shizukuExpression = "f01"
	shizukuExpressionUnhappy  shizukuExpression = "f02"
	shizukuExpressionAnger    shizukuExpression = "f03"
	shizukuExpressionSurprise shizukuExpression = "f04"
)

type shizukuMotion = string

const (
	shizukuMotionIdle      shizukuMotion = "idle"
	shizukuMotionTapBody   shizukuMotion = "tap_body"
	shizukuMotionPinchIn   shizukuMotion = "pinch_in"
	shizukuMotionPinchOut  shizukuMotion = "pinch_out"
	shizukuMotionShake     shizukuMotion = "shake"
	shizukuMotionFlickHead shizukuMotion = "flick_head"
)

var shizukuDriver = model.EmotionDriver{
	model.Happiness: {Motion: shizukuMotionTapBody, Expresion: shizukuExpressionSurprise},
	model.Sadness:   {Motion: shizukuMotionShake, Expresion: shizukuExpressionUnhappy},
	model.Anger:     {Motion: shizukuMotionPinchOut, Expresion: shizukuExpressionAnger},
	model.Fear:      {Motion: shizukuMotionPinchIn, Expresion: shizukuExpressionUnhappy},
	model.Surprise:  {Motion: shizukuMotionPinchIn, Expresion: shizukuExpressionSurprise},
	model.Disgust:   {Motion: shizukuMotionPinchOut, Expresion: shizukuExpressionUnhappy},
}

type Shizuku struct {
	emotionDriver model.EmotionDriver
	lastStatus    map[model.Emotion]int
}

func (char *Shizuku) GetDriverbyEmotion(emotion map[model.Emotion]int) model.ModelDriver {
	topEmos := model.GetTopEmotion(emotion, 1)
	char.lastStatus = emotion
	var topEmo model.Emotion
	if len(topEmos) > 0 {
		topEmo = topEmos[0]
	}

	driver := shizukuDriver[topEmo]

	fmt.Println(emotion, topEmo)
	if emotion[topEmo] <= 5 {
		driver.Motion = shizukuMotionIdle
	}

	return driver
}

func (char *Shizuku) SetEmotionDriver(emotionDriver model.EmotionDriver) {
	char.emotionDriver = shizukuDriver
}
