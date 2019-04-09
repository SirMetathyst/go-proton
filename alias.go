package proton

import (
	"errors"
)

var (
	// ErrAliasIDUndefined ...
	ErrAliasIDUndefined = errors.New("proton: alias: id undefined")
	// ErrAliasIDContainsWhitespace ...
	ErrAliasIDContainsWhitespace = errors.New("proton: alias: id contains whitespace")
	// ErrAliasValueUndefined ...
	ErrAliasValueUndefined = errors.New("proton: alias: value undefined")
	// ErrAliasValueContainsWhitespace ...
	ErrAliasValueContainsWhitespace = errors.New("proton: alias: value contains whitespace")
)

// Alias ...
type Alias struct {
	id    string
	value string
}

// NewAlias ...
func NewAlias(id, value string) (*Alias, error) {
	if id == "" {
		return nil, ErrAliasIDUndefined
	}
	if containsWhitespace(id) {
		return nil, ErrAliasIDContainsWhitespace
	}
	if value == "" {
		return nil, ErrAliasValueUndefined
	}
	if containsWhitespace(value) {
		return nil, ErrAliasValueContainsWhitespace
	}
	return &Alias{id, value}, nil
}

// ID ...
func (a *Alias) ID() String {
	return String(a.id)
}

// Value ...
func (a *Alias) Value() String {
	return String(a.value)
}
