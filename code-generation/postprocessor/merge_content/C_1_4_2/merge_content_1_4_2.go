package postprocessor

import proton "github.com/SirMetathyst/go-proton"

// MergeContentPostProcessor_C_1_4_2 ...
func MergeContentPostProcessor_C_1_4_2(fi []proton.FI) ([]proton.FI, error) {
	fl := make(map[string][]proton.FI)
	inl := make([]proton.FI, 0)

	for _, f := range fi {
		fl[f.File()] = append(fl[f.File()], proton.NewFileInfo(f.File(), f.FileContent(), f.Generator()))
	}

	for k, v := range fl {

		file := k
		fileContent := ""
		generator := ""

		for i, f := range v {
			fileContent += f.FileContent()
			generator += f.Generator()
			if i != len(v)-1 {
				generator += ","
			}
		}

		inl = append(inl, proton.NewFileInfo(file, fileContent, generator))
	}

	return inl, nil
}
