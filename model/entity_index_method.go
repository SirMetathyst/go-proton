package model

import (
	"bytes"
)

// EntityIndexMethod ...
type EntityIndexMethod struct {
	id     string
	member map[string]*Member
}

// NewEntityIndexMethodWithID ...
func NewEntityIndexMethodWithID(id string) *EntityIndexMethod {
	return &EntityIndexMethod{
		id:     id,
		member: make(map[string]*Member, 0)}
}

// NewEntityIndexMethod ...
func NewEntityIndexMethod() *EntityIndexMethod {
	return NewEntityIndexMethodWithID("Untitled")
}

// GetID ...
func (m *EntityIndexMethod) GetID() String {
	return String(m.id)
}

// SetID ...
func (m *EntityIndexMethod) SetID(id string) *EntityIndexMethod {
	m.id = id
	return m
}

// GetMemberWithID ...
func (m *EntityIndexMethod) GetMemberWithID(id string) *Member {
	mbr, _ := m.member[id]
	return mbr
}

// GetMember ...
func (m *EntityIndexMethod) GetMember() []*Member {
	slice := make([]*Member, 0)
	for _, mbr := range m.member {
		slice = append(slice, mbr)
	}
	return slice
}

// AddMember ...
func (m *EntityIndexMethod) AddMember(overwrite bool, member ...*Member) *EntityIndexMethod {
	for _, mbr := range member {
		_, exist := m.member[mbr.GetID().String()]
		if overwrite || !overwrite && !exist {
			m.member[mbr.GetID().String()] = mbr
		}
	}
	return m
}

// CreateMember ...
func (m *EntityIndexMethod) CreateMember(overwrite bool) *Member {
	mbr := NewMember()
	m.AddMember(overwrite, mbr)
	return mbr
}

// String ...
func (m EntityIndexMethod) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(m.GetID()))
	buffer.WriteRune('(')

	buffer.WriteString(") ")
	for _, Member := range m.GetMember() {
		buffer.WriteString("Member{")
		buffer.WriteString(Member.String())
		buffer.WriteRune('}')
	}
	return buffer.String()
}
