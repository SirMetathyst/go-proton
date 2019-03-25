package proton

import "fmt"

var (
	ErrEntityIndexMethodMemberListTriedToAddNilMember         = fmt.Errorf("entitas(entity index method member list): tried to add nil entity index method member")
	ErrEntityIndexMethodMemberListTriedToAddDuplicateMemberID = fmt.Errorf("entitas(entity index method member list): tried to add entity index method member with duplicate id")
)

// EIMML ...
type EIMML struct {
	l []*EIMM
}

// NewEntityIndexMethodMemberList ...
func NewEntityIndexMethodMemberList() *EIMML {
	return &EIMML{}
}

// AddMember ...
func (eimml *EIMML) AddMember(m *EIMM) error {
	if m == nil {
		return ErrEntityIndexMethodMemberListTriedToAddNilMember
	}
	if eimml.MemberWithID(m.ID().String()) != nil {
		return ErrEntityIndexMethodMemberListTriedToAddDuplicateMemberID
	}
	eimml.l = append(eimml.l, m)
	return nil
}

// MemberWithID ...
func (eimml *EIMML) MemberWithID(id string) *EIMM {
	for _, m := range eimml.l {
		if m.ID().EqualTo(id) {
			return m
		}
	}
	return nil
}

// MemberSlice ...
func (eimml *EIMML) MemberSlice() []*EIMM {
	return eimml.l
}
