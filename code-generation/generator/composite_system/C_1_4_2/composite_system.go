package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpCompositeSystemGenerator_C_1_6_1", CompositeSystemGenerator_C_1_4_2, false)
}

// CompositeSystemGenerator_C_1_4_2 ...
func CompositeSystemGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("CompositeSystem.cs", CompositeSystem_C_1_4_2(new(bytes.Buffer)), "CompositeSystemGenerator_C_1_4_2"))
	return slice, nil
}
