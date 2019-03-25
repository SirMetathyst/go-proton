package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-proton"
	"github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentGenerator_E_1_4_2", ComponentGenerator_E_1_4_2, false)
}

// eventComponentID ...
func eventComponentID(cp *proton.CP) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	return proton.String(proton.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentGenerator_E_1_4_2 ...
func ComponentGenerator_E_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	for _, cp := range md.ComponentSlice() {
		slice = append(slice, proton.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		if cp.EventTarget() != proton.NoTarget {
			slice = append(slice, proton.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+"Listener.cs", EventComponent_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		}
	}
	return slice, nil
}
