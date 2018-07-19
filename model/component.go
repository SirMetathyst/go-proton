package model

import (
	"fmt"
)

var (
	ErrComponentIDUndefined = fmt.Errorf("Component: `ID` Undefined.")
)

// EventTarget ...
type EventTarget int

const (
	NoTarget EventTarget = iota
	SelfTarget
	AnyTarget
)

// IsValid ...
func (e EventTarget) IsValid() bool {
	return e >= 0 || e <= 2
}

// AsValid ...
func (e EventTarget) AsValid() EventTarget {
	if e.IsValid() {
		return e
	}
	return NoTarget
}

// String ...
func (e EventTarget) String() String {
	switch e {
	case SelfTarget:
		return String("Self")
	case AnyTarget:
		return String("Any")
	}
	return String("")
}

// EventType ...
type EventType int

const (
	AddedEvent EventType = iota
	RemovedEvent
)

// IsValid ...
func (t EventType) IsValid() bool {
	return t == 0 || t == 1
}

// AsValid ...
func (t EventType) AsValid() EventType {
	if t.IsValid() {
		return t
	}
	return AddedEvent
}

// String ...
func (t EventType) String() String {
	switch t {
	case AddedEvent:
		return String("Added")
	case RemovedEvent:
		return String("Removed")
	}
	return String("")
}

// CleanupMode ...
type CleanupMode int

const (
	NoCleanup CleanupMode = iota
	DestroyEntity
	RemoveComponent
)

// IsValid ...
func (m CleanupMode) IsValid() bool {
	return m == 0 || m == 1
}

// AsValid ...
func (m CleanupMode) AsValid() CleanupMode {
	if m.IsValid() {
		return m
	}
	return NoCleanup
}

// String ...
func (m CleanupMode) String() String {
	switch m {
	case DestroyEntity:
		return String("DestroyEntity")
	case RemoveComponent:
		return String("RemoveComponent")
	}
	return String("")
}

// CP ...
type CP struct {
	id, flagPrefix string
	unique         bool
	eventTarget    EventTarget
	eventType      EventType
	eventPriority  int
	listener       bool
	cleanupMode    CleanupMode
	ctxList        CL
	memberList     ML
}

// NewComponent ...
func NewComponent(id, flagPrefix string, unique bool, eventTarget EventTarget, eventType EventType, eventPriority int, listener bool, cleanupMode CleanupMode, ctxList CL, memberList ML) (*CP, error) {
	if id == "" {
		return nil, ErrComponentIDUndefined
	}
	return &CP{id, flagPrefix, unique, eventTarget.AsValid(), eventType.AsValid(), eventPriority, listener, cleanupMode.AsValid(), ctxList, memberList}, nil
}

// ID ...
func (cp *CP) ID() String {
	return String(cp.id)
}

// HasFlagPrefix ...
func (cp *CP) HasFlagPrefix() bool {
	return cp.flagPrefix != ""
}

// FlagPrefix ...
func (cp *CP) FlagPrefix() String {
	return String(cp.flagPrefix)
}

// FlagPrefixOrDefault ...
func (cp *CP) FlagPrefixOrDefault() String {
	if cp.HasFlagPrefix() {
		return String(cp.flagPrefix)
	}
	return String("is")
}

// IsUnique ...
func (cp *CP) IsUnique() bool {
	return cp.unique
}

// IsEvent ...
func (cp *CP) IsEvent() bool {
	return cp.eventTarget > 0
}

// EventTarget ...
func (cp *CP) EventTarget() EventTarget {
	return cp.eventTarget
}

// EventType ...
func (cp *CP) EventType() EventType {
	return cp.eventType
}

// EventPriority ...
func (cp *CP) EventPriority() int {
	return cp.eventPriority
}

// IsListener ...
func (cp *CP) IsListener() bool {
	return cp.listener
}

// IsCleanup ...
func (cp *CP) IsCleanup() bool {
	return cp.cleanupMode > 0
}

// CleanupMode ...
func (cp *CP) CleanupMode() CleanupMode {
	return cp.cleanupMode
}

// ContextWithID ...
func (cp *CP) ContextWithID(id string) *C {
	return cp.ctxList.ContextWithID(id)
}

// ContextList ...
func (cp *CP) ContextList() []*C {
	return cp.ctxList.ContextList()
}

// MembersWithEntityIndex ...
func (cp *CP) MembersWithEntityIndex() []*M {
	return cp.memberList.MembersWithEntityIndex()
}

// MemberList ...
func (cp *CP) MemberList() []*M {
	return cp.memberList.MemberList()
}
