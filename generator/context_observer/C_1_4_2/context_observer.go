package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ContextObserverGenerator_C_1_4_2 ...
func ContextObserverGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Contexts.cs", ContextObserver_C_1_4_2(md.ContextList(), new(bytes.Buffer)), "ContextObserverGenerator_C_1_4_2"))
	return slice, nil
}
