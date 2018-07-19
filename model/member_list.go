package model

import "fmt"

var (
	ErrMemberListTriedToAddNilMember         = fmt.Errorf("MemberList: Tried to add nil `Member`")
	ErrMemberListTriedToAddDuplicateMemberID = fmt.Errorf("MemberList: Tried to add `Member` with duplicate `ID`")
)

// ML ...
type ML struct {
	l []*M
}

// NewMemberList ...
func NewMemberList() ML {
	return ML{}
}

// AddMember ...
func (ml *ML) AddMember(m *M) error {
	if m == nil {
		return ErrMemberListTriedToAddNilMember
	}
	if ml.MemberWithID(m.ID().String()) != nil {
		return ErrMemberListTriedToAddDuplicateMemberID
	}
	ml.l = append(ml.l, m)
	return nil
}

// MemberWithID ...
func (ml *ML) MemberWithID(id string) *M {
	for _, m := range ml.l {
		if m.ID().EqualTo(id) {
			return m
		}
	}
	return nil
}

// MembersWithEntityIndex ...
func (ml *ML) MembersWithEntityIndex() []*M {
	slice := make([]*M, 0)
	for _, m := range ml.l {
		if m.EntityIndex() > 0 {
			slice = append(slice, m)
		}
	}
	return slice
}

// MemberList ...
func (ml *ML) MemberList() []*M {
	return ml.l
}
