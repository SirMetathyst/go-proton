package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpComponentLookupGenerator_C_1_4_2", ComponentLookupGenerator_C_1_4_2, false)
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

// ComponentLookupGenerator_C_1_4_2 ...
func ComponentLookupGenerator_C_1_4_2(md *proton.Model) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"ComponentsLookup.cs", ComponentLookup_C_1_4_2(c, md.ComponentsWithContextID(c.ID().String()), new(bytes.Buffer)), "ComponentLookupGenerator_C_1_4_2"))
	}
	return slice, nil
}
