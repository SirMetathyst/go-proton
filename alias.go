package proton

import (
	"fmt"
	"unicode"
)

var (
	ErrAliasIDUndefined             = fmt.Errorf("entitas(alias): id undefined")
	ErrAliasIDContainsWhitespace    = fmt.Errorf("entitas(alias): id contains whitespace")
	ErrAliasValueUndefined          = fmt.Errorf("entitas(alias): value undefined")
	ErrAliasValueContainsWhitespace = fmt.Errorf("entitas(alias): value contains whitespace")
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
	for _, v := range id {
		if unicode.IsSpace(v) {
			return nil, ErrAliasIDContainsWhitespace
		}
	}
	if value == "" {
		return nil, ErrAliasValueUndefined
	}
	for _, v := range value {
		if unicode.IsSpace(v) {
			return nil, ErrAliasValueContainsWhitespace
		}
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
