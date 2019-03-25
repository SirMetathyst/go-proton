package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpContextsGenerator_C_1_4_2", ContextsGenerator_C_1_4_2, false)
}

// ContextsGenerator_C_1_4_2 ...
func ContextsGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Contexts.cs", Contexts_C_1_4_2(md.ContextSlice(), new(bytes.Buffer)), "ContextsGenerator_C_1_4_2"))
	return slice, nil
}
