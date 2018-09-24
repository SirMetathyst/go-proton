package postprocessor

import (
	"fmt"

	entitas "github.com/SirMetathyst/go-entitas"
)

// fileContent ...
type fileContent interface {
	FileContent() string
}

// PrintFileContentPostProcessor_C_1_4_2 ...
func PrintFileContentPostProcessor_C_1_4_2(fi []entitas.FI) ([]entitas.FI, error) {
	for _, f := range fi {
		fmt.Println(f.FileContent())
	}

	return fi, nil
}
