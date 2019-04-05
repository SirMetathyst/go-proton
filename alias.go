package proton

import (
	"errors"
)

var (
	ErrAliasIDUndefined             = errors.New("proton: alias: id undefined")
	ErrAliasIDContainsWhitespace    = errors.New("proton: alias: id contains whitespace")
	ErrAliasValueUndefined          = errors.New("proton: alias: value undefined")
	ErrAliasValueContainsWhitespace = errors.New("proton: alias: value contains whitespace")
)

// A ...
type A struct {
	id, value string
}

// NewAlias ...
func NewAlias(id, value string) (*A, error) {
	if id == "" {
		return nil, ErrAliasIDUndefined
	}
	if ContainsWhitespace(id) {
		return nil, ErrAliasIDContainsWhitespace
	}
	if value == "" {
		return nil, ErrAliasValueUndefined
	}
	if ContainsWhitespace(value) {
		return nil, ErrAliasValueContainsWhitespace
	}
	return &A{id, value}, nil
}

// ID ...
func (a *A) ID() String {
	return String(a.id)
}

// Value ...
func (a *A) Value() String {
	return String(a.value)
}
