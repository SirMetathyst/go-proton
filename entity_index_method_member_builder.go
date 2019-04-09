package proton

import "errors"

var (
	// ErrEntityIndexMethodMemberBuilderMemberAlreadyBuilt ...
	ErrEntityIndexMethodMemberBuilderMemberAlreadyBuilt = errors.New("proton: entity index method member builder: entity index method member is already built")
	// ErrEntityIndexMethodMemberBuilderAliasShouldNotBeNil ...
	ErrEntityIndexMethodMemberBuilderAliasShouldNotBeNil = errors.New("proton: entity index method member builder: alias list should not be nil")
	// ErrEntityIndexMethodMemberBuilderMemberShouldNotBeNil ...
	ErrEntityIndexMethodMemberBuilderMemberShouldNotBeNil = errors.New("proton: entity index method member builder: entity index method member list should not be nil")
)

// EntityIndexMethodMemberBuilder ...
type EntityIndexMethodMemberBuilder struct {
	aliasList                   *AliasList
	entityIndexMethodMemberList *EntityIndexMethodMemberList
	alias                       *Alias
	built                       bool
	id                          string
	value                       string
}

// NewEntityIndexMethodMemberBuilder ...
func NewEntityIndexMethodMemberBuilder(
	aliasList *AliasList,
	entityIndexMethodMemberList *EntityIndexMethodMemberList) *EntityIndexMethodMemberBuilder {
	if aliasList == nil {
		panic(ErrEntityIndexMethodMemberBuilderAliasShouldNotBeNil)
	}
	if entityIndexMethodMemberList == nil {
		panic(ErrEntityIndexMethodMemberBuilderMemberShouldNotBeNil)
	}
	return &EntityIndexMethodMemberBuilder{
		aliasList:                   aliasList,
		entityIndexMethodMemberList: entityIndexMethodMemberList}
}

// SetID ...
func (eimmb *EntityIndexMethodMemberBuilder) SetID(id string) *EntityIndexMethodMemberBuilder {
	eimmb.id = id
	return eimmb
}

// SetValue ...
func (eimmb *EntityIndexMethodMemberBuilder) SetValue(value string) *EntityIndexMethodMemberBuilder {
	eimmb.value = value
	return eimmb
}

// SetAlias ...
func (eimmb *EntityIndexMethodMemberBuilder) SetAlias(id string) *EntityIndexMethodMemberBuilder {
	eimmb.alias = eimmb.aliasList.AliasWithID(id)
	return eimmb
}

// Build ...
func (eimmb *EntityIndexMethodMemberBuilder) Build() error {
	if eimmb.built {
		return ErrEntityIndexMethodMemberBuilderMemberAlreadyBuilt
	}
	if eimmb.alias != nil {
		entityIndexMethodMember, err := NewEntityIndexMethodMemberAlias(eimmb.id, eimmb.alias)
		if err != nil {
			return err
		}
		eimmb.entityIndexMethodMemberList.AddMember(entityIndexMethodMember)
		return nil
	}
	entityIndexMethodMember, err := NewEntityIndexMethodMember(eimmb.id, eimmb.value)
	if err != nil {
		return err
	}
	err = eimmb.entityIndexMethodMemberList.AddMember(entityIndexMethodMember)
	if err != nil {
		return err
	}
	eimmb.built = true
	return nil
}
