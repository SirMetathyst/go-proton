package model

import (
	"bytes"
)

// Component ...
type Component struct {
	id, prefix              string
	unique                  bool
	eventBinding, eventType int
	context                 map[string]*Context
	member                  map[string]*Member
}

// NewComponentWithID ...
func NewComponentWithID(id string) *Component {
	return &Component{
		id:           id,
		unique:       false,
		eventBinding: 0,
		eventType:    0,
		prefix:       "",
		context:      make(map[string]*Context, 0),
		member:       make(map[string]*Member, 0)}
}

// NewComponent ...
func NewComponent() *Component {
	return NewComponentWithID("")
}

// GetID ...
func (c *Component) GetID() String {
	return String(c.id)
}

// SetID ...
func (c *Component) SetID(id string) *Component {
	if c.id == "" {
		c.id = id
	}
	return c
}

// HasPrefix ...
func (c *Component) HasPrefix() bool {
	return c.prefix != ""
}

// GetPrefix ...
func (c *Component) GetPrefix() String {
	return String(c.prefix)
}

// GetPrefixOrDefault ...
func (c *Component) GetPrefixOrDefault() String {
	if c.HasPrefix() {
		return String(c.prefix)
	} else if len(c.GetMember()) <= 0 {
		return String("is")
	}
	return String("has")
}

// SetPrefix ...
func (c *Component) SetPrefix(prefix string) *Component {
	c.prefix = prefix
	return c
}

// IsUnique ...
func (c *Component) IsUnique() bool {
	return c.unique
}

// SetUnique ...
func (c *Component) SetUnique(value bool) *Component {
	c.unique = value
	return c
}

// IsEvent ...
func (c *Component) IsEvent() bool {
	return c.eventBinding > 0
}

// SetEventBinding ...
func (c *Component) SetEventBinding(value int) *Component {
	c.eventBinding = InRange(value, 0, 2)
	return c
}

// GetEventBinding ...
func (c *Component) GetEventBinding() int {
	return c.eventBinding
}

// GetEventBindingString ...
func (c *Component) GetEventBindingString() String {
	switch c.eventBinding {
	case 1:
		return String("self")
	case 2:
		return String("global")
	}
	return String("")
}

// SetEventType ...
func (c *Component) SetEventType(value int) *Component {
	c.eventType = InRange(value, 0, 1)
	return c
}

// GetEventType ...
func (c *Component) GetEventType() int {
	return c.eventType
}

// GetEventTypeSuffix ...
func (c *Component) GetEventTypeSuffix() String {
	switch c.eventType {
	case 1:
		return String("Removed")
	}
	return String("")
}

// GetEntityIndexCount ...
func (c *Component) GetEntityIndexCount() int {
	count := 0
	for _, m := range c.GetMember() {
		if m.GetEntityIndex() > 0 {
			count++
		}
	}
	return count
}

// GetContextWithID ...
func (c *Component) GetContextWithID(id string) *Context {
	ctx, _ := c.context[id]
	return ctx
}

// GetContext ...
func (c *Component) GetContext() []*Context {
	slice := make([]*Context, 0)
	for _, ctx := range c.context {
		slice = append(slice, ctx)
	}
	return slice
}

// GetDefaultContext ...
func (c *Component) GetDefaultContext() *Context {
	for _, ctx := range c.context {
		if ctx.IsDefault() {
			return ctx
		}
	}
	return nil
}

// AddContext ...
func (c *Component) AddContext(overwrite bool, context ...*Context) *Component {
	for _, ctx := range context {
		if ctx != nil {
			_, exist := c.context[ctx.GetID().String()]
			if overwrite || !overwrite && !exist {
				c.context[ctx.GetID().String()] = ctx
			}
		}
	}
	return c
}

// GetMemberWithID ...
func (c *Component) GetMemberWithID(id string) *Member {
	m, _ := c.member[id]
	return m
}

// GetMember ...
func (c *Component) GetMember() []*Member {
	slice := make([]*Member, 0)
	for _, m := range c.member {
		slice = append(slice, m)
	}
	return slice
}

// AddMember ...
func (c *Component) AddMember(overwrite bool, member ...*Member) *Component {
	for _, m := range member {
		_, exist := c.member[m.GetID().String()]
		if overwrite || !overwrite && !exist {
			c.member[m.GetID().String()] = m
		}
	}
	return c
}

// CreateMember ...
func (c *Component) CreateMember(overwrite bool, id string) *Member {
	m := NewMemberWithID(id)
	c.AddMember(overwrite, m)
	return m
}

// String ...
func (c Component) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(c.GetID()))
	buffer.WriteRune('(')
	if c.IsUnique() {
		buffer.WriteString("[unique]")
	}
	if c.IsEvent() {
		buffer.WriteString("[event:")
		buffer.WriteString(c.GetEventBindingString().String())
		buffer.WriteRune(']')
		buffer.WriteString("[eventType:")
		buffer.WriteString(c.GetEventTypeSuffix().String())
		buffer.WriteRune(']')
	}

	if c.HasPrefix() {
		buffer.WriteString("[prefix:")
		buffer.WriteString(c.GetPrefix().String())
		buffer.WriteRune(']')
	}

	buffer.WriteString(") ")
	for _, ctx := range c.GetContext() {
		buffer.WriteString("Context(")
		buffer.WriteString(ctx.String())
		buffer.WriteRune(')')
	}
	buffer.WriteString(") ")
	for _, m := range c.GetMember() {
		buffer.WriteString("Member{")
		buffer.WriteString(m.String())
		buffer.WriteRune('}')
	}

	return buffer.String()
}
