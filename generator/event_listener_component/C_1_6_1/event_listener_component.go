package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventListenerComponentGenerator_C_1_6_1 ...
func EventListenerComponentGenerator_C_1_6_1(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range m.GetComponentWithEvent() {
		for _, ctx := range c.GetContext() {
			slice = append(slice, proton.NewFileInfo("Events/Components/"+c.GetID().WithoutComponentSuffix().String()+"ListenerComponent.cs", EventListenerComponent_C_1_6_1(ctx, c, new(bytes.Buffer)), "EventListenerComponentGenerator_C_1_6_1"))
		}
	}
	return slice, nil
}
