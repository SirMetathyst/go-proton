package proton

import (
	"errors"
)

var (
	// ErrComponentMemberIDUndefined ...
	ErrComponentMemberIDUndefined = errors.New("proton: component member: id undefined")
	// ErrComponentMemberIDContainsWhitespace ...
	ErrComponentMemberIDContainsWhitespace = errors.New("proton: component member: id contains whitespace")
	// ErrComponentMemberValueUndefined ...
	ErrComponentMemberValueUndefined = errors.New("proton: component member: value undefined")
	// ErrComponentMemberValueContainsWhitespace ...
	ErrComponentMemberValueContainsWhitespace = errors.New("proton: component member: value contains whitespace")
	// ErrComponentMemberAliasShouldNotBeNil ...
	ErrComponentMemberAliasShouldNotBeNil = errors.New("proton: component member: alias should not be nil")
	// ErrComponentMemberEntityIndexInvalid ...
	ErrComponentMemberEntityIndexInvalid = errors.New("proton: component member: entity index is invalid")
)

// ComponentMember ...
type ComponentMember struct {
	id              string
	value           string
	alias           *Alias
	entityIndexType EntityIndexType
}

// NewComponentMemberAlias ...
func NewComponentMemberAlias(
	id string,
	alias *Alias,
	entityIndexType EntityIndexType) (*ComponentMember, error) {

	if id == "" {
		return nil, ErrComponentMemberIDUndefined
	}
	if containsWhitespace(id) {
		return nil, ErrComponentMemberIDContainsWhitespace
	}
	if alias == nil {
		return nil, ErrComponentMemberAliasShouldNotBeNil
	}
	if !entityIndexType.IsValid() {
		return nil, ErrComponentMemberEntityIndexInvalid
	}
	return &ComponentMember{
		id:              id,
		value:           "",
		alias:           alias,
		entityIndexType: entityIndexType}, nil
}

// NewComponentMember ...
func NewComponentMember(
	id string,
	value string,
	entityIndexType EntityIndexType) (*ComponentMember, error) {

	if id == "" {
		return nil, ErrComponentMemberIDUndefined
	}
	if containsWhitespace(id) {
		return nil, ErrComponentMemberIDContainsWhitespace
	}
	if value == "" {
		return nil, ErrComponentMemberValueUndefined
	}
	if containsWhitespace(value) {
		return nil, ErrComponentMemberValueContainsWhitespace
	}
	if !entityIndexType.IsValid() {
		return nil, ErrComponentMemberEntityIndexInvalid
	}

	return &ComponentMember{
		id:              id,
		value:           value,
		alias:           nil,
		entityIndexType: entityIndexType}, nil
}

// ID ...
func (cm *ComponentMember) ID() String {
	return String(cm.id)
}

// EntityIndexType ...
func (cm *ComponentMember) EntityIndexType() EntityIndexType {
	return cm.entityIndexType
}

// Value ...
func (cm *ComponentMember) Value() String {
	if cm.alias != nil {
		return cm.alias.Value()
	}
	return String(cm.value)
}

// Alias ...
func (cm *ComponentMember) Alias() *Alias {
	return cm.alias
}
