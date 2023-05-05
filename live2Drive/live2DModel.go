package live2drive

import (
	"os"
	"path/filepath"
	"strings"
)

var live2dModelFile = "./static/live2d/"

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
