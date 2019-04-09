package proton

import "errors"

var (
	// ErrEntityIndexBuilderContextListShouldNotBeNil ...
	ErrEntityIndexBuilderContextListShouldNotBeNil = errors.New("proton: entity index builder: context list should not be nil")
	// ErrEntityIndexBuilderAliasListShouldNotBeNil ...
	ErrEntityIndexBuilderAliasListShouldNotBeNil = errors.New("proton: entity index builder: alias list should not be nil")
	// ErrEntityIndexBuilderEntityIndexListShouldNotBeNil ...
	ErrEntityIndexBuilderEntityIndexListShouldNotBeNil = errors.New("proton: entity index builder: entity index list should not be nil")
	// ErrEntityIndexBuilderEntityIndexAlreadyBuilt ...
	ErrEntityIndexBuilderEntityIndexAlreadyBuilt = errors.New("proton: entity index builder: entity index is already built")
	// ErrEntityIndexBuilderEntityIndexMustHaveContext ...
	ErrEntityIndexBuilderEntityIndexMustHaveContext = errors.New("proton: entity index builder: entity index does not have a context")
	// ErrEntityIndexBuilderEntityIndexMustHaveMethod ...
	ErrEntityIndexBuilderEntityIndexMustHaveMethod = errors.New("proton: entity index builder: entity index does not have a method")
)

// EntityIndexBuilder ...
type EntityIndexBuilder struct {
	contextList                 *ContextList
	aliasList                   *AliasList
	entityIndexList             *EntityIndexList
	context                     *Context
	targetEntityIndexMethodList *EntityIndexMethodList
	id                          string
	isPrimary                   bool
	built                       bool
}

// NewEntityIndexBuilder ...
func NewEntityIndexBuilder(
	contextList *ContextList,
	aliasList *AliasList,
	entityIndexList *EntityIndexList) *EntityIndexBuilder {
	if contextList == nil {
		panic(ErrEntityIndexBuilderContextListShouldNotBeNil)
	}
	if aliasList == nil {
		panic(ErrEntityIndexBuilderAliasListShouldNotBeNil)
	}
	if entityIndexList == nil {
		panic(ErrEntityIndexBuilderEntityIndexListShouldNotBeNil)
	}
	return &EntityIndexBuilder{
		contextList:                 contextList,
		aliasList:                   aliasList,
		entityIndexList:             entityIndexList,
		targetEntityIndexMethodList: NewEntityIndexMethodList(),
	}
}

// SetID ...
func (eib *EntityIndexBuilder) SetID(id string) *EntityIndexBuilder {
	eib.id = id
	return eib
}

// SetPrimary ...
func (eib *EntityIndexBuilder) SetPrimary(isPrimary bool) *EntityIndexBuilder {
	eib.isPrimary = isPrimary
	return eib
}

// AddContext ...
func (eib *EntityIndexBuilder) AddContext(id string) *EntityIndexBuilder {
	context := eib.contextList.ContextWithID(id)
	eib.context = context
	return eib
}

// NewMethod ...
func (eib *EntityIndexBuilder) NewMethod() *EntityIndexMethodBuilder {
	return NewEntityIndexMethodBuilder(eib.aliasList, eib.targetEntityIndexMethodList)
}

// Build ...
func (eib *EntityIndexBuilder) Build() error {
	if eib.built {
		return ErrEntityIndexBuilderEntityIndexAlreadyBuilt
	}
	if eib.context == nil {
		return ErrEntityIndexBuilderEntityIndexMustHaveContext
	}
	if len(eib.targetEntityIndexMethodList.EntityIndexMethodSlice()) == 0 {
		return ErrEntityIndexBuilderEntityIndexMustHaveMethod
	}
	entityIndex, err := NewEntityIndex(eib.id, eib.isPrimary, eib.context, eib.targetEntityIndexMethodList)
	if err != nil {
		return err
	}
	err = eib.entityIndexList.AddEntityIndex(entityIndex)
	if err != nil {
		return err
	}
	eib.built = true
	return nil
}
