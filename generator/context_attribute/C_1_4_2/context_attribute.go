package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ContextAttributeGenerator_C_1_4_2 ...
func ContextAttributeGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Attribute.cs", ContextAttribute_C_1_4_2(c, new(bytes.Buffer)), "ContextAttributeGenerator_C_1_4_2"))
	}
	return slice, nil
}
