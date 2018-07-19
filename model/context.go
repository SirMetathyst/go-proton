package model

import "fmt"

var (
	ErrContextIDUndefined = fmt.Errorf("Context: `ID` Undefined.")
)

// C ...
type C struct {
	id string
}

// NewContext ...
func NewContext(id string) (*C, error) {
	if id == "" {
		return nil, ErrContextIDUndefined
	}
	return &C{id}, nil
}

// ID ...
func (c *C) ID() String {
	return String(c.id)
}

// String ...
func (c C) String() string {
	return "[" + c.ID().String() + "]"
}
