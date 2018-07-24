package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ComponentMatcherGenerator_C_1_4_2 ...
func ComponentMatcherGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		for _, c := range cp.ContextList() {
			slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
