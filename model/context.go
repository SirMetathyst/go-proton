package model

import "bytes"

// Context ...
type Context struct {
	id        string
	isDefault bool
}

// NewContextWithID ...
func NewContextWithID(id string) *Context {
	return &Context{id: id}
}

// NewContext ...
func NewContext() *Context {
	return new(Context)
}

// GetID ...
func (c *Context) GetID() String {
	return String(c.id)
}

// SetID ...
func (c *Context) SetID(id string) *Context {
	c.id = id
	return c
}

// IsDefault ...
func (c *Context) IsDefault() bool {
	return c.isDefault
}

// SetDefault ...
func (c *Context) SetDefault(value bool) *Context {
	c.isDefault = value
	return c
}

// String ...
func (c Context) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(c.GetID()))
	buffer.WriteRune('(')
	if c.IsDefault() {
		buffer.WriteString("Default")
	}
	buffer.WriteRune(')')
	return buffer.String()
}
