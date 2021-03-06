package postprocessor

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddPostProcessor("FileHeaderPostProcessor_C_1_4_2", FileHeaderPostProcessor_C_1_4_2, true)
}

// FileHeaderPostProcessor_C_1_4_2 ...
func FileHeaderPostProcessor_C_1_4_2(p *codegeneration.P, fileInfo []proton.FileInfo) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, f := range fileInfo {
		b := new(bytes.Buffer)
		FileHeader_1_4_2(f.FileContent(), f.Generator(), b)
		slice = append(slice, proton.NewFileInfo(f.File(), b.String(), f.FileContent()))
	}
	return slice, nil
}
