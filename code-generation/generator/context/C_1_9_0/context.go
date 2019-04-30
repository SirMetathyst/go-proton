package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpContextGenerator_C_1_9_0", ContextGenerator_C_1_9_0, false)
}

// ContextGenerator_C_1_9_0 ...
func ContextGenerator_C_1_9_0(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithContextSuffix().ToUpperFirst().String()+".cs", Context_C_1_9_0(c, new(bytes.Buffer)), "ContextGenerator_C_1_9_0"))
	}
	return slice, nil
}
