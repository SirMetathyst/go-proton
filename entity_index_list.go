package proton

import "fmt"

var (
	ErrEntityIndexListTriedToAddNilEntityIndex         = fmt.Errorf("EntityIndexList: Tried to add nil `EntityIndex`")
	ErrEntityIndexListTriedToAddDuplicateEntityIndexID = fmt.Errorf("EntityIndexList: Tried to add `EntityIndex` with duplicate `ID`")
)

// EIL ...
type EIL struct {
	l []*EI
}

// NewEntityIndexList ...
func NewEntityIndexList() *EIL {
	return &EIL{}
}

// AddEntityIndex ...
func (eil *EIL) AddEntityIndex(ei *EI) error {
	if ei == nil {
		return ErrEntityIndexListTriedToAddNilEntityIndex
	}
	if eil.EntityIndexWithID(ei.ID().String()) != nil {
		return ErrEntityIndexListTriedToAddDuplicateEntityIndexID
	}
	eil.l = append(eil.l, ei)
	return nil
}

// EntityIndexWithID ...
func (eil *EIL) EntityIndexWithID(id string) *EI {
	for _, ei := range eil.l {
		if ei.ID().EqualTo(id) {
			return ei
		}
	}
	return nil
}

// EntityIndexList ...
func (eil *EIL) EntityIndexList() []*EI {
	return eil.l
}
