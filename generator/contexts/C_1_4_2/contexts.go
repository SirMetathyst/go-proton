package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ContextsGenerator_C_1_4_2 ...
func ContextsGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", Contexts_C_1_4_2(m.GetContext(), new(bytes.Buffer)), "ContextsGenerator_C_1_4_2"))
	return slice, nil
}
