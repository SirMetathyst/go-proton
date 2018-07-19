package model

import "fmt"

// EIML ...
type EIML struct {
	l []*EIM
}

// NewEntityIndexMethodList ...
func NewEntityIndexMethodList() EIML {
	return EIML{}
}

// AddEntityIndexMethod ...
func (eiml *EIML) AddEntityIndexMethod(eim *EIM) {
	eiml.l = append(eiml.l, eim)
}

// EntityIndexMethodWithID ...
func (eiml *EIML) EntityIndexMethodWithID(id string) *EIM {
	for _, eim := range eiml.l {
		if eim.ID().EqualTo(id) {
			return eim
		}
	}
	return nil
}

// EntityIndexMethodList ...
func (eiml *EIML) EntityIndexMethodList() []*EIM {
	return eiml.l
}

// IsUnique ...
func (eiml *EIML) IsUnique() (String, error) {
	lookup := make(map[String]bool, 0)
	for _, eim := range eiml.l {
		if _, ok := lookup[eim.ID()]; ok {
			return eim.ID(), fmt.Errorf("IsUnique: Found duplicate of `%s`", eim.ID())
		}
		lookup[eim.ID()] = true
	}
	return "", nil
}
