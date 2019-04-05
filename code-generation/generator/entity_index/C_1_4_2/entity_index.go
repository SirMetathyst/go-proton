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
func EntityIndexGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", EntityIndex_C_1_4_2(md.ComponentsWithEntityIndex(), md.EntityIndexList(), new(bytes.Buffer)), "EntityIndexGenerator_C_1_4_2"))
	return slice, nil
}
