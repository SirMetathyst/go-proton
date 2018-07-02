package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentEntityInterfaceGenerator_C_1_4_2 ...
func ComponentEntityInterfaceGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponent() {
		if len(c.GetContext()) > 1 {
			b := new(bytes.Buffer)
			if len(c.GetMember()) == 0 {
				ComponentEntityInterfaceFlag_C_1_4_2(c, b)
			} else {
				ComponentEntityInterface_C_1_4_2(c, b)
			}
			slice = append(slice, proton.NewFileInfo("Components/Interfaces/I"+c.GetID().WithoutComponentSuffix().ToUpperFirst().String()+"Entity.cs", b.String(), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			for _, ctx := range c.GetContext() {
				slice = append(slice, proton.NewFileInfo(ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+c.GetID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentEntityInterfaceLink_C_1_4_2(ctx, c, new(bytes.Buffer)), "ComponentEntityInterfaceGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
