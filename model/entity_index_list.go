package model

import "fmt"

// EIL ...
type EIL struct {
	l []*EI
}

// NewEntityIndexList ...
func NewEntityIndexList() EIL {
	return EIL{}
}

// AddEntityIndex ...
func (eil *EIL) AddEntityIndex(ei *EI) {
	eil.l = append(eil.l, ei)
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

// IsUnique ...
func (eil *EIL) IsUnique() (String, error) {
	lookup := make(map[String]bool, 0)
	for _, ei := range eil.l {
		if _, ok := lookup[ei.ID()]; ok {
			return ei.ID(), fmt.Errorf("IsUnique: Found duplicate of `%s`", ei.ID())
		}
		lookup[ei.ID()] = true
	}
	return "", nil
}
