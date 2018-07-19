package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ContextGenerator_C_1_4_2 ...
func ContextGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextList() {
		slice = append(slice, proton.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithContextSuffix().ToUpperFirst().String()+".cs", Context_C_1_4_2(c, new(bytes.Buffer)), "ContextGenerator_C_1_4_2"))
	}
	return slice, nil
}
