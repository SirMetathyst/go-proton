package postprocessor

// MergeContentPostProcessor_C_1_4_2 ...
func MergeContentPostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	fl := make(map[string][]fI)
	inl := make([]fI, 0)

	for _, cv := range v {
		f := cv.(fII)
		fl[f.File()] = append(fl[f.File()], newFileInfo(f.File(), f.FileContent(), f.Generator()))
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

		inl = append(inl, newFileInfo(file, fileContent, generator))
	}

	outl := make([]interface{}, 0)
	for _, f := range inl {
		outl = append(outl, f)
	}

	return outl, nil
}
