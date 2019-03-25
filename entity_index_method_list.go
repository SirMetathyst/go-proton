package proton

import "fmt"

var (
	ErrEntityIndexMethodListTriedToAddNilEntityIndexMethod         = fmt.Errorf("EntityIndexMethodList: Tried to add nil `EntityIndexMethod`")
	ErrEntityIndexMethodListTriedToAddDuplicateEntityIndexMethodID = fmt.Errorf("EntityIndexMethodList: Tried to add `EntityIndexMethod` with duplicate `ID`")
)

// EIML ...
type EIML struct {
	l []*EIM
}

// NewEntityIndexMethodList ...
func NewEntityIndexMethodList() *EIML {
	return &EIML{}
}

// AddEntityIndexMethod ...
func (eiml *EIML) AddEntityIndexMethod(eim *EIM) error {
	if eim == nil {
		return ErrEntityIndexMethodListTriedToAddNilEntityIndexMethod
	}
	//if eiml.EntityIndexMethodWithID(eim.ID().String()) != nil {
	//	return ErrEntityIndexMethodListTriedToAddDuplicateEntityIndexMethodID
	//}
	eiml.l = append(eiml.l, eim)
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

// EntityIndexMethodList ...
func (eiml *EIML) EntityIndexMethodList() []*EIM {
	return eiml.l
}
