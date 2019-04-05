package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentMatcherGenerator_C_1_4_2", ComponentMatcherGenerator_C_1_4_2, false)
}

// eventComponentID ...
func eventComponentID(c *proton.C, cp *proton.CP) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
		optionalContextID = c.ID().String()
	}
	return proton.String(optionalContextID + proton.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentMatcherGenerator_C_1_4_2 ...
func ComponentMatcherGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, cp := range md.ComponentSlice() {
		for _, c := range cp.ContextSlice() {
			slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, false, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
			if cp.EventTarget() != proton.NoTarget {
				slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+eventComponentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, true, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
