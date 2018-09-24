package postprocessor

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

// FileHeaderPostProcessor_C_1_4_2 ...
func FileHeaderPostProcessor_C_1_4_2(fi []entitas.FI) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, f := range fi {
		b := new(bytes.Buffer)
		FileHeader_1_4_2(f.FileContent(), f.Generator(), b)
		slice = append(slice, entitas.NewFileInfo(f.File(), b.String(), f.FileContent()))
	}
	return slice, nil
}
