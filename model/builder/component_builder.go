package modelbuilder

import (
	"fmt"

	"github.com/SirMetathyst/proton/model"
)

var (
	ErrComponentBuilderComponentAlreadyBuilt    = fmt.Errorf("ComponentBuilder: Component already built.")
	ErrComponentBuilderComponentMustHaveContext = fmt.Errorf("ComponentBuilder: Component does not have a context.")
)

// CPB ...
type CPB struct {
	outList          *model.CPL
	targetCtxList    model.CL
	targetMemberList model.ML
	ctxList          *model.CL
	dctx             string
	aliasList        *model.AL
	built            bool
	id, flagPrefix   string
	unique           bool
	eventTarget      model.EventTarget
	eventType        model.EventType
	eventPriority    int
	listener         bool
	cleanupMode      model.CleanupMode
}

// NewComponentBuilder ...
func NewComponentBuilder(outList *model.CPL, ctxList *model.CL, dctx string, aliasList *model.AL) *CPB {
	return &CPB{outList: outList, ctxList: ctxList, dctx: dctx, aliasList: aliasList}
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
func (cpb *CPB) SetEventTarget(t model.EventTarget) *CPB {
	cpb.eventTarget = t
	return cpb
}

// SetEventType ...
func (cpb *CPB) SetEventType(e model.EventType) *CPB {
	cpb.eventType = e
	return cpb
}

// SetEventPriority ...
func (cpb *CPB) SetEventPriority(v int) *CPB {
	cpb.eventPriority = v
	return cpb
}

// SetListener ...
func (cpb *CPB) SetListener(v bool) *CPB {
	cpb.listener = v
	return cpb
}

// SetCleanupMode ...
func (cpb *CPB) SetCleanupMode(m model.CleanupMode) *CPB {
	cpb.cleanupMode = m
	return cpb
}

// AddContext ...
func (cpb *CPB) AddContext(id string) *CPB {
	ctx := cpb.ctxList.ContextWithID(id)
	cpb.targetCtxList.AddContext(ctx)
	return cpb
}

// NewMember ...
func (cpb *CPB) NewMember() *MB {
	return NewMemberBuilder(&cpb.targetMemberList, cpb.aliasList)
}

// Build ...
func (cpb *CPB) Build() error {
	if len(cpb.targetCtxList.ContextList()) == 0 {
		dctx := cpb.ctxList.ContextWithID(cpb.dctx)
		if dctx == nil {
			return ErrComponentBuilderComponentMustHaveContext
		}
		cpb.targetCtxList.AddContext(dctx)
	}
	if cpb.built {
		return ErrComponentBuilderComponentAlreadyBuilt
	}
	cp, err := model.NewComponent(cpb.id, cpb.flagPrefix, cpb.unique, cpb.eventTarget, cpb.eventType, cpb.eventPriority, cpb.listener, cpb.cleanupMode, cpb.targetCtxList, cpb.targetMemberList)
	if err != nil {
		return err
	}
	err = cpb.outList.AddComponent(cp)
	if err != nil {
		return err
	}
	return nil
}
