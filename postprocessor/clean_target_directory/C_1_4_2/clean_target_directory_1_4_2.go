package postprocessor

import (
	"os"

	entitas "github.com/SirMetathyst/go-entitas"
)

var (
	CleanTargetDirectory = "Generated"
)

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(fi []entitas.FI) ([]entitas.FI, error) {
	os.RemoveAll(CleanTargetDirectory)
	return fi, nil
}
