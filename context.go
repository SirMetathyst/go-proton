package proton

import (
	"errors"
)

var (
	// ErrContextIDUndefined ...
	ErrContextIDUndefined = errors.New("proton: context: id undefined")
	// ErrContextIDContainsWhitespace ...
	ErrContextIDContainsWhitespace = errors.New("proton: context: id contains whitespace")
)

// Context ...
type Context struct {
	id string
}

// NewContext ...
func NewContext(id string) (*Context, error) {
	if id == "" {
		return nil, ErrContextIDUndefined
	}
	if containsWhitespace(id) {
		return nil, ErrContextIDContainsWhitespace
	}
	return &Context{id}, nil
}

// ID ...
func (c *Context) ID() String {
	return String(c.id)
}
