package proton

import "errors"

var (
	// ErrComponentMemberListTriedToAddNilMember ...
	ErrComponentMemberListTriedToAddNilMember = errors.New("proton: component member list: tried to add nil component member")
	// ErrComponentMemberListTriedToAddDuplicateMemberID ...
	ErrComponentMemberListTriedToAddDuplicateMemberID = errors.New("proton: component member list: tried to add component member with duplicate id")
)

// ComponentMemberList ...
type ComponentMemberList struct {
	componentMemberSlice []*ComponentMember
}

// NewComponentMemberList ...
func NewComponentMemberList() *ComponentMemberList {
	return &ComponentMemberList{}
}

// AddComponentMember ...
func (cml *ComponentMemberList) AddComponentMember(componentMember *ComponentMember) error {
	if componentMember == nil {
		panic(ErrComponentMemberListTriedToAddNilMember)
	}
	if cml.ComponentMemberWithID(componentMember.ID().String()) != nil {
		return ErrComponentMemberListTriedToAddDuplicateMemberID
	}
	cml.componentMemberSlice = append(cml.componentMemberSlice, componentMember)
	return nil
}

// ComponentMemberWithID ...
func (cml *ComponentMemberList) ComponentMemberWithID(id string) *ComponentMember {
	for _, componentMember := range cml.componentMemberSlice {
		if componentMember.ID().EqualTo(id) {
			return componentMember
		}
	}
	return nil
}

// ComponentMembersWithEntityIndex ...
func (cml *ComponentMemberList) ComponentMembersWithEntityIndex() []*ComponentMember {
	slice := make([]*ComponentMember, 0)
	for _, componentMember := range cml.componentMemberSlice {
		if componentMember.EntityIndexType() > 0 {
			slice = append(slice, componentMember)
		}
	}
	return slice
}

// ComponentMemberSlice ...
func (cml *ComponentMemberList) ComponentMemberSlice() []*ComponentMember {
	return cml.componentMemberSlice
}
