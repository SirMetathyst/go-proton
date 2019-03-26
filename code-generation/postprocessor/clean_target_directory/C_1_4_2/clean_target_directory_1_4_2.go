package postprocessor

import (
	"os"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

// CleanTargetDirectoryPostProcessor_C_1_4_2 ...
func CleanTargetDirectoryPostProcessor_C_1_4_2(p *codegeneration.P, fi []proton.FI) ([]proton.FI, error) {
	os.RemoveAll(p.OutputFolder())
	return fi, nil
}
