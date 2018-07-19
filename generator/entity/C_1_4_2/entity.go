package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EntityGenerator_C_1_4_2 ...
func EntityGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Entity.cs", Entity_C_1_4_2(c, new(bytes.Buffer)), "EntityGenerator_C_1_4_2"))
	}
	return slice, nil
}
