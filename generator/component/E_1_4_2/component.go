package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentGenerator_E_1_4_2 ...
func ComponentGenerator_E_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		slice = append(slice, proton.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
	}
	return slice, nil
}
