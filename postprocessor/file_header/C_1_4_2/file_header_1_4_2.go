package postprocessor

import (
	"bytes"

	"github.com/SirMetathyst/proton"
)

// FileHeaderPostProcessor_C_1_4_2 ...
func FileHeaderPostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, v := range f {
		b := new(bytes.Buffer)
		FileHeader_1_4_2(v.FileContent, v.Generator, b)
		slice = append(slice, proton.NewFileInfo(v.File, b.String(), v.FileContent))
	}
	return slice, nil
}
