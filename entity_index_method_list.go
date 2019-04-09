package proton

import "errors"

var (
	// ErrEntityIndexMethodListTriedToAddNilEntityIndexMethod ...
	ErrEntityIndexMethodListTriedToAddNilEntityIndexMethod = errors.New("proton: entity index method list: tried to add nil entity index method")
	// ErrEntityIndexMethodListTriedToAddDuplicateEntityIndexMethodID ...
	ErrEntityIndexMethodListTriedToAddDuplicateEntityIndexMethodID = errors.New("proton: entity index method list: tried to add entity index method with duplicate id")
)

// EntityIndexMethodList ...
type EntityIndexMethodList struct {
	entityIndexMethodSlice []*EntityIndexMethod
}

// NewEntityIndexMethodList ...
func NewEntityIndexMethodList() *EntityIndexMethodList {
	return &EntityIndexMethodList{}
}

// AddEntityIndexMethod ...
func (eiml *EntityIndexMethodList) AddEntityIndexMethod(entityIndexMethod *EntityIndexMethod) error {
	if entityIndexMethod == nil {
		panic(ErrEntityIndexMethodListTriedToAddNilEntityIndexMethod)
	}
	//if eiml.EntityIndexMethodWithID(eim.ID().String()) != nil {
	//	return ErrEntityIndexMethodListTriedToAddDuplicateEntityIndexMethodID
	//}
	eiml.entityIndexMethodSlice = append(eiml.entityIndexMethodSlice, entityIndexMethod)
	return nil
}

// EntityIndexMethodWithID ...
//func (eiml *EIML) EntityIndexMethodWithID(id string) *EIM {
//	for _, eim := range eiml.l {
//		if eim.ID().EqualTo(id) {
//			return eim
//		}
//	}
//	return nil
//}

// EntityIndexMethodSlice ...
func (eiml *EntityIndexMethodList) EntityIndexMethodSlice() []*EntityIndexMethod {
	return eiml.entityIndexMethodSlice
}
