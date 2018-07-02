package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EntityIndexGenerator_C_1_4_2 ...
func EntityIndexGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", EntityIndex_C_1_4_2(m.GetComponentWithEntityIndex(), m.GetEntityIndex(), new(bytes.Buffer)), "EntityIndexGenerator_C_1_4_2"))
	return slice, nil
}
