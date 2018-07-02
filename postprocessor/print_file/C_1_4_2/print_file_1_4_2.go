package postprocessor

import (
	"fmt"

	"github.com/SirMetathyst/proton"
)

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	for _, cf := range f {
		fmt.Println(cf.File)
	}

	return f, nil
}
