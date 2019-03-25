package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpEventListenerInterfaceGenerator_C_1_6_1", EventListenerInterfaceGenerator_C_1_6_1, false)
}

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
