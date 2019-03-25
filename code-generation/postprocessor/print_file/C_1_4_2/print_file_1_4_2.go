package postprocessor

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
)

// file ...
type file interface {
	File() string
}

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(fi []proton.FI) ([]proton.FI, error) {
	for _, f := range fi {
		fmt.Println(f.File())
	}
	return fi, nil
}
