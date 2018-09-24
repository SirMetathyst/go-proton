package postprocessor

import entitas "github.com/SirMetathyst/go-entitas"

// MergeContentPostProcessor_C_1_4_2 ...
func MergeContentPostProcessor_C_1_4_2(fi []entitas.FI) ([]entitas.FI, error) {
	fl := make(map[string][]entitas.FI)
	inl := make([]entitas.FI, 0)

	for _, f := range fi {
		fl[f.File()] = append(fl[f.File()], entitas.NewFileInfo(f.File(), f.FileContent(), f.Generator()))
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

		inl = append(inl, entitas.NewFileInfo(file, fileContent, generator))
	}

	return inl, nil
}
