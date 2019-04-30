package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpContextObserverGenerator_C_1_4_2", ContextObserverGenerator_C_1_4_2, false)
}

// ContextObserverGenerator_C_1_4_2 ...
func ContextObserverGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", ContextObserver_C_1_4_2(md.ContextSlice(), new(bytes.Buffer)), "ContextObserverGenerator_C_1_4_2"))
	return slice, nil
}
