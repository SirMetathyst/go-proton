package proton

import "errors"

var (
	// ErrEntityIndexIDUndefined ...
	ErrEntityIndexIDUndefined = errors.New("proton: entity index: id undefined")
	// ErrEntityIndexMustHaveContext ...
	ErrEntityIndexMustHaveContext = errors.New("proton: entity index: entity index does not have a context")
	// ErrEntityIndexMustHaveMethod ...
	ErrEntityIndexMustHaveMethod = errors.New("proton: entity index: entity index does not have a method")
)

// EntityIndexType ...
type EntityIndexType int

const (
	// EntityIndexTypeNone ...
	EntityIndexTypeNone EntityIndexType = iota
	// EntityIndexTypeSingle ...
	EntityIndexTypeSingle
	// EntityIndexTypeMultiple ...
	EntityIndexTypeMultiple
)

// IsValid ...
func (i EntityIndexType) IsValid() bool {
	return i >= 0 || i <= 2
}

// String ...
func (i EntityIndexType) String() string {
	switch i {
	case EntityIndexTypeNone:
		return ""
	case EntityIndexTypeSingle:
		return "PrimaryEntityIndex"
	case EntityIndexTypeMultiple:
		return "EntityIndex"
	}
	return "UNKNOWN"
}

// EntityIndex ...
type EntityIndex struct {
	id                    string
	isPrimary             bool
	context               *Context
	entityIndexMethodList *EntityIndexMethodList
}

// NewEntityIndex ...
func NewEntityIndex(
	id string,
	isPrimary bool,
	context *Context,
	entityIndexMethodList *EntityIndexMethodList) (*EntityIndex, error) {
	if id == "" {
		return nil, ErrEntityIndexIDUndefined
	}
	if context == nil {
		panic(ErrEntityIndexMustHaveContext)
	}
	if len(entityIndexMethodList.EntityIndexMethodSlice()) == 0 {
		return nil, ErrEntityIndexMustHaveMethod
	}
	return &EntityIndex{
		id:                    id,
		isPrimary:             isPrimary,
		context:               context,
		entityIndexMethodList: entityIndexMethodList}, nil
}

// ID ...
func (ei *EntityIndex) ID() String {
	return String(ei.id)
}

// IsPrimary ...
func (ei *EntityIndex) IsPrimary() bool {
	return ei.isPrimary
}

// Context ...
func (ei *EntityIndex) Context() *Context {
	return ei.context
}

// EntityIndexMethodSlice ...
func (ei *EntityIndex) EntityIndexMethodSlice() []*EntityIndexMethod {
	return ei.entityIndexMethodList.EntityIndexMethodSlice()
}
