package proton

import "errors"

var (
	// ErrComponentIDUndefined ...
	ErrComponentIDUndefined = errors.New("proton: component: id undefined")
	// ErrComponentEventTargetInvalid ...
	ErrComponentEventTargetInvalid = errors.New("proton: component: event target is invalid")
	// ErrComponentEventTypeInvalid ...
	ErrComponentEventTypeInvalid = errors.New("proton: component: event type is Invalid")
	// ErrComponentCleanupModeInvalid ...
	ErrComponentCleanupModeInvalid = errors.New("proton: component: cleanup mode is invalid")
)

// EventTarget ...
type EventTarget int

const (
	// EventTargetNone ...
	EventTargetNone EventTarget = iota
	// EventTargetSelf ...
	EventTargetSelf
	// EventTargetAny ...
	EventTargetAny
)

// IsValid ...
func (e EventTarget) IsValid() bool {
	return e >= 0 || e <= 2
}

// String ...
func (e EventTarget) String() String {
	switch e {
	case EventTargetNone:
		return String("")
	case EventTargetSelf:
		return String("Self")
	case EventTargetAny:
		return String("Any")
	}
	return String("UNKNOWN")
}

// EventType ...
type EventType int

const (
	// EventAdded ...
	EventAdded EventType = iota
	// EventAddedOrRemoved ...
	EventAddedOrRemoved
	// EventRemoved ...
	EventRemoved
)

// IsValid ...
func (t EventType) IsValid() bool {
	return t >= 0 || t <= 2
}

// String ...
func (t EventType) String() String {
	switch t {
	case EventAdded:
		return String("Added")
	case EventAddedOrRemoved:
		return String("AddedOrRemoved")
	case EventRemoved:
		return String("Removed")

	}
	return String("UNKNOWN")
}

// CleanupMode ...
type CleanupMode int

const (
	// CleanupModeNone ...
	CleanupModeNone CleanupMode = iota
	// CleanupModeEntity ...
	CleanupModeEntity
	// CleanupModeComponent ...
	CleanupModeComponent
)

// IsValid ...
func (m CleanupMode) IsValid() bool {
	return m == 0 || m == 2
}

// String ...
func (m CleanupMode) String() String {
	switch m {
	case CleanupModeNone:
		return String("")
	case CleanupModeEntity:
		return String("DestroyEntity")
	case CleanupModeComponent:
		return String("RemoveComponent")
	}
	return String("UNKNOWN")
}

// Component ...
type Component struct {
	contextList         *ContextList
	componentMemberList *ComponentMemberList
	id                  string
	flagPrefix          string
	unique              bool
	eventTarget         EventTarget
	eventType           EventType
	eventPriority       int
	cleanupMode         CleanupMode
}

// NewComponent ...
func NewComponent(
	id string,
	flagPrefix string,
	unique bool,
	eventTarget EventTarget,
	eventType EventType,
	eventPriority int,
	cleanupMode CleanupMode,
	contextList *ContextList,
	componentMemberList *ComponentMemberList) (*Component, error) {

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
	return &Component{
		contextList:         contextList,
		componentMemberList: componentMemberList,
		id:                  id,
		flagPrefix:          flagPrefix,
		unique:              unique,
		eventTarget:         eventTarget,
		eventType:           eventType,
		eventPriority:       eventPriority,
		cleanupMode:         cleanupMode}, nil
}

// ID ...
func (cp *Component) ID() String {
	return String(cp.id)
}

// HasFlagPrefix ...
func (cp *Component) HasFlagPrefix() bool {
	return cp.flagPrefix != ""
}

// FlagPrefix ...
func (cp *Component) FlagPrefix() String {
	return String(cp.flagPrefix)
}

// FlagPrefixOrDefault ...
func (cp *Component) FlagPrefixOrDefault() String {
	if cp.HasFlagPrefix() {
		return String(cp.flagPrefix)
	}
	return String("is")
}

// IsUnique ...
func (cp *Component) IsUnique() bool {
	return cp.unique
}

// IsEvent ...
func (cp *Component) IsEvent() bool {
	return cp.eventTarget > 0
}

// EventTarget ...
func (cp *Component) EventTarget() EventTarget {
	return cp.eventTarget
}

// EventType ...
func (cp *Component) EventType() EventType {
	return cp.eventType
}

// EventPriority ...
func (cp *Component) EventPriority() int {
	return cp.eventPriority
}

// IsCleanup ...
func (cp *Component) IsCleanup() bool {
	return cp.cleanupMode > 0
}

// CleanupMode ...
func (cp *Component) CleanupMode() CleanupMode {
	return cp.cleanupMode
}

// ContextWithID ...
func (cp *Component) ContextWithID(id string) *Context {
	return cp.contextList.ContextWithID(id)
}

// ContextSlice ...
func (cp *Component) ContextSlice() []*Context {
	return cp.contextList.ContextSlice()
}

// MembersWithEntityIndex ...
func (cp *Component) MembersWithEntityIndex() []*ComponentMember {
	return cp.componentMemberList.MembersWithEntityIndex()
}

// MemberSlice ...
func (cp *Component) MemberSlice() []*ComponentMember {
	return cp.componentMemberList.MemberSlice()
}
