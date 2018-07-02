package postprocessor

import (
	"fmt"

	"github.com/SirMetathyst/proton"
)

// PrintFileContentPostProcessor_C_1_4_2 ...
func PrintFileContentPostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	for _, cf := range f {
		fmt.Println(cf.FileContent)
	}

	return f, nil
}
