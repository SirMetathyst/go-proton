package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentEntityGenerator_C_1_4_2", ComponentEntityGenerator_C_1_4_2, false)
}

// eventComponentID ...
func eventComponentID(c *proton.Context, cp *proton.Component) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.EventTypeRemoved {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
		optionalContextID = c.ID().String()
	}
	return proton.String(optionalContextID + proton.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// eventComponentInterfaceID ...
func eventComponentInterfaceID(cp *proton.Component) proton.String {
	var eventTypeSuffix = ""
	if cp.EventType() == proton.EventTypeRemoved {
		eventTypeSuffix = "Removed"
	}
	return proton.String(proton.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentEntityGenerator_C_1_4_2 ...
func ComponentEntityGenerator_C_1_4_2(m *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range m.ComponentSlice() {
		for _, c := range cp.ContextSlice() {
			b := new(bytes.Buffer)

			if len(cp.MemberSlice()) == 0 {
				ComponentEntityFlag_C_1_4_2(c, cp, b)
			} else {
				ComponentEntity_C_1_4_2(c, cp, false, b)
			}

			slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))

			if cp.EventTarget() != proton.EventTargetNone {
				b = new(bytes.Buffer)
				ComponentEntity_C_1_4_2(c, cp, true, b)
				slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+eventComponentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
