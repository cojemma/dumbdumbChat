package live2drive

import (
	"fmt"
	"testing"
)

func TestListAllModel(t *testing.T) {
	live2dModelFile = "../static/live2d/"

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test list all model files",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(ListAllModel())
		})
	}
}
