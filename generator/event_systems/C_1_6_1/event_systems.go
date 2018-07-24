package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// EventSystemsGenerator_C_1_6_1 ...
func EventSystemsGenerator_C_1_6_1(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextList() {
		ecplist := make([]*entitas.CP, 0)
		cplist := md.ComponentsWithContextID(c.ID().String())
		for _, cp := range cplist {
			if cp.IsEvent() {
				ecplist = append(ecplist, cp)
			}
		}
		if len(ecplist) > 0 {
			slice = append(slice, entitas.NewFileInfo("Events/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+"EventSystems.cs", EventSystems_C_1_6_1(c, ecplist, new(bytes.Buffer)), "EventSystemsGenerator_C_1_6_1"))
		}
	}
	return slice, nil
}
