package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpEntityIndexGenerator_C_1_4_2", EntityIndexGenerator_C_1_4_2, false)
}

// EntityIndexGenerator_C_1_4_2 ...
func EntityIndexGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Contexts.cs", EntityIndex_C_1_4_2(md.ComponentsWithEntityIndex(), md.EntityIndexList(), new(bytes.Buffer)), "EntityIndexGenerator_C_1_4_2"))
	return slice, nil
}
