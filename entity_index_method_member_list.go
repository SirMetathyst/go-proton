package proton

import "errors"

var (
	// ErrEntityIndexMethodMemberListTriedToAddNilMember ...
	ErrEntityIndexMethodMemberListTriedToAddNilMember = errors.New("proton: entity index method member list: tried to add nil entity index method member")
	// ErrEntityIndexMethodMemberListTriedToAddDuplicateMemberID ...
	ErrEntityIndexMethodMemberListTriedToAddDuplicateMemberID = errors.New("proton: entity index method member list: tried to add entity index method member with duplicate id")
)

// EntityIndexMethodMemberList ...
type EntityIndexMethodMemberList struct {
	entityIndexMethodMemberSlice []*EntityIndexMethodMember
}

// NewEntityIndexMethodMemberList ...
func NewEntityIndexMethodMemberList() *EntityIndexMethodMemberList {
	return &EntityIndexMethodMemberList{}
}

// AddMember ...
func (eimml *EntityIndexMethodMemberList) AddMember(entityIndexMethodMember *EntityIndexMethodMember) error {
	if entityIndexMethodMember == nil {
		panic(ErrEntityIndexMethodMemberListTriedToAddNilMember)
	}
	if eimml.MemberWithID(entityIndexMethodMember.ID().String()) != nil {
		return ErrEntityIndexMethodMemberListTriedToAddDuplicateMemberID
	}
	eimml.entityIndexMethodMemberSlice = append(eimml.entityIndexMethodMemberSlice, entityIndexMethodMember)
	return nil
}

// MemberWithID ...
func (eimml *EntityIndexMethodMemberList) MemberWithID(id string) *EntityIndexMethodMember {
	for _, entityIndexMethodMember := range eimml.entityIndexMethodMemberSlice {
		if entityIndexMethodMember.ID().EqualTo(id) {
			return entityIndexMethodMember
		}
	}
	return nil
}

// MemberSlice ...
func (eimml *EntityIndexMethodMemberList) MemberSlice() []*EntityIndexMethodMember {
	return eimml.entityIndexMethodMemberSlice
}
