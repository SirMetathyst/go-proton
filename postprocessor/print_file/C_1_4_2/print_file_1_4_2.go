package postprocessor

import (
	"fmt"

	entitas "github.com/SirMetathyst/go-entitas"
)

// file ...
type file interface {
	File() string
}

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(fi []entitas.FI) ([]entitas.FI, error) {
	for _, f := range fi {
		fmt.Println(f.File())
	}
	return fi, nil
}
