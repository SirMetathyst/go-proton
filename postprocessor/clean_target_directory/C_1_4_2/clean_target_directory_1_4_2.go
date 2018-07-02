package postprocessor

import (
	"os"

	"github.com/SirMetathyst/proton"
)

var (
	Directory = "Generated"
)

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	os.RemoveAll(Directory)
	return f, nil
}
