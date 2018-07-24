package postprocessor

import (
	"os"
)

var (
	CleanTargetDirectory = "Generated"
)

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	os.RemoveAll(CleanTargetDirectory)
	return v, nil
}
