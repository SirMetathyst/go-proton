package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpContextMatcherGenerator_C_1_4_2", ContextMatcherGenerator_C_1_4_2, false)
}

// ContextMatcherGenerator_C_1_4_2 ...
func ContextMatcherGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"Matcher.cs", ContextMatcher_C_1_4_2(c, new(bytes.Buffer)), "ContextMatcherGenerator_E_1_4_2"))
	}
	return slice, nil
}
