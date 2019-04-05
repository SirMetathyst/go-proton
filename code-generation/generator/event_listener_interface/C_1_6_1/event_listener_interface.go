package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpEventListenerInterfaceGenerator_C_1_6_1", EventListenerInterfaceGenerator_C_1_6_1, false)
}

// EventListenerInterfaceGenerator_C_1_6_1 ...
func EventListenerInterfaceGenerator_C_1_6_1(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, cp := range md.ComponentSlice() {
		if cp.IsEvent() {
			for _, c := range cp.ContextSlice() {
				slice = append(slice, proton.NewFileInfo("Events/Interfaces/I"+componentID(c, cp).String()+".cs", EventListenerInterface_C_1_6_1(c, cp, new(bytes.Buffer)), "EventListenerInterfaceGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
