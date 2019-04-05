package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {

	codegeneration.AddGenerator("CSharpCompositeSystemGenerator_C_1_6_1", CompositeSystemGenerator_C_1_4_2, false)
}

// CompositeSystemGenerator_C_1_4_2 ...
func CompositeSystemGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	slice = append(slice, proton.NewFileInfo("CompositeSystem.cs", CompositeSystem_C_1_4_2(new(bytes.Buffer)), "CompositeSystemGenerator_C_1_4_2"))
	return slice, nil
}
