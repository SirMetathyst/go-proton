package postprocessor

import (
	"fmt"
)

// fileContent ...
type fileContent interface {
	FileContent() string
}

// PrintFileContentPostProcessor_C_1_4_2 ...
func PrintFileContentPostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	for _, cv := range v {
		f := cv.(fileContent)
		fmt.Println(f.FileContent())
	}

	return v, nil
}
