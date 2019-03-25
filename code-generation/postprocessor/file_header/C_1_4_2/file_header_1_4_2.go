package postprocessor

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

// FileHeaderPostProcessor_C_1_4_2 ...
func FileHeaderPostProcessor_C_1_4_2(fi []proton.FI) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, f := range fi {
		b := new(bytes.Buffer)
		FileHeader_1_4_2(f.FileContent(), f.Generator(), b)
		slice = append(slice, proton.NewFileInfo(f.File(), b.String(), f.FileContent()))
	}
	return slice, nil
}
