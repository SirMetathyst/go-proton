package proton

import "fmt"

var (
	ErrAliasListTriedToAddNilAlias         = fmt.Errorf("entitas(alias list): tried to add nil alias")
	ErrAliasListTriedToAddDuplicateAliasID = fmt.Errorf("entitas(alias list): tried to add alias with duplicate id")
)

// AL ...
type AL struct {
	l []*A
	m map[string]*A
}

// NewAliasList ...
func NewAliasList() *AL {
	return &AL{
		l: make([]*A, 0),
		m: make(map[string]*A),
	}
}

// AddAlias ...
func (al *AL) AddAlias(a *A) error {
	if a == nil {
		return ErrAliasListTriedToAddNilAlias
	}
	if al.AliasWithID(a.ID().String()) != nil {
		return ErrAliasListTriedToAddDuplicateAliasID
	}
	al.l = append(al.l, a)
	al.m[a.ID().String()] = a
	return nil
}

// AliasWithID ...
func (al *AL) AliasWithID(id string) *A {
	v, ok := al.m[id]
	if ok {
		return v
	}
	return nil
}

// AliasSlice ...
func (al *AL) AliasSlice() []*A {
	return al.l
}
