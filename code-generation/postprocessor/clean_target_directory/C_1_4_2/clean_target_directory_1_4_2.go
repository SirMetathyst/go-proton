package postprocessor

import (
	"os"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddPostProcessor("CleanTargetDirectoryPostProcessor_C_1_4_2", CleanTargetDirectoryPostProcessor_C_1_4_2, true)
}

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(p *codegeneration.P, fileInfo []proton.FileInfo) ([]proton.FileInfo, error) {
	os.RemoveAll(p.OutputFolder())
	return fileInfo, nil
}
