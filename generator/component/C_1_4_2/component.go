package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"

	"github.com/SirMetathyst/proton/model"
)

// ComponentGenerator_1_4_2 ...
func ComponentGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponent() {
		slice = append(slice, proton.NewFileInfo(c.GetID().AsComponentPath().String(), Component_C_1_4_2(c, new(bytes.Buffer)), "ComponentGenerator_C_1_4_2"))
	}
	return slice, nil
}
