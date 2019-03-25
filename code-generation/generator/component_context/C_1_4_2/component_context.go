package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpComponentContextGenerator_C_1_4_2", ComponentContextGenerator_C_1_4_2, false)
}

// ComponentContextGenerator_1_4_2 ...
func ComponentContextGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentSlice() {
		if cp.IsUnique() {
			for _, c := range cp.ContextSlice() {
				b := new(bytes.Buffer)

				if len(cp.MemberSlice()) == 0 {
					ComponentContextFlag_C_1_4_2(c, cp, b)
				} else {
					ComponentContext_C_1_4_2(c, cp, b)
				}

				slice = append(slice, entitas.NewFileInfo(c.ID().ToUpperFirst().String()+"/Components/"+c.ID().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentContextGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
