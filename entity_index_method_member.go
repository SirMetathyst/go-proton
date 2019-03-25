package proton

import "fmt"

var (
	ErrEntityIndexMethodMemberIDUndefined    = fmt.Errorf("entitas(entity index method member): id undefined")
	ErrEntityIndexMethodMemberValueUndefined = fmt.Errorf("entitas(entity index method member): value undefined")
	ErrEntityIndexMethodMemberAliasUndefined = fmt.Errorf("entitas(entity index method member): alias undefined")
)

// EIMM ...
type EIMM struct {
	id, value string
	alias     *A
}

// NewEntityIndexMethodMemberAlias ...
func NewEntityIndexMethodMemberAlias(id string, alias *A) (*EIMM, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodMemberIDUndefined
	}
	if alias == nil {
		return nil, ErrEntityIndexMethodMemberValueUndefined
	}
	return &EIMM{id, "", alias}, nil
}

// NewEntityIndexMethodMember ...
func NewEntityIndexMethodMember(id, value string) (*EIMM, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodMemberIDUndefined
	}
	if value == "" {
		return nil, ErrEntityIndexMethodMemberValueUndefined
	}

	sid := String(id).ToLowerFirst().String()

	return &EIMM{sid, value, nil}, nil
}

// ID ...
func (eimm *EIMM) ID() String {
	return String(eimm.id)
}

// Value ...
func (eimm *EIMM) Value() String {
	if eimm.alias != nil {
		return eimm.alias.Value()
	}
	return String(eimm.value)
}

// Alias ...
func (eimm *EIMM) Alias() *A {
	return eimm.alias
}
