package proton

import "errors"

var (
	// ErrComponentBuilderContextListShouldNotBeNil ...
	ErrComponentBuilderContextListShouldNotBeNil = errors.New("proton: component builder: context list should not be nil")
	// ErrComponentBuilderAliasListShouldNotBeNil ...
	ErrComponentBuilderAliasListShouldNotBeNil = errors.New("proton: component builder: alias list should not be nil")
	// ErrComponentBuilderComponentListShouldNotBeNil ...
	ErrComponentBuilderComponentListShouldNotBeNil = errors.New("proton: component builder: component list should not be nil")
	// ErrComponentBuilderComponentAlreadyBuilt ...
	ErrComponentBuilderComponentAlreadyBuilt = errors.New("proton: component builder: component is already built")
	// ErrComponentBuilderComponentMustHaveContext ...
	ErrComponentBuilderComponentMustHaveContext = errors.New("proton: component builder: component does not have a context")
)

// ComponentBuilder ...
type ComponentBuilder struct {
	componentList             *ComponentList
	targetContextList         *ContextList
	targetComponentMemberList *ComponentMemberList
	contextList               *ContextList
	aliasList                 *AliasList
	built                     bool
	defaultContext            string
	id                        string
	flagPrefix                string
	unique                    bool
	eventTarget               EventTarget
	eventType                 EventType
	eventPriority             int
	cleanupMode               CleanupMode
}

// NewComponentBuilder ...
func NewComponentBuilder(
	contextList *ContextList,
	aliasList *AliasList,
	componentList *ComponentList,
	defaultContext string) *ComponentBuilder {

	if contextList == nil {
		panic(ErrComponentBuilderContextListShouldNotBeNil)
	}
	if aliasList == nil {
		panic(ErrComponentBuilderAliasListShouldNotBeNil)
	}
	if componentList == nil {
		panic(ErrComponentBuilderComponentListShouldNotBeNil)
	}
	return &ComponentBuilder{
		contextList:               contextList,
		aliasList:                 aliasList,
		componentList:             componentList,
		targetContextList:         NewContextList(),
		targetComponentMemberList: NewComponentMemberList(),
		defaultContext:            defaultContext,
	}
}

// SetID ...
func (cpb *ComponentBuilder) SetID(id string) *ComponentBuilder {
	cpb.id = id
	return cpb
}

// SetFlagPrefix ...
func (cpb *ComponentBuilder) SetFlagPrefix(p string) *ComponentBuilder {
	cpb.flagPrefix = p
	return cpb
}

// SetUnique ...
func (cpb *ComponentBuilder) SetUnique(u bool) *ComponentBuilder {
	cpb.unique = u
	return cpb
}

// SetEventTarget ...
func (cpb *ComponentBuilder) SetEventTarget(t EventTarget) *ComponentBuilder {
	cpb.eventTarget = t
	return cpb
}

// SetEventType ...
func (cpb *ComponentBuilder) SetEventType(e EventType) *ComponentBuilder {
	cpb.eventType = e
	return cpb
}

// SetEventPriority ...
func (cpb *ComponentBuilder) SetEventPriority(v int) *ComponentBuilder {
	cpb.eventPriority = v
	return cpb
}

// SetCleanupMode ...
func (cpb *ComponentBuilder) SetCleanupMode(m CleanupMode) *ComponentBuilder {
	cpb.cleanupMode = m
	return cpb
}

// AddContext ...
func (cpb *ComponentBuilder) AddContext(id string) error {
	ctx := cpb.contextList.ContextWithID(id)
	return cpb.targetContextList.AddContext(ctx)
}

// NewMember ...
func (cpb *ComponentBuilder) NewMember() *ComponentMemberBuilder {
	return NewComponentMemberBuilder(cpb.aliasList, cpb.targetComponentMemberList)
}

// Build ...
func (cpb *ComponentBuilder) Build() error {
	if cpb.built {
		return ErrComponentBuilderComponentAlreadyBuilt
	}
	if len(cpb.targetContextList.ContextSlice()) == 0 {
		defaultContext := cpb.contextList.ContextWithID(cpb.defaultContext)
		if defaultContext == nil {
			return ErrComponentBuilderComponentMustHaveContext
		}
		cpb.targetContextList.AddContext(defaultContext)
	}
	component, err := NewComponent(cpb.id, cpb.flagPrefix, cpb.unique, cpb.eventTarget, cpb.eventType, cpb.eventPriority, cpb.cleanupMode, cpb.targetContextList, cpb.targetComponentMemberList)
	if err != nil {
		return err
	}
	err = cpb.componentList.AddComponent(component)
	if err != nil {
		return err
	}
	cpb.built = true
	return nil
}
