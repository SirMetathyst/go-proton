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

// AddMember ...
func (cml *ComponentMemberList) AddMember(componentMember *ComponentMember) error {
	if componentMember == nil {
		return ErrComponentMemberListTriedToAddNilMember
	}
	if cml.MemberWithID(componentMember.ID().String()) != nil {
		return ErrComponentMemberListTriedToAddDuplicateMemberID
	}
	cml.componentMemberSlice = append(cml.componentMemberSlice, componentMember)
	return nil
}

// MemberWithID ...
func (cml *ComponentMemberList) MemberWithID(id string) *ComponentMember {
	for _, componentMember := range cml.componentMemberSlice {
		if componentMember.ID().EqualTo(id) {
			return componentMember
		}
	}
	return nil
}

// MembersWithEntityIndex ...
func (cml *ComponentMemberList) MembersWithEntityIndex() []*ComponentMember {
	slice := make([]*ComponentMember, 0)
	for _, componentMember := range cml.componentMemberSlice {
		if componentMember.EntityIndexType() > 0 {
			slice = append(slice, componentMember)
		}
	}
	return slice
}

// MemberSlice ...
func (cml *ComponentMemberList) MemberSlice() []*ComponentMember {
	return cml.componentMemberSlice
}
