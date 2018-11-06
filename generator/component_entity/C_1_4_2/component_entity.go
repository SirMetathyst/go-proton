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

// eventComponentInterfaceID ...
func eventComponentInterfaceID(cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	return entitas.String(entitas.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentEntityGenerator_C_1_4_2 ...
func ComponentEntityGenerator_C_1_4_2(m *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range m.ComponentList() {
		for _, c := range cp.ContextList() {
			b := new(bytes.Buffer)

			if len(cp.MemberList()) == 0 {
				ComponentEntityFlag_C_1_4_2(c, cp, b)
			} else {
				ComponentEntity_C_1_4_2(c, cp, false, b)
			}

			slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))

			if cp.EventTarget() != entitas.NoTarget {
				b = new(bytes.Buffer)
				ComponentEntity_C_1_4_2(c, cp, true, b)
				slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+eventComponentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", b.String(), "ComponentEntityGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
