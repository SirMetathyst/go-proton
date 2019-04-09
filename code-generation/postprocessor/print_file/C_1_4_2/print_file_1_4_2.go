package postprocessor

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddPostProcessor("PrintFilePostProcessor_C_1_4_2", PrintFilePostProcessor_C_1_4_2, false)
}

// file ...
type file interface {
	File() string
}

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(p *codegeneration.P, fileInfo []proton.FileInfo) ([]proton.FileInfo, error) {
	for _, f := range fileInfo {
		fmt.Println(f.File())
	}
	return fileInfo, nil
}
