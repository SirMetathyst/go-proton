package postprocessor

import (
	"bytes"
)

// FileHeaderPostProcessor_C_1_4_2 ...
func FileHeaderPostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	slice := make([]interface{}, 0)
	for _, cv := range v {
		f := cv.(fII)
		b := new(bytes.Buffer)
		FileHeader_1_4_2(f.FileContent(), f.Generator(), b)
		slice = append(slice, newFileInfo(f.File(), b.String(), f.FileContent()))
	}
	return slice, nil
}
