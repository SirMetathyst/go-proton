package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpEntityIndexGenerator_C_1_4_2", EntityIndexGenerator_C_1_4_2, false)
}

// EntityIndexGenerator_C_1_4_2 ...
func EntityIndexGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", EntityIndex_C_1_4_2(md.ComponentsWithEntityIndex(), md.EntityIndexSlice(), new(bytes.Buffer)), "EntityIndexGenerator_C_1_4_2"))
	return slice, nil
}
