package proton

import "fmt"

var (
	ErrComponentMemberIDUndefined        = fmt.Errorf("entitas(component member): id undefined")
	ErrComponentMemberValueUndefined     = fmt.Errorf("entitas(component member): value undefined")
	ErrComponentMemberAliasUndefined     = fmt.Errorf("entitas(component member): alias undefined")
	ErrComponentMemberEntityIndexInvalid = fmt.Errorf("entitas(component member): entity index is invalid")
)

// CM ...
type CM struct {
	id, value   string
	alias       *A
	entityIndex EntityIndex
}

// NewComponentMemberAlias ...
func NewComponentMemberAlias(id string, alias *A, entityIndex EntityIndex) (*CM, error) {
	if id == "" {
		return nil, ErrComponentMemberIDUndefined
	}
	if alias == nil {
		return nil, ErrComponentMemberValueUndefined
	}
	if !entityIndex.IsValid() {
		return nil, ErrComponentMemberEntityIndexInvalid
	}
	return &CM{id, "", alias, entityIndex}, nil
}

// NewComponentMember ...
func NewComponentMember(id, value string, entityIndex EntityIndex) (*CM, error) {
	if id == "" {
		return nil, ErrComponentMemberIDUndefined
	}
	if value == "" {
		return nil, ErrComponentMemberValueUndefined
	}
	if !entityIndex.IsValid() {
		return nil, ErrComponentMemberEntityIndexInvalid
	}

	sid := String(id).ToLowerFirst().String()

	return &CM{sid, value, nil, entityIndex}, nil
}

// ID ...
func (cm *CM) ID() String {
	return String(cm.id)
}

// EntityIndex ...
func (cm *CM) EntityIndex() EntityIndex {
	return cm.entityIndex
}

// Value ...
func (cm *CM) Value() String {
	if cm.alias != nil {
		return cm.alias.Value()
	}
	return String(cm.value)
}

// Alias ...
func (cm *CM) Alias() *A {
	return cm.alias
}
