package proton

import "errors"

var (
	// ErrEntityIndexMethodBuilderAliasListShouldNotBeNil ...
	ErrEntityIndexMethodBuilderAliasListShouldNotBeNil = errors.New("proton: entity index method builder: alias list should not be nil")
	// ErrEntityIndexMethodBuilderEntityIndexMethodListShouldNotBeNil ...
	ErrEntityIndexMethodBuilderEntityIndexMethodListShouldNotBeNil = errors.New("proton: entity index method builder: entity index method list should not be nil")
	// ErrEntityIndexMethodBuilderEntityIndexMethodAlreadyBuilt ...
	ErrEntityIndexMethodBuilderEntityIndexMethodAlreadyBuilt = errors.New("proton: entity index method builder: entity index method is already built")
)

// EntityIndexMethodBuilder ...
type EntityIndexMethodBuilder struct {
	aliasList                         *AliasList
	targetEntityIndexMethodMemberList *EntityIndexMethodMemberList
	entityIndexMethodList             *EntityIndexMethodList
	id                                string
	built                             bool
}

// NewEntityIndexMethodBuilder ...
func NewEntityIndexMethodBuilder(
	aliasList *AliasList,
	entityIndexMethodList *EntityIndexMethodList) *EntityIndexMethodBuilder {
	if aliasList == nil {
		panic(ErrEntityIndexMethodBuilderAliasListShouldNotBeNil)
	}
	if entityIndexMethodList == nil {
		panic(ErrEntityIndexMethodBuilderEntityIndexMethodListShouldNotBeNil)
	}
	return &EntityIndexMethodBuilder{
		aliasList:                         aliasList,
		entityIndexMethodList:             entityIndexMethodList,
		targetEntityIndexMethodMemberList: NewEntityIndexMethodMemberList(),
	}
}

// SetID ...
func (eimb *EntityIndexMethodBuilder) SetID(id string) *EntityIndexMethodBuilder {
	eimb.id = id
	return eimb
}

// NewMember ...
func (eimb *EntityIndexMethodBuilder) NewMember() *EntityIndexMethodMemberBuilder {
	return NewEntityIndexMethodMemberBuilder(eimb.aliasList, eimb.targetEntityIndexMethodMemberList)
}

// Build ...
func (eimb *EntityIndexMethodBuilder) Build() error {
	if eimb.built {
		return ErrEntityIndexMethodBuilderEntityIndexMethodAlreadyBuilt
	}
	entityIndexMethod, err := NewEntityIndexMethod(eimb.id, eimb.targetEntityIndexMethodMemberList)
	if err != nil {
		return err
	}
	err = eimb.entityIndexMethodList.AddEntityIndexMethod(entityIndexMethod)
	if err != nil {
		return err
	}
	eimb.built = true
	return nil
}
