package postprocessor

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

// fileContent ...
type fileContent interface {
	FileContent() string
}

// PrintFileContentPostProcessor_C_1_4_2 ...
func PrintFileContentPostProcessor_C_1_4_2(p *codegeneration.P, fi []proton.FI) ([]proton.FI, error) {
	for _, f := range fi {
		fmt.Println(f.FileContent())
	}

	return fi, nil
}
