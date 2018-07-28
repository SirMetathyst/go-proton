package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ContextMatcherGenerator_E_1_4_2 ...
func ContextMatcherGenerator_E_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Matcher.cs", ContextMatcher_E_1_4_2(c, new(bytes.Buffer)), "ContextMatcherGenerator_E_1_4_2"))
	}
	return slice, nil
}