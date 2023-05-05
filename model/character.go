package model

type Emotion string

const (
	Happiness = "happiness"
	Sadness   = "sadness"
	Anger     = "anger"
	Fear      = "fear"
	Surprise  = "surprise"
	Disgust   = "disgust"
)

func ConvertEmotionStruct(emo EmotionStatus) map[Emotion]int {
	return map[Emotion]int{
		Happiness: emo.Happiness,
		Sadness:   emo.Sadness,
		Anger:     emo.Anger,
		Fear:      emo.Fear,
		Surprise:  emo.Surprise,
		Disgust:   emo.Disgust,
	}
}

func GetTopEmotion(emo map[Emotion]int, count int) []Emotion {
	emoCopy := make(map[Emotion]int)
	for k, v := range emo {
		emoCopy[k] = v
	}

	emoTops := []Emotion{}
	for i := 0; i < count; i++ {
		var maxEmotion Emotion
		for emotion, score := range emoCopy {
			if score >= emoCopy[maxEmotion] {
				maxEmotion = emotion
			}
		}
		emoTops = append(emoTops, maxEmotion)
		delete(emoCopy, maxEmotion)
	}

	return emoTops
}

type EmotionStatus struct {
	Happiness int `json:"happiness"`
	Sadness   int `json:"sadness"`
	Anger     int `json:"anger"`
	Fear      int `json:"fear"`
	Surprise  int `json:"surprise"`
	Disgust   int `json:"disgust"`
}

type Live2dCharacterConfig struct {
	ModelFile     string        `json:"modelFile"`
	EmotionDriver EmotionDriver `json:"emotionDriver"`
}

type EmotionDriver = map[Emotion]ModelDriver
