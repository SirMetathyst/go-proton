package postprocessor

import (
	"github.com/SirMetathyst/proton"
)

// MergeContentPostProcessor_C_1_4_2 ...
func MergeContentPostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	fl := make(map[string][]proton.FileInfo)
	flst := make([]proton.FileInfo, 0)

	for _, cf := range f {
		fl[cf.File] = append(fl[cf.File], cf)
	}

	for k, v := range fl {

		NewFile := proton.FileInfo{}
		NewFile.File = k

		for i, f := range v {
			NewFile.FileContent += f.FileContent
			NewFile.Generator += f.Generator
			if i != len(v)-1 {
				NewFile.Generator += ","
			}
		}

		flst = append(flst, NewFile)
	}

	return flst, nil
}
