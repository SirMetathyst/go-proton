package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentMatcherGenerator_C_1_4_2 ...
func ComponentMatcherGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponent() {
		for _, ctx := range c.GetContext() {
			slice = append(slice, proton.NewFileInfo(ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+c.GetID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(ctx, c, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
