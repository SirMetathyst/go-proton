package model

import "fmt"

var (
	ErrAliasIDUndefined    = fmt.Errorf("Alias: `ID` Undefined.")
	ErrAliasValueUndefined = fmt.Errorf("Alias: `Value` Undefined.")
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
	if value == "" {
		return nil, ErrAliasValueUndefined
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

// String ...
func (a A) String() string {
	return "[" + a.ID().String() + ":\"" + a.Value().String() + "\"]"
}
