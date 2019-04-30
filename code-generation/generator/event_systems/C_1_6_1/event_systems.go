package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpEventSystemsGenerator_C_1_6_1", EventSystemsGenerator_C_1_6_1, false)
}

// EventSystemsGenerator_C_1_6_1 ...
func EventSystemsGenerator_C_1_6_1(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextSlice() {
		ecplist := make([]*proton.Component, 0)
		cplist := md.ComponentsWithContextID(c.ID().String())
		for _, cp := range cplist {
			if cp.IsEvent() {
				ecplist = append(ecplist, cp)
			}
		}
		if len(ecplist) > 0 {
			slice = append(slice, proton.NewFileInfo("Events/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"EventSystems.cs", EventSystems_C_1_6_1(c, ecplist, new(bytes.Buffer)), "EventSystemsGenerator_C_1_6_1"))
		}
	}
	return slice, nil
}
