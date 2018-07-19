package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentMatcherGenerator_C_1_4_2 ...
func ComponentMatcherGenerator_C_1_4_2(m *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range m.ComponentList() {
		for _, c := range cp.ContextList() {
			slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
