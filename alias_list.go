package proton

import "errors"

var (
	// ErrAliasListTriedToAddNilAlias ...
	ErrAliasListTriedToAddNilAlias = errors.New("proton: alias list: tried to add nil alias")
	// ErrAliasListTriedToAddDuplicateAliasID ...
	ErrAliasListTriedToAddDuplicateAliasID = errors.New("proton: alias list: tried to add alias with duplicate id")
)

// AliasList ...
type AliasList struct {
	aliasSlice []*Alias
	aliasMap   map[string]*Alias
}

// NewAliasList ...
func NewAliasList() *AliasList {
	return &AliasList{make([]*Alias, 0), make(map[string]*Alias)}
}

// AddAlias ...
func (al *AliasList) AddAlias(alias *Alias) error {
	if alias == nil {
		panic(ErrAliasListTriedToAddNilAlias)
	}
	if al.AliasWithID(alias.ID().String()) != nil {
		return ErrAliasListTriedToAddDuplicateAliasID
	}
	al.aliasSlice = append(al.aliasSlice, alias)
	al.aliasMap[alias.ID().String()] = alias
	return nil
}

// AliasWithID ...
func (al *AliasList) AliasWithID(id string) *Alias {
	v, ok := al.aliasMap[id]
	if ok {
		return v
	}
	return nil
}

// AliasSlice ...
func (al *AliasList) AliasSlice() []*Alias {
	return al.aliasSlice
}
