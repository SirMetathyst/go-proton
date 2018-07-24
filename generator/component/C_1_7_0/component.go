package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ComponentGenerator_C_1_7_0 ...
func ComponentGenerator_C_1_7_0(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		if !cp.IsListener() {
			slice = append(slice, entitas.NewFileInfo("../Generated_Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_C_1_7_0(cp, new(bytes.Buffer)), "ComponentGenerator_C_1_7_0"))
		}
	}
	return slice, nil
}
