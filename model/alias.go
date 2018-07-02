package model

// Alias ...
type Alias struct {
	id, value string
}

// NewAliasWith ...
func NewAliasWith(id, value string) *Alias {
	return &Alias{id, value}
}

// NewAlias ...
func NewAlias() *Alias {
	return &Alias{}
}

// GetID ...
func (a *Alias) GetID() String {
	return String(a.id)
}

// SetID ...
func (a *Alias) SetID(id string) *Alias {
	if a.id == "" {
		a.id = id
	}
	return a
}

// GetValue ...
func (a *Alias) GetValue() String {
	return String(a.value)
}

// SetValue ...
func (a *Alias) SetValue(value string) *Alias {
	a.value = value
	return a
}

// String ...
func (a Alias) String() string {
	return "[" + a.GetID().String() + ":\"" + a.GetValue().String() + "\"]"
}
