package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentContextGenerator_C_1_4_2", ComponentContextGenerator_C_1_4_2, false)
}

// ComponentContextGenerator_1_4_2 ...
func ComponentContextGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, cp := range md.ComponentSlice() {
		if cp.IsUnique() {
			for _, c := range cp.ContextSlice() {
				b := new(bytes.Buffer)

				if len(cp.MemberSlice()) == 0 {
					ComponentContextFlag_C_1_4_2(c, cp, b)
				} else {
					ComponentContext_C_1_4_2(c, cp, b)
				}

				slice = append(slice, proton.NewFileInfo(c.ID().ToUpperFirst().String()+"/Components/"+c.ID().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentContextGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
