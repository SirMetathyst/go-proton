package proton

import "errors"

var (
	// ErrEntityIndexListTriedToAddNilEntityIndex ...
	ErrEntityIndexListTriedToAddNilEntityIndex = errors.New("proton: entity index list: tried to add nil entity index")
	// ErrEntityIndexListTriedToAddDuplicateEntityIndexID ...
	ErrEntityIndexListTriedToAddDuplicateEntityIndexID = errors.New("proton: entity index list: tried to add entity index with duplicate id")
)

// EntityIndexList ...
type EntityIndexList struct {
	entityIndexSlice []*EntityIndex
}

// NewEntityIndexList ...
func NewEntityIndexList() *EntityIndexList {
	return &EntityIndexList{}
}

// AddEntityIndex ...
func (eil *EntityIndexList) AddEntityIndex(entityIndex *EntityIndex) error {
	if entityIndex == nil {
		panic(ErrEntityIndexListTriedToAddNilEntityIndex)
	}
	if eil.EntityIndexWithID(entityIndex.ID().String()) != nil {
		return ErrEntityIndexListTriedToAddDuplicateEntityIndexID
	}
	eil.entityIndexSlice = append(eil.entityIndexSlice, entityIndex)
	return nil
}

// EntityIndexWithID ...
func (eil *EntityIndexList) EntityIndexWithID(id string) *EntityIndex {
	for _, entityIndex := range eil.entityIndexSlice {
		if entityIndex.ID().EqualTo(id) {
			return entityIndex
		}
	}
	return nil
}

// EntityIndexList ...
func (eil *EntityIndexList) EntityIndexSlice() []*EntityIndex {
	return eil.entityIndexSlice
}
