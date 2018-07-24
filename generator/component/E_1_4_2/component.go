package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ComponentGenerator_E_1_4_2 ...
func ComponentGenerator_E_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		slice = append(slice, entitas.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
	}
	return slice, nil
}
