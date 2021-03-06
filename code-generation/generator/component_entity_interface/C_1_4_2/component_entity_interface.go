package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentEntityInterfaceGenerator_C_1_4_2", ComponentEntityInterfaceGenerator_C_1_4_2, false)
}

// ComponentEntityInterfaceGenerator_C_1_4_2 ...
func ComponentEntityInterfaceGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentSlice() {
		if len(cp.ContextSlice()) > 1 {
			b := new(bytes.Buffer)
			if len(cp.MemberSlice()) == 0 {
				ComponentEntityInterfaceFlag_C_1_4_2(cp, b)
			} else {
				ComponentEntityInterface_C_1_4_2(cp, b)
			}
			slice = append(slice, proton.NewFileInfo("Components/Interfaces/I"+cp.ID().WithoutComponentSuffix().ToUpperFirst().String()+"Entity.cs", b.String(), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			for _, c := range cp.ContextSlice() {
				slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentEntityInterfaceLink_C_1_4_2(c, cp, new(bytes.Buffer)), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
