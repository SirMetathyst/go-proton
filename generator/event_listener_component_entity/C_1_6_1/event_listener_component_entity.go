package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// EventListenerComponentEntityGenerator_C_1_6_1 ...
func EventListenerComponentEntityGenerator_C_1_6_1(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsEvent() {
			for _, c := range cp.ContextList() {
				slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+componentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", EventListenerComponentEntity_C_1_6_1(c, cp, new(bytes.Buffer)), "EventListenerEntityGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
