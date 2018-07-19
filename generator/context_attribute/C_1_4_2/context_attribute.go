package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ContextAttributeGenerator_C_1_4_2 ...
func ContextAttributeGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Attribute.cs", ContextAttribute_C_1_4_2(c, new(bytes.Buffer)), "ContextAttributeGenerator_C_1_4_2"))
	}
	return slice, nil
}
