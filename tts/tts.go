package tts

import (
	"bytes"
	"dumbdumbChat/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Mkdir("./static/audio/", 0755)
}

var (
	speechKey     = ""
	speechRegion  = ""
	outputMP3     = "./static/audio/output.mp3"
	ttsConfigFile = "./tts/config.json"
	ttsConfig     = model.TTSConfig{}
)

func init() {
	ttsConfig = GetTTSVoice()
}

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
	<voice xml:lang='%s' xml:gender='%s' name='%s' style='%s'> 
		%s
	</voice> 
</speak>`, ttsConfig.TTSLanguage, ttsConfig.TTSGender, ttsConfig.TTSVoiceName, xiaoXiaoEmoMap[topeEmo[0]], text)

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

	return outputMP3
}

func SetTTSVoice(config model.TTSConfig) {
	file, _ := os.Create(ttsConfigFile)
	defer file.Close()

	file.Seek(0, 0) // 將檔案指標移至開頭
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.Encode(config)
}

func GetTTSVoice() model.TTSConfig {
	file, _ := os.Open(ttsConfigFile)
	defer file.Close()

	config := model.TTSConfig{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	return config
}
