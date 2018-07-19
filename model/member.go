package model

import "fmt"

var (
	ErrMemberIDUndefined        = fmt.Errorf("Member: `ID` Undefined.")
	ErrMemberValueUndefined     = fmt.Errorf("Member: `Value` Undefined.")
	ErrMemberAliasUndefined     = fmt.Errorf("Member: `Alias` Undefined.")
	ErrMemberEntityIndexInvalid = fmt.Errorf("Member: `EntityIndex` is Invalid.")
)

// M ...
type M struct {
	id, value   string
	alias       *A
	entityIndex EntityIndex
}

// NewMemberAlias ...
func NewMemberAlias(id string, alias *A, entityIndex EntityIndex) (*M, error) {
	if id == "" {
		return nil, ErrMemberIDUndefined
	}
	if alias == nil {
		return nil, ErrMemberAliasUndefined
	}
	if !entityIndex.IsValid() {
		return nil, ErrMemberEntityIndexInvalid
	}
	return &M{id, "", alias, entityIndex}, nil
}

// NewMember ...
func NewMember(id, value string, entityIndex EntityIndex) (*M, error) {
	if id == "" {
		return nil, ErrMemberIDUndefined
	}
	if value == "" {
		return nil, ErrAliasValueUndefined
	}
	if !entityIndex.IsValid() {
		return nil, ErrMemberEntityIndexInvalid
	}

	sid := String(id).ToLowerFirst().String()

	return &M{sid, value, nil, entityIndex}, nil
}

// ID ...
func (m *M) ID() String {
	return String(m.id)
}

// EntityIndex ...
func (m *M) EntityIndex() EntityIndex {
	return m.entityIndex
}

// Value ...
func (m *M) Value() String {
	if m.alias != nil {
		return m.alias.Value()
	}
	return String(m.value)
}

// Alias ...
func (m *M) Alias() *A {
	return m.alias
}
