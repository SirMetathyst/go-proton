package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventListenerComponentEntityGenerator_C_1_6_1 ...
func EventListenerComponentEntityGenerator_C_1_6_1(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsEvent() {
			for _, c := range cp.ContextList() {
				slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+ComponentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", EventListenerComponentEntity_C_1_6_1(c, cp, new(bytes.Buffer)), "EventListenerEntityGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
