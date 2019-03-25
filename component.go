package proton

import (
	"fmt"
)

var (
	ErrComponentIDUndefined        = fmt.Errorf("Component: `ID` Undefined.")
	ErrComponentEventTargetInvalid = fmt.Errorf("Component: `EventTarget` Invalid.")
	ErrComponentEventTypeInvalid   = fmt.Errorf("Component: `EventType` Invalid.")
	ErrComponentCleanupModeInvalid = fmt.Errorf("Component: `CleanupMode` Invalid.")
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
	cl             *CL
	cml            *CML
	id, flagPrefix string
	unique         bool
	eventTarget    EventTarget
	eventType      EventType
	eventPriority  int
	cleanupMode    CleanupMode
}

// NewComponent ...
func NewComponent(id, flagPrefix string, unique bool, eventTarget EventTarget, eventType EventType, eventPriority int, cleanupMode CleanupMode, cl *CL, cml *CML) (*CP, error) {
	if id == "" {
		return nil, ErrComponentIDUndefined
	}
	if !eventTarget.IsValid() {
		return nil, ErrComponentEventTargetInvalid
	}
	if !eventType.IsValid() {
		return nil, ErrComponentEventTypeInvalid
	}
	if !cleanupMode.IsValid() {
		return nil, ErrComponentCleanupModeInvalid
	}
	return &CP{cl, cml, id, flagPrefix, unique, eventTarget, eventType, eventPriority, cleanupMode}, nil
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
	return cp.cl.ContextWithID(id)
}

// ContextSlice ...
func (cp *CP) ContextSlice() []*C {
	return cp.cl.ContextSlice()
}

// MembersWithEntityIndex ...
func (cp *CP) MembersWithEntityIndex() []*CM {
	return cp.cml.MembersWithEntityIndex()
}

// MemberSlice ...
func (cp *CP) MemberSlice() []*CM {
	return cp.cml.MemberSlice()
}
