package model

import "fmt"

var (
	ErrAliasListTriedToAddNilAlias         = fmt.Errorf("AliasList: Tried to add nil `Alias`")
	ErrAliasListTriedToAddDuplicateAliasID = fmt.Errorf("AliasList: Tried to add `Alias` with duplicate `ID`")
)

// AL ...
type AL struct {
	l []*A
}

// NewAliasList ...
func NewAliasList() AL {
	return AL{}
}

// indexOf ...
func (al *AL) indexOf(a *A) int {
	for i, ca := range al.l {
		if ca.ID() == a.ID() {
			return i
		}
	}
	return -1
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
	return nil
}

// AddReplaceAlias ...
func (al *AL) AddReplaceAlias(a *A) error {
	i := al.indexOf(a)
	if i != -1 {
		al.l[i] = a
	} else {
		return al.AddAlias(a)
	}
	return nil
}

// AliasWithID ...
func (al *AL) AliasWithID(id string) *A {
	for _, a := range al.l {
		if a.ID().EqualTo(id) {
			return a
		}
	}
	return nil
}

// AliasList ...
func (al *AL) AliasList() []*A {
	return al.l
}
