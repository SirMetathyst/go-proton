package proton

import "errors"

var (
	// ErrComponentMemberBuilderMemberAlreadyBuilt ...
	ErrComponentMemberBuilderMemberAlreadyBuilt = errors.New("proton: component member builder: component member is already built")
	// ErrComponentMemberBuilderAliasListShouldNotBeNil ...
	ErrComponentMemberBuilderAliasListShouldNotBeNil = errors.New("proton: component member builder: alias list should not be nil")
	// ErrComponentMemberBuilderMemberShouldNotBeNil ...
	ErrComponentMemberBuilderMemberShouldNotBeNil = errors.New("proton: component member builder: component member list should not be nil")
	// ErrComponentMemberBuilderAliasNotFound ...
	ErrComponentMemberBuilderAliasNotFound = errors.New("proton: component member builder: the given alias id was not found")
)

// ComponentMemberBuilder ...
type ComponentMemberBuilder struct {
	aliasList           *AliasList
	componentMemberList *ComponentMemberList
	useAlias            bool
	alias               *Alias
	built               bool
	id                  string
	value               string
	entityIndexType     EntityIndexType
}

// NewComponentMemberBuilder ...
func NewComponentMemberBuilder(
	aliasList *AliasList,
	componentMemberList *ComponentMemberList) *ComponentMemberBuilder {
	if aliasList == nil {
		panic(ErrComponentMemberBuilderAliasListShouldNotBeNil)
	}
	if componentMemberList == nil {
		panic(ErrComponentMemberBuilderMemberShouldNotBeNil)
	}
	return &ComponentMemberBuilder{
		aliasList:           aliasList,
		componentMemberList: componentMemberList,
		useAlias:            false}
}

// SetID ...
func (cpmb *ComponentMemberBuilder) SetID(id string) *ComponentMemberBuilder {
	cpmb.id = id
	return cpmb
}

// SetValue ...
func (cpmb *ComponentMemberBuilder) SetValue(value string) *ComponentMemberBuilder {
	cpmb.value = value
	return cpmb
}

// SetEntityIndexType ...
func (cpmb *ComponentMemberBuilder) SetEntityIndexType(value EntityIndexType) *ComponentMemberBuilder {
	cpmb.entityIndexType = value
	return cpmb
}

// SetAlias ...
func (cpmb *ComponentMemberBuilder) SetAlias(id string) *ComponentMemberBuilder {
	cpmb.alias = cpmb.aliasList.AliasWithID(id)
	cpmb.useAlias = true
	return cpmb
}

// Build ...
func (cpmb *ComponentMemberBuilder) Build() error {
	if cpmb.built {
		return ErrComponentMemberBuilderMemberAlreadyBuilt
	}
	if cpmb.useAlias {
		componentMember, err := NewComponentMemberAlias(cpmb.id, cpmb.alias, cpmb.entityIndexType)
		if err != nil {
			if err == ErrComponentMemberAliasShouldNotBeNil {
				return ErrComponentMemberBuilderAliasNotFound
			}
			return err
		}
		cpmb.componentMemberList.AddComponentMember(componentMember)
		cpmb.built = true
		return nil
	}
	componentMember, err := NewComponentMember(cpmb.id, cpmb.value, cpmb.entityIndexType)
	if err != nil {
		return err
	}
	err = cpmb.componentMemberList.AddComponentMember(componentMember)
	if err != nil {
		return err
	}
	cpmb.built = true
	return nil
}
