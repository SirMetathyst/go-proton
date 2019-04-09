package proton

import "errors"

var (
	// ErrEntityIndexMethodIDUndefined ...
	ErrEntityIndexMethodIDUndefined = errors.New("proton: entity index method: id undefined")
)

// EntityIndexMethod ...
type EntityIndexMethod struct {
	id                          string
	entityIndexMethodMemberList *EntityIndexMethodMemberList
}

// NewEntityIndexMethod ...
func NewEntityIndexMethod(
	id string,
	entityIndexMethodMemberList *EntityIndexMethodMemberList) (*EntityIndexMethod, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodIDUndefined
	}
	return &EntityIndexMethod{
		id:                          id,
		entityIndexMethodMemberList: entityIndexMethodMemberList}, nil
}

// ID ...
func (eim *EntityIndexMethod) ID() String {
	return String(eim.id)
}

// MemberSlice ...
func (eim *EntityIndexMethod) MemberSlice() []*EntityIndexMethodMember {
	return eim.entityIndexMethodMemberList.MemberSlice()
}
