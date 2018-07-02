package model

import (
	"bytes"

	"github.com/SirMetathyst/proton/configuration"
)

// Provider ...
type Provider func(*configuration.C) (*M, error)

// M ...
type M struct {
	target      string
	namespace   string
	context     map[string]*Context
	alias       map[string]*Alias
	component   map[string]*Component
	entityIndex map[string]*EntityIndex
}

// NewModel ...
func NewModel() *M {
	return &M{
		target:      "entitas_csharp",
		namespace:   "",
		context:     make(map[string]*Context, 0),
		alias:       make(map[string]*Alias, 0),
		component:   make(map[string]*Component, 0),
		entityIndex: make(map[string]*EntityIndex, 0)}
}

var (
	model = NewModel()
)

// GetModel ...
func GetModel() *M {
	return model
}

// GetTarget ...
func (m *M) GetTarget() String {
	return String(m.target)
}

// SetTarget ...
func (m *M) SetTarget(value string) *M {
	m.target = value
	return m
}

// GetNamespace ...
func (m *M) GetNamespace() String {
	return String(m.namespace)
}

// SetNamespace ...
func (m *M) SetNamespace(value string) *M {
	m.namespace = value
	return m
}

// GetContextWithID ...
func (m *M) GetContextWithID(id string) *Context {
	ctx, _ := m.context[id]
	return ctx
}

// GetContext ...
func (m *M) GetContext() []*Context {
	slice := make([]*Context, 0)
	for _, ctx := range m.context {
		slice = append(slice, ctx)
	}
	return slice
}

// GetDefaultContext ...
func (m *M) GetDefaultContext() *Context {
	for _, ctx := range m.context {
		if ctx.IsDefault() {
			return ctx
		}
	}
	return nil
}

// AddContext ...
func (m *M) AddContext(overwrite bool, context ...*Context) *M {
	for _, ctx := range context {
		if ctx != nil {
			_, exist := m.context[ctx.GetID().String()]
			if overwrite || !overwrite && !exist {
				m.context[ctx.GetID().String()] = ctx
			}
		}
	}
	return m
}

// CreateContext ...
func (m *M) CreateContext(overwrite bool, id string) *Context {
	ctx := NewContextWithID(id)
	m.AddContext(overwrite, ctx)
	return ctx
}

// GetAliasWithID ...
func (m *M) GetAliasWithID(id string) *Alias {
	alias, _ := m.alias[id]
	return alias
}

// GetAlias ...
func (m *M) GetAlias() []*Alias {
	slice := make([]*Alias, 0)
	for _, alias := range m.alias {
		slice = append(slice, alias)
	}
	return slice
}

// AddAlias ...
func (m *M) AddAlias(overwrite bool, alias ...*Alias) *M {
	for _, alias := range alias {
		if alias != nil {
			_, exist := m.context[alias.GetID().String()]
			if overwrite || !overwrite && !exist {
				m.alias[alias.GetID().String()] = alias
			}
		}
	}
	return m
}

// GetComponentWithID ...
func (m *M) GetComponentWithID(id string) *Component {
	c, _ := m.component[id]
	return c
}

// GetComponentWithContextID ...
func (m *M) GetComponentWithContextID(id string) []*Component {
	slice := make([]*Component, 0)
	for _, c := range m.component {
		if ctx := c.GetContextWithID(id); ctx != nil {
			slice = append(slice, c)
		}
	}
	return slice
}

// GetUniqueComponent ...
func (m *M) GetUniqueComponent() []*Component {
	slice := make([]*Component, 0)
	for _, c := range m.component {
		if c.IsUnique() {
			slice = append(slice, c)
		}
	}
	return slice
}

// GetComponentWithEntityIndex ...
func (m *M) GetComponentWithEntityIndex() []*Component {
	slice := make([]*Component, 0)
	for _, c := range m.component {
		if c.GetEntityIndexCount() > 0 {
			slice = append(slice, c)
		}
	}
	return slice
}

// GetComponentWithEvent ...
func (m *M) GetComponentWithEvent() []*Component {
	slice := make([]*Component, 0)
	for _, c := range m.component {
		if c.IsEvent() {
			slice = append(slice, c)
		}
	}
	return slice
}

// GetComponent ...
func (m *M) GetComponent() []*Component {
	slice := make([]*Component, 0)
	for _, c := range m.component {
		slice = append(slice, c)
	}
	return slice
}

// AddComponent ...
func (m *M) AddComponent(overwrite bool, component ...*Component) *M {
	for _, component := range component {
		if component != nil {
			_, exist := m.component[component.GetID().String()]
			if overwrite || !overwrite && !exist {
				m.component[component.GetID().String()] = component
			}
		}
	}
	return m
}

// CreateComponent ...
func (m *M) CreateComponent(overwrite bool, id string) *Component {
	c := NewComponentWithID(id)
	m.AddComponent(overwrite, c)
	return c
}

// GetEntityIndexWithID ...
func (m *M) GetEntityIndexWithID(id string) *EntityIndex {
	i, _ := m.entityIndex[id]
	return i
}

// GetEntityIndex ...
func (m *M) GetEntityIndex() []*EntityIndex {
	slice := make([]*EntityIndex, 0)
	for _, i := range m.entityIndex {
		slice = append(slice, i)
	}
	return slice
}

// AddEntityIndex ...
func (m *M) AddEntityIndex(overwrite bool, index ...*EntityIndex) *M {
	for _, i := range index {
		_, exist := m.entityIndex[i.GetID().String()]
		if overwrite || !overwrite && !exist {
			m.entityIndex[i.GetID().String()] = i
		}
	}
	return m
}

// CreateEntityIndex ...
func (m *M) CreateEntityIndex(overwrite bool) *EntityIndex {
	i := NewEntityIndex()
	m.AddEntityIndex(overwrite, i)
	return i
}

// Refresh ...
func (m *M) Refresh() {
	for _, c := range m.GetComponent() {
		if len(c.GetContext()) == 0 {
			ctx := m.GetDefaultContext()
			if ctx != nil {
				c.AddContext(true, ctx)
			}
		}
	}
}

func (m *M) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("Target(")
	buffer.WriteString(m.GetTarget().String())
	buffer.WriteString(")\n\nNamespace(")
	buffer.WriteString(m.GetNamespace().String())
	buffer.WriteString(")\n\n")
	for _, Context := range m.GetContext() {
		buffer.WriteString("Context(")
		buffer.WriteString(Context.String())
		buffer.WriteString(")\n")
	}
	buffer.WriteString("\n")
	for _, Alias := range m.GetAlias() {
		buffer.WriteString("Alias(")
		buffer.WriteString(Alias.String())
		buffer.WriteString(")\n")
	}
	buffer.WriteString("\n")
	for _, Component := range m.GetComponent() {
		buffer.WriteString("Component(")
		buffer.WriteString(Component.String())
		buffer.WriteString(")\n")
	}
	buffer.WriteString("\n")
	for _, EntityIndex := range m.GetEntityIndex() {
		buffer.WriteString("EntityIndex(")
		buffer.WriteString(EntityIndex.String())
		buffer.WriteString(")\n")
	}
	return buffer.String()
}
