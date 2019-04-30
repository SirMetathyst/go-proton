package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpEventSystemGenerator_C_1_6_1", EventSystemGenerator_C_1_6_1, false)
}

// EventSystemGenerator_C_1_6_1 ...
func EventSystemGenerator_C_1_6_1(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentSlice() {
		if cp.IsEvent() {
			for _, c := range cp.ContextSlice() {
				b := new(bytes.Buffer)
				if cp.EventTarget() == proton.EventTargetAny {
					EventSystemAnyTarget_C_1_6_1(c, cp, b)
				} else {
					EventSystemSelfTarget_C_1_6_1(c, cp, b)
				}
				slice = append(slice, proton.NewFileInfo("Events/Systems/"+componentID(c, cp).String()+".cs", b.String(), "EventSystemGenerator_C_1_6_1"))
			}
		}
	}
	return slice, nil
}
