package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentEntityInterfaceGenerator_C_1_4_2 ...
func ComponentEntityInterfaceGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if len(cp.ContextList()) > 1 {
			b := new(bytes.Buffer)
			if len(cp.MemberList()) == 0 {
				ComponentEntityInterfaceFlag_C_1_4_2(cp, b)
			} else {
				ComponentEntityInterface_C_1_4_2(cp, b)
			}
			slice = append(slice, proton.NewFileInfo("Components/Interfaces/I"+cp.ID().WithoutComponentSuffix().ToUpperFirst().String()+"Entity.cs", b.String(), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			for _, c := range cp.ContextList() {
				slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentEntityInterfaceLink_C_1_4_2(c, cp, new(bytes.Buffer)), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
