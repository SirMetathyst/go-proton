package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// EntityIndexGenerator_C_1_4_2 ...
func EntityIndexGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Contexts.cs", EntityIndex_C_1_4_2(md.ComponentsWithEntityIndex(), md.EntityIndexList(), new(bytes.Buffer)), "EntityIndexGenerator_C_1_4_2"))
	return slice, nil
}
