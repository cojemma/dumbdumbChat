package live2drive

import (
	"dumbdumbChat/model"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

var (
	live2dModelFile = "./static/live2d/"
)

func ListAllModel() map[string]string {
	modelDirs, _ := filepath.Glob(live2dModelFile + "*")
	allModel := map[string]string{}

	for _, d := range modelDirs {
		info, _ := os.Stat(d)
		modelFile, _ := filepath.Glob(d + "/*model*")

		if len(modelFile) > 0 {
			fileURL := strings.TrimLeft(modelFile[0], ".")
			fileURL = strings.ReplaceAll(fileURL, "\\", "/")
			allModel[info.Name()] = fileURL
		}
	}

	return allModel
}

func SaveModelConfig(config model.Live2dCharacterConfig) {
	filePath := live2dModelFile + config.ModelName + "/live2dConfig.json"
	file, _ := os.Create(filePath)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.Encode(config)
}

func LoadModelConfig(modelName string) model.Live2dCharacterConfig {
	filePath := live2dModelFile + modelName + "/live2dConfig.json"
	file, _ := os.Open(filePath)
	config := model.Live2dCharacterConfig{}

	encoder := json.NewDecoder(file)
	encoder.Decode(&config)

	return config
}
