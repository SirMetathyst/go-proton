package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// ComponentContextGenerator_1_4_2 ...
func ComponentContextGenerator_C_1_4_2(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if cp.IsUnique() {
			for _, c := range cp.ContextList() {
				b := new(bytes.Buffer)

				if len(cp.MemberList()) == 0 {
					ComponentContextFlag_C_1_4_2(c, cp, b)
				} else {
					ComponentContext_C_1_4_2(c, cp, b)
				}

				slice = append(slice, proton.NewFileInfo(c.ID().String()+"/Components/"+c.ID().String()+cp.ID().WithComponentSuffix().String()+".cs", b.String(), "ComponentContextGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
