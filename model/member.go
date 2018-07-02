package model

// Member ...
type Member struct {
	id, value   string
	alias       *Alias
	entityIndex int
}

// NewMemberWithAlias ...
func NewMemberWithAlias(id string, alias *Alias, entityIndex int) *Member {
	return &Member{id, "", alias, InRange(entityIndex, 0, 2)}
}

// NewMemberWithValue ...
func NewMemberWithValue(id, value string, entityIndex int) *Member {
	return &Member{id, value, nil, InRange(entityIndex, 0, 2)}
}

// NewMemberWithID ...
func NewMemberWithID(id string) *Member {
	return &Member{id, "", nil, 0}
}

// NewMember ...
func NewMember() *Member {
	return &Member{id: "Undefined"}
}

// GetID ...
func (m *Member) GetID() String {
	return String(m.id)
}

// SetID ...
func (m *Member) SetID(id string) *Member {
	m.id = id
	return m
}

// GetEntityIndex ...
func (m *Member) GetEntityIndex() int {
	return m.entityIndex
}

// GetEntityIndexString ...
func (m *Member) GetEntityIndexString() String {
	switch m.entityIndex {
	case 1:
		return String("PrimaryEntityIndex")
	case 2:
		return String("EntityIndex")
	}
	return String("")
}

// SetEntityIndex ...
func (m *Member) SetEntityIndex(Type int) *Member {
	m.entityIndex = InRange(Type, 0, 2)
	return m
}

// GetValue ...
func (m *Member) GetValue() String {
	if m.alias != nil {
		return m.alias.GetValue()
	}
	return String(m.value)
}

// SetValue ...
func (m *Member) SetValue(value string) *Member {
	m.value = value
	return m
}

// GetAliasValue ...
func (m *Member) GetAliasValue() String {
	if m.alias != nil {
		return m.alias.GetValue()
	}
	return String("")
}

// GetAlias ...
func (m *Member) GetAlias() *Alias {
	return m.alias
}

// SetAlias ...
func (m *Member) SetAlias(alias *Alias) *Member {
	m.alias = alias
	return m
}

// String ...
func (m Member) String() string {
	return "[" + m.GetID().String() + ":\"" + m.GetValue().String() + ":\"" + m.GetEntityIndexString().String() + "\"]"
}
