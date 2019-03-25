package proton

import (
	"fmt"
	"unicode"
)

var (
	ErrContextIDUndefined          = fmt.Errorf("entitas(context): id undefined")
	ErrContextIDContainsWhitespace = fmt.Errorf("entitas(context): id contains whitespace")
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
	for _, v := range id {
		if unicode.IsSpace(v) {
			return nil, ErrContextIDContainsWhitespace
		}
	}
	return &C{id}, nil
}

// ID ...
func (c *C) ID() String {
	return String(c.id)
}
