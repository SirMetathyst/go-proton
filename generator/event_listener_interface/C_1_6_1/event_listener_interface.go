package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// EventListenerInterfaceGenerator_C_1_6_1 ...
func EventListenerInterfaceGenerator_C_1_6_1(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentSlice() {
		if cp.IsEvent() {
			for _, c := range cp.ContextSlice() {
				slice = append(slice, entitas.NewFileInfo("Events/Interfaces/I"+componentID(c, cp).String()+".cs", EventListenerInterface_C_1_6_1(c, cp, new(bytes.Buffer)), "EventListenerInterfaceGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
