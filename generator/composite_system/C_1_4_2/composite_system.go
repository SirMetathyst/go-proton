package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// CompositeSystemGenerator_C_1_4_2 ...
func CompositeSystemGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("CompositeSystem.cs", CompositeSystem_C_1_4_2(new(bytes.Buffer)), "CompositeSystemGenerator_C_1_4_2"))
	return slice, nil
}
