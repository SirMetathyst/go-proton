package postprocessor

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddPostProcessor("PrintFileContentPostProcessor_C_1_4_2", PrintFileContentPostProcessor_C_1_4_2, false)
}

// fileContent ...
type fileContent interface {
	FileContent() string
}

// PrintFileContentPostProcessor_C_1_4_2 ...
func PrintFileContentPostProcessor_C_1_4_2(p *codegeneration.P, fileInfo []proton.FileInfo) ([]proton.FileInfo, error) {
	for _, f := range fileInfo {
		fmt.Println(f.FileContent())
	}

	return fileInfo, nil
}
