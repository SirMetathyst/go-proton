package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// EventSystemsGenerator_C_1_6_1 ...
func EventSystemsGenerator_C_1_6_1(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, c := range md.ContextList() {
		ecplist := make([]*model.CP, 0)
		cplist := md.ComponentsWithContextID(c.ID().String())
		for _, cp := range cplist {
			if cp.IsEvent() {
				ecplist = append(ecplist, cp)
			}
		}
		if len(ecplist) > 0 {
			slice = append(slice, proton.NewFileInfo("Events/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"EventSystems.cs", EventSystems_C_1_6_1(c, ecplist, new(bytes.Buffer)), "EventSystemsGenerator_C_1_6_1"))
		}
	}
	return slice, nil
}
