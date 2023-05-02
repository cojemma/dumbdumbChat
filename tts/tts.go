package tts

import (
	"bytes"
	"dumbdumbChat/model"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	outputMP3 = "./static/audio/output.mp3"
)

func init() {
	os.Mkdir("./static/audio/", 0755)
}

var (
	speechKey    = ""
	speechRegion = ""
)

func SetKeyandRegion(key, region string) {
	speechKey = key
	speechRegion = region
}

var xiaoXiaoEmoMap = map[model.Emotion]string{
	model.Happiness: "chat",
	model.Sadness:   "sad",
	model.Anger:     "angry",
	model.Fear:      "fearful",
	model.Surprise:  "cheerful",
	model.Disgust:   "disgruntled",
}

func TexttoSpeech(text string, emo map[model.Emotion]int) string {
	url := fmt.Sprintf("https://%s.tts.speech.microsoft.com/cognitiveservices/v1", speechRegion)

	topeEmo := model.GetTopEmotion(emo, 1)

	ttsReq := fmt.Sprintf(`<speak version='1.0' xml:lang='en-US'> 
	<voice xml:lang='zh-CN' xml:gender='Female' name='zh-CN-XiaoxiaoNeural' style='%s'> 
		%s
	</voice> 
</speak>`, xiaoXiaoEmoMap[topeEmo[0]], text)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(ttsReq))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", speechKey)
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("X-Microsoft-OutputFormat", "audio-16khz-128kbitrate-mono-mp3")
	req.Header.Add("User-Agent", "curl")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	mp3, _ := os.Create(outputMP3)
	defer mp3.Close()

	mp3.Write(body)
	// err = os.WriteFile(outputMP3, body, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return outputMP3
}
