package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ComponentEntityGenerator_C_1_4_2 ...
func ComponentEntityGenerator_C_1_4_2(m *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range m.ComponentList() {
		for _, c := range cp.ContextList() {
			b := new(bytes.Buffer)

			if len(cp.MemberList()) == 0 {
				ComponentEntityFlag_C_1_4_2(c, cp, b)
			} else {
				ComponentEntity_C_1_4_2(c, cp, b)
			}

			slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))
		}
	}
	return slice, nil
}
