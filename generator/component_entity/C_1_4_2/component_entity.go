package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentEntityGenerator_C_1_4_2 ...
func ComponentEntityGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponent() {
		for _, ctx := range c.GetContext() {
			b := new(bytes.Buffer)

			if len(c.GetMember()) == 0 {
				ComponentEntityFlag_C_1_4_2(ctx, c, b)
			} else {
				ComponentEntity_C_1_4_2(ctx, c, b)
			}

			slice = append(slice, proton.NewFileInfo(ctx.GetID().WithoutContextSuffix().String()+"/Components/"+ctx.GetID().WithoutContextSuffix().String()+c.GetID().WithComponentSuffix().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
