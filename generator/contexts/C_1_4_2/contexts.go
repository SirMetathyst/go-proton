package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ContextsGenerator_C_1_4_2 ...
func ContextsGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Contexts.cs", Contexts_C_1_4_2(md.ContextList(), new(bytes.Buffer)), "ContextsGenerator_C_1_4_2"))
	return slice, nil
}