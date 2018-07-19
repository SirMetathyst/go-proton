package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// CompositeSystemGenerator_C_1_4_2 ...
func CompositeSystemGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("CompositeSystem.cs", CompositeSystem_C_1_4_2(new(bytes.Buffer)), "CompositeSystemGenerator_C_1_4_2"))
	return slice, nil
}
