package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpContextsGenerator_C_1_4_2", ContextsGenerator_C_1_4_2, false)
}

// ContextsGenerator_C_1_4_2 ...
func ContextsGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", Contexts_C_1_4_2(md.ContextSlice(), new(bytes.Buffer)), "ContextsGenerator_C_1_4_2"))
	return slice, nil
}
