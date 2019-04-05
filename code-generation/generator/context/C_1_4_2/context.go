package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpContextGenerator_C_1_4_2", ContextGenerator_C_1_4_2, false)
}

// ContextGenerator_C_1_4_2 ...
func ContextGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithContextSuffix().ToUpperFirst().String()+".cs", Context_C_1_4_2(c, new(bytes.Buffer)), "ContextGenerator_C_1_4_2"))
	}
	return slice, nil
}
