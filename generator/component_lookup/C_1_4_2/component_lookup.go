package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ComponentLookupGenerator_C_1_4_2 ...
func ComponentLookupGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"ComponentsLookup.cs", ComponentLookup_C_1_4_2(c, md.ComponentsWithContextID(c.ID().String()), new(bytes.Buffer)), "ComponentLookupGenerator_C_1_4_2"))
	}
	return slice, nil
}
