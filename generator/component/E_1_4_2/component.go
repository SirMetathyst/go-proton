package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// eventComponentID ...
func eventComponentID(c *entitas.C, cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextList()) > 1 {
		optionalContextID = c.ID().String()
	}
	return entitas.String(optionalContextID + entitas.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentGenerator_E_1_4_2 ...
func ComponentGenerator_E_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		slice = append(slice, entitas.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		if cp.EventTarget() != entitas.NoTarget {
			slice = append(slice, entitas.NewFileInfo("Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+"Listener.cs", EventComponent_E_1_4_2(cp, new(bytes.Buffer)), "ComponentGenerator_E_1_4_2"))
		}
	}
	return slice, nil
}
