package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentEntityGenerator_C_1_4_2 ...
func ComponentEntityGenerator_C_1_4_2(m *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range m.ComponentList() {
		for _, c := range cp.ContextList() {
			b := new(bytes.Buffer)

			if len(cp.MemberList()) == 0 {
				ComponentEntityFlag_C_1_4_2(c, cp, b)
			} else {
				ComponentEntity_C_1_4_2(c, cp, b)
			}

			slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().String()+"/Components/"+c.ID().WithoutContextSuffix().String()+cp.ID().WithComponentSuffix().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
