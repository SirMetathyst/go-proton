package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// EventSystemGenerator_C_1_6_1 ...
func EventSystemGenerator_C_1_6_1(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsEvent() {
			for _, c := range cp.ContextList() {
				b := new(bytes.Buffer)
				if cp.EventTarget() == entitas.AnyTarget {
					EventSystemAnyTarget_C_1_6_1(c, cp, b)
				} else {
					EventSystemSelfTarget_C_1_6_1(c, cp, b)
				}
				slice = append(slice, entitas.NewFileInfo("Events/Systems/"+componentID(c, cp).String()+".cs", b.String(), "EventSystemGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
