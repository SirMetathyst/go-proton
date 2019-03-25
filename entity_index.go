package proton

import "fmt"

var (
	ErrEntityIndexIDUndefined     = fmt.Errorf("EntityIndex: `ID` Undefined.")
	ErrEntityIndexMustHaveContext = fmt.Errorf("EntityIndex: EntityIndex does not have a `Context`.")
	ErrEntityIndexMustHaveMethod  = fmt.Errorf("EntityIndex: EntityIndex does not have a `Method`.")
)

// EntityIndex ...
type EntityIndex int

const (
	NoEntityIndex EntityIndex = iota
	SingleEntityIndex
	MultipleEntityIndex
)

// IsValid ...
func (i EntityIndex) IsValid() bool {
	return i >= 0 || i <= 2
}

// String ...
func (i EntityIndex) String() string {
	switch i {
	case SingleEntityIndex:
		return "PrimaryEntityIndex"
	case MultipleEntityIndex:
		return "EntityIndex"
	}
	return ""
}

// EI ...
type EI struct {
	id        string
	isPrimary bool
	c         *C
	eiml      *EIML
}

// NewEntityIndex ...
func NewEntityIndex(id string, isPrimary bool, c *C, eiml *EIML) (*EI, error) {
	if id == "" {
		return nil, ErrEntityIndexIDUndefined
	}
	if c == nil {
		return nil, ErrEntityIndexMustHaveContext
	}
	if len(eiml.EntityIndexMethodList()) == 0 {
		return nil, ErrEntityIndexMustHaveMethod
	}
	return &EI{id, isPrimary, c, eiml}, nil
}

// ID ...
func (ei *EI) ID() String {
	return String(ei.id)
}

// IsPrimary ...
func (ei *EI) IsPrimary() bool {
	return ei.isPrimary
}

// Context ...
func (ei *EI) Context() *C {
	return ei.c
}

// EntityIndexMethodList ...
func (ei *EI) EntityIndexMethodList() []*EIM {
	return ei.eiml.EntityIndexMethodList()
}
