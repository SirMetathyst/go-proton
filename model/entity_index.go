package model

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

// AsValid ...
func (i EntityIndex) AsValid() EntityIndex {
	if i.IsValid() {
		return i
	}
	return NoEntityIndex
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
	eiml      EIML
}

// NewEntityIndex ...
func NewEntityIndex(id string, isPrimary bool, c *C, eiml EIML) (*EI, error) {
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
