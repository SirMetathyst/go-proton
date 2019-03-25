package proton

import "fmt"

var (
	ErrComponentListTriedToAddNilComponent         = fmt.Errorf("entitas(component list): tried to add nil component")
	ErrComponentListTriedToAddDuplicateComponentID = fmt.Errorf("entitas(component list): tried to add component with duplicate id")
)

// CPL ...
type CPL struct {
	l []*CP
}

// NewComponentList ...
func NewComponentList() *CPL {
	return &CPL{}
}

// AddComponent ...
func (cpl *CPL) AddComponent(cp *CP) error {
	if cp == nil {
		return ErrComponentListTriedToAddNilComponent
	}
	if cpl.ComponentWithID(cp.ID().String()) != nil {
		return ErrComponentListTriedToAddDuplicateComponentID
	}
	cpl.l = append(cpl.l, cp)
	return nil
}

// ComponentWithID ...
func (cpl *CPL) ComponentWithID(id string) *CP {
	for _, cp := range cpl.l {
		if cp.ID().EqualTo(id) {
			return cp
		}
	}
	return nil
}

// ComponentsWithContextID ...
func (cpl *CPL) ComponentsWithContextID(id string) []*CP {
	slice := make([]*CP, 0)
	for _, cp := range cpl.l {
		c := cp.ContextWithID(id)
		if c != nil {
			slice = append(slice, cp)
		}
	}
	return slice
}

// ComponentsWithEntityIndex ...
func (cpl *CPL) ComponentsWithEntityIndex() []*CP {
	slice := make([]*CP, 0)
	for _, cp := range cpl.l {
		m := cp.MembersWithEntityIndex()
		if len(m) > 0 {
			slice = append(slice, cp)
		}
	}
	return slice
}

// ComponentSlice ...
func (cpl *CPL) ComponentSlice() []*CP {
	return cpl.l
}
