package chatcore

import (
	"dumbdumbChat/chatAI"
	live2drive "dumbdumbChat/live2Drive"
	"dumbdumbChat/model"
	"dumbdumbChat/tts"
	"dumbdumbChat/wsMessage"
	"encoding/json"
	"fmt"
	"os"
)

var (
	keyFile                            = "./key.json"
	live2dModel live2drive.Live2Driver = &live2drive.Shizuku{}
)

func init() {
	keys := GetKey()
	SetKey(keys)
}

func ChatCore() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			ChatCore()
		}
	}()

	userMessageC := make(chan string)
	wsMessage.RegisterReceiveStream(userMessageC)

	for {
		userMessage := <-userMessageC
		chatAIResp := chatAI.ChatAI(userMessage)

		if chatAIResp.Message.Content == "" {
			continue
		}

		driver := live2dModel.GetDriverbyEmotion(chatAIResp.Emotion)

		ttsAudio := tts.TexttoSpeech(chatAIResp.Message.Content, chatAIResp.Emotion)
		ttsAudio = ttsAudio[1:]

		wsMessage.SendMessage(model.Chat, chatAIResp.Message)
		wsMessage.SendMessage(model.Live2Drive, driver)
		wsMessage.SendMessage(model.ChatVoice, ttsAudio)
	}
}

func SetKey(setKeyReq model.SetKeyRequest) {
	chatAI.SetOpenAIKey(setKeyReq.ChatKey)
	tts.SetKeyandRegion(setKeyReq.TTSKey, setKeyReq.TTSRegion)

	saveKey(setKeyReq)
}

func saveKey(setKeyReq model.SetKeyRequest) {
	file, _ := os.Create(keyFile)
	defer file.Close()

	file.Seek(0, 0) // 將檔案指標移至開頭
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.Encode(setKeyReq)
}

func GetKey() model.SetKeyRequest {
	file, _ := os.Open(keyFile)
	defer file.Close()

	keys := model.SetKeyRequest{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&keys)

	return keys
}

func ChangeLive2d(modelName string) {
	if modelName == "shizuku" {
		return
	}

	live2dModel = new(live2drive.SelfLoadModel)

	modelConfig := live2drive.LoadModelConfig(modelName)
	fmt.Printf("%+v\n", modelConfig)
	live2dModel.SetEmotionDriver(modelConfig.EmotionDriver)
}
