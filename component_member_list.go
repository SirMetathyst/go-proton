package proton

import "fmt"

var (
	ErrComponentMemberListTriedToAddNilMember         = fmt.Errorf("entitas(component member list): tried to add nil component member")
	ErrComponentMemberListTriedToAddDuplicateMemberID = fmt.Errorf("entitas(component member list): tried to add component member with duplicate id")
)

// CML ...
type CML struct {
	l []*CM
}

// NewComponentMemberList ...
func NewComponentMemberList() *CML {
	return &CML{}
}

// AddMember ...
func (cml *CML) AddMember(m *CM) error {
	if m == nil {
		return ErrComponentMemberListTriedToAddNilMember
	}
	if cml.MemberWithID(m.ID().String()) != nil {
		return ErrComponentMemberListTriedToAddDuplicateMemberID
	}
	cml.l = append(cml.l, m)
	return nil
}

// MemberWithID ...
func (cml *CML) MemberWithID(id string) *CM {
	for _, m := range cml.l {
		if m.ID().EqualTo(id) {
			return m
		}
	}
	return nil
}

// MembersWithEntityIndex ...
func (cml *CML) MembersWithEntityIndex() []*CM {
	slice := make([]*CM, 0)
	for _, m := range cml.l {
		if m.EntityIndex() > 0 {
			slice = append(slice, m)
		}
	}
	return slice
}

// MemberSlice ...
func (cml *CML) MemberSlice() []*CM {
	return cml.l
}
