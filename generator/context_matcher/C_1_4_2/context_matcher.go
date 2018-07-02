package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ContextMatcherGenerator_C_1_4_2 ...
func ContextMatcherGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, ctx := range m.GetContext() {
		slice = append(slice, proton.NewFileInfo(ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+"/"+ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+"Matcher.cs", ContextMatcher_C_1_4_2(ctx, new(bytes.Buffer)), "ContextMatcherGenerator_C_1_4_2"))
	}
	return slice, nil
}
