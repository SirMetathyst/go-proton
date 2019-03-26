package postprocessor

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

// file ...
type file interface {
	File() string
}

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(p *codegeneration.P, fi []proton.FI) ([]proton.FI, error) {
	for _, f := range fi {
		fmt.Println(f.File())
	}
	return fi, nil
}
