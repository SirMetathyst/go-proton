package proton

import "errors"

var (
	// ErrComponentListTriedToAddNilComponent ...
	ErrComponentListTriedToAddNilComponent = errors.New("proton: component list: tried to add nil component")
	// ErrComponentListTriedToAddDuplicateComponentID ...
	ErrComponentListTriedToAddDuplicateComponentID = errors.New("proton: component list: tried to add component with duplicate id")
)

// ComponentList ...
type ComponentList struct {
	componentSlice []*Component
}

// NewComponentList ...
func NewComponentList() *ComponentList {
	return &ComponentList{}
}

// AddComponent ...
func (cpl *ComponentList) AddComponent(component *Component) error {
	if component == nil {
		return ErrComponentListTriedToAddNilComponent
	}
	if cpl.ComponentWithID(component.ID().String()) != nil {
		return ErrComponentListTriedToAddDuplicateComponentID
	}
	cpl.componentSlice = append(cpl.componentSlice, component)
	return nil
}

// ComponentWithID ...
func (cpl *ComponentList) ComponentWithID(id string) *Component {
	for _, component := range cpl.componentSlice {
		if component.ID().EqualTo(id) {
			return component
		}
	}
	return nil
}

// ComponentsWithContextID ...
func (cpl *ComponentList) ComponentsWithContextID(id string) []*Component {
	slice := make([]*Component, 0)
	for _, component := range cpl.componentSlice {
		context := component.ContextWithID(id)
		if context != nil {
			slice = append(slice, component)
		}
	}
	return slice
}

// ComponentsWithEntityIndex ...
func (cpl *ComponentList) ComponentsWithEntityIndex() []*Component {
	slice := make([]*Component, 0)
	for _, component := range cpl.componentSlice {
		memberSlice := component.MembersWithEntityIndex()
		if len(memberSlice) > 0 {
			slice = append(slice, component)
		}
	}
	return slice
}

// ComponentSlice ...
func (cpl *ComponentList) ComponentSlice() []*Component {
	return cpl.componentSlice
}
