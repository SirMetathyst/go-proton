package proton

import "fmt"

var (
	ErrEntityIndexMethodIDUndefined = fmt.Errorf("EntityIndexMethod: `ID` Undefined.")
)

// EIM ...
type EIM struct {
	id    string
	eimml *EIMML
}

// NewEntityIndexMethod ...
func NewEntityIndexMethod(id string, eimml *EIMML) (*EIM, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodIDUndefined
	}
	return &EIM{id, eimml}, nil
}

// ID ...
func (eim *EIM) ID() String {
	return String(eim.id)
}

// MemberSlice ...
func (eim *EIM) MemberSlice() []*EIMM {
	return eim.eimml.MemberSlice()
}
