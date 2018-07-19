package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventListenerInterfaceGenerator_C_1_6_1 ...
func EventListenerInterfaceGenerator_C_1_6_1(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsEvent() {
			for _, c := range cp.ContextList() {
				slice = append(slice, proton.NewFileInfo("Events/Interfaces/I"+ComponentID(c, cp).String()+".cs", EventListenerInterface_C_1_6_1(c, cp, new(bytes.Buffer)), "EventListenerInterfaceGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
