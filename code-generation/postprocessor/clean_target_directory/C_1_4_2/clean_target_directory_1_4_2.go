package postprocessor

import (
	"os"

	proton "github.com/SirMetathyst/go-proton"
)

var (
	CleanTargetDirectory = "src-gen"
)

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(fi []proton.FI) ([]proton.FI, error) {
	os.RemoveAll(CleanTargetDirectory)
	return fi, nil
}
