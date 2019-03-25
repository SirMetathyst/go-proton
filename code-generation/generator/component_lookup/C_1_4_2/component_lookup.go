package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpComponentLookupGenerator_C_1_4_2", ComponentLookupGenerator_C_1_4_2, false)
}

// eventComponentID ...
func eventComponentID(c *entitas.C, cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextSlice()) > 1 {
		optionalContextID = c.ID().String()
	}
	return entitas.String(optionalContextID + entitas.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentLookupGenerator_C_1_4_2 ...
func ComponentLookupGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"ComponentsLookup.cs", ComponentLookup_C_1_4_2(c, md.ComponentsWithContextID(c.ID().String()), new(bytes.Buffer)), "ComponentLookupGenerator_C_1_4_2"))
	}
	return slice, nil
}
