package proton

import "errors"

var (
	// ErrEntityIndexMethodMemberIDUndefined ...
	ErrEntityIndexMethodMemberIDUndefined = errors.New("proton: entity index method member: id undefined")
	// ErrEntityIndexMethodMemberValueUndefined ...
	ErrEntityIndexMethodMemberValueUndefined = errors.New("proton: entity index method member: value undefined")
	// ErrEntityIndexMethodMemberAliasUndefined ...
	ErrEntityIndexMethodMemberAliasUndefined = errors.New("proton: entity index method member: alias undefined")
)

// EntityIndexMethodMember ...
type EntityIndexMethodMember struct {
	id    string
	value string
	alias *Alias
}

// NewEntityIndexMethodMemberAlias ...
func NewEntityIndexMethodMemberAlias(id string, alias *Alias) (*EntityIndexMethodMember, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodMemberIDUndefined
	}
	if alias == nil {
		panic(ErrEntityIndexMethodMemberValueUndefined)
	}
	return &EntityIndexMethodMember{
		id:    id,
		value: "",
		alias: alias}, nil
}

// NewEntityIndexMethodMember ...
func NewEntityIndexMethodMember(id, value string) (*EntityIndexMethodMember, error) {
	if id == "" {
		return nil, ErrEntityIndexMethodMemberIDUndefined
	}
	if value == "" {
		return nil, ErrEntityIndexMethodMemberValueUndefined
	}

	sid := String(id).ToLowerFirst().String()

	return &EntityIndexMethodMember{
		id:    sid,
		value: value,
		alias: nil}, nil
}

// ID ...
func (eimm *EntityIndexMethodMember) ID() String {
	return String(eimm.id)
}

// Value ...
func (eimm *EntityIndexMethodMember) Value() String {
	if eimm.alias != nil {
		return eimm.alias.Value()
	}
	return String(eimm.value)
}

// Alias ...
func (eimm *EntityIndexMethodMember) Alias() *Alias {
	return eimm.alias
}
