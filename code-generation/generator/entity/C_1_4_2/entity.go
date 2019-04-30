package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpEntityGenerator_C_1_4_2", EntityGenerator_C_1_4_2, false)
}

// EntityGenerator_C_1_4_2 ...
func EntityGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Entity.cs", Entity_C_1_4_2(c, new(bytes.Buffer)), "EntityGenerator_C_1_4_2"))
	}
	return slice, nil
}
