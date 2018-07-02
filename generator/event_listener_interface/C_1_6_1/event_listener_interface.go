package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventListenerInterfaceGenerator_C_1_6_1 ...
func EventListenerInterfaceGenerator_C_1_6_1(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponentWithEvent() {
		for _, ctx := range c.GetContext() {
			slice = append(slice, proton.NewFileInfo("Events/Interfaces/I"+ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+c.GetID().WithoutComponentSuffix().String()+"Listener.cs", EventListenerInterface_C_1_6_1(ctx, c, new(bytes.Buffer)), "EventListenerInterfaceGenerator_C_1_6_1"))
		}
	}
	return slice, nil
}
