package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentGenerator_E_1_4_2", ComponentGenerator_E_1_4_2, false)
}

// eventComponentID ...
func eventComponentID(cp *proton.Component) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.EventTypeRemoved {
		eventTypeSuffix = "Removed"
	}
	return proton.String(proton.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentGenerator_E_1_4_2 ...
func ComponentGenerator_E_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentSlice() {
		slice = append(slice, proton.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		if cp.EventTarget() != proton.EventTargetNone {
			slice = append(slice, proton.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+"Listener.cs", EventComponent_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		}
	}
	return slice, nil
}
