package model

import "bytes"

// EntityIndex ...
type EntityIndex struct {
	id        string
	isPrimary bool
	context   *Context
	method    map[string]*EntityIndexMethod
}

// NewEntityIndexWith ...
func NewEntityIndexWith(id string, context *Context) *EntityIndex {
	return &EntityIndex{
		id:        id,
		isPrimary: false,
		context:   context,
		method:    make(map[string]*EntityIndexMethod, 0)}
}

// NewEntityIndexWithID ...
func NewEntityIndexWithID(id string) *EntityIndex {
	return NewEntityIndexWith(id, nil)
}

// NewEntityIndex ...
func NewEntityIndex() *EntityIndex {
	return NewEntityIndexWithID("Untitled")
}

// GetID ...
func (i *EntityIndex) GetID() String {
	return String(i.id)
}

// SetID ...
func (i *EntityIndex) SetID(id string) *EntityIndex {
	i.id = id
	return i
}

// GetContext ...
func (i *EntityIndex) GetContext() *Context {
	return i.context
}

// SetContext ...
func (i *EntityIndex) SetContext(context *Context) *EntityIndex {
	i.context = context
	return i
}

// IsPrimary ...
func (i *EntityIndex) IsPrimary() bool {
	return i.isPrimary
}

// SetPrimary ...
func (i *EntityIndex) SetPrimary(value bool) *EntityIndex {
	i.isPrimary = value
	return i
}

// GetEntityIndexMethodWithID ...
func (i *EntityIndex) GetEntityIndexMethodWithID(id string) *EntityIndexMethod {
	m, _ := i.method[id]
	return m
}

// GetEntityIndexMethod ...
func (i *EntityIndex) GetEntityIndexMethod() []*EntityIndexMethod {
	slice := make([]*EntityIndexMethod, 0)
	for _, m := range i.method {
		slice = append(slice, m)
	}
	return slice
}

// AddEntityIndexMethod ...
func (i *EntityIndex) AddEntityIndexMethod(overwrite bool, method ...*EntityIndexMethod) *EntityIndex {
	for _, m := range method {
		_, exist := i.method[m.GetID().String()]
		if overwrite || !overwrite && !exist {
			i.method[m.GetID().String()] = m
		}
	}
	return i
}

// CreateEntityIndexMethod ...
func (i *EntityIndex) CreateEntityIndexMethod(overwrite bool) *EntityIndexMethod {
	m := NewEntityIndexMethod()
	i.AddEntityIndexMethod(overwrite, m)
	return m
}

// String ...
func (i EntityIndex) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(i.GetID()))
	buffer.WriteRune('(')
	if i.IsPrimary() {
		buffer.WriteString("[primary]")
	}
	buffer.WriteString(") ")
	for _, Method := range i.GetEntityIndexMethod() {
		buffer.WriteString("Method{")
		buffer.WriteString(Method.String())
		buffer.WriteString("} ")
	}
	return buffer.String()
}
