package proton

import (
	"fmt"
)

var (
	ErrComponentBuilderContextListShouldNotBeNil   = fmt.Errorf("ComponentBuilder: `ContextList` should not be nil.")
	ErrComponentBuilderAliasListShouldNotBeNil     = fmt.Errorf("ComponentBuilder: `AliasList` should not be nil.")
	ErrComponentBuilderComponentListShouldNotBeNil = fmt.Errorf("ComponentBuilder: `ComponentList` should not be nil.")
	ErrComponentBuilderComponentAlreadyBuilt       = fmt.Errorf("ComponentBuilder: `Component` already built.")
	ErrComponentBuilderComponentMustHaveContext    = fmt.Errorf("ComponentBuilder: `Component` does not have a context.")
)

// CPB ...
type CPB struct {
	cpl            *CPL
	tcl            *CL
	tcml           *CML
	cl             *CL
	al             *AL
	built          bool
	dctx           string
	id, flagPrefix string
	unique         bool
	eventTarget    EventTarget
	eventType      EventType
	eventPriority  int
	cleanupMode    CleanupMode
}

// NewComponentBuilder ...
func NewComponentBuilder(cl *CL, al *AL, cpl *CPL, dctx string) *CPB {
	if cl == nil {
		panic(ErrComponentBuilderContextListShouldNotBeNil)
	}
	if al == nil {
		panic(ErrComponentBuilderAliasListShouldNotBeNil)
	}
	if cpl == nil {
		panic(ErrComponentBuilderComponentListShouldNotBeNil)
	}
	return &CPB{
		cl:   cl,
		al:   al,
		cpl:  cpl,
		tcl:  NewContextList(),
		tcml: NewComponentMemberList(),
		dctx: dctx,
	}
}

// SetID ...
func (cpb *CPB) SetID(id string) *CPB {
	cpb.id = id
	return cpb
}

// SetFlagPrefix ...
func (cpb *CPB) SetFlagPrefix(p string) *CPB {
	cpb.flagPrefix = p
	return cpb
}

// SetUnique ...
func (cpb *CPB) SetUnique(u bool) *CPB {
	cpb.unique = u
	return cpb
}

// SetEventTarget ...
func (cpb *CPB) SetEventTarget(t EventTarget) *CPB {
	cpb.eventTarget = t
	return cpb
}

// SetEventType ...
func (cpb *CPB) SetEventType(e EventType) *CPB {
	cpb.eventType = e
	return cpb
}

// SetEventPriority ...
func (cpb *CPB) SetEventPriority(v int) *CPB {
	cpb.eventPriority = v
	return cpb
}

// SetCleanupMode ...
func (cpb *CPB) SetCleanupMode(m CleanupMode) *CPB {
	cpb.cleanupMode = m
	return cpb
}

// AddContext ...
func (cpb *CPB) AddContext(id string) *CPB {
	ctx := cpb.cl.ContextWithID(id)
	cpb.tcl.AddContext(ctx)
	return cpb
}

// NewMember ...
func (cpb *CPB) NewMember() *CPMB {
	return NewComponentMemberBuilder(cpb.al, cpb.tcml)
}

// Build ...
func (cpb *CPB) Build() error {
	if cpb.built {
		return ErrComponentBuilderComponentAlreadyBuilt
	}
	if len(cpb.tcl.ContextSlice()) == 0 {
		dctx := cpb.cl.ContextWithID(cpb.dctx)
		if dctx == nil {
			return ErrComponentBuilderComponentMustHaveContext
		}
		cpb.tcl.AddContext(dctx)
	}
	cp, err := NewComponent(cpb.id, cpb.flagPrefix, cpb.unique, cpb.eventTarget, cpb.eventType, cpb.eventPriority, cpb.cleanupMode, cpb.tcl, cpb.tcml)
	if err != nil {
		return err
	}
	err = cpb.cpl.AddComponent(cp)
	if err != nil {
		return err
	}
	cpb.built = true
	return nil
}
