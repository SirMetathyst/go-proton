package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventSystemGenerator_C_1_6_1 ...
func EventSystemGenerator_C_1_6_1(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsEvent() {
			for _, c := range cp.ContextList() {
				b := new(bytes.Buffer)
				if cp.EventTarget() == model.AnyTarget {
					EventSystemAnyTarget_C_1_6_1(c, cp, b)
				} else {
					EventSystemSelfTarget_C_1_6_1(c, cp, b)
				}
				slice = append(slice, proton.NewFileInfo("Events/Systems/"+ComponentID(c, cp).String()+".cs", b.String(), "EventSystemGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
