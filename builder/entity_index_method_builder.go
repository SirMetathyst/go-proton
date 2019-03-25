package builder

import (
	"fmt"

	"github.com/SirMetathyst/go-proton"
)

var (
	ErrEntityIndexMethodBuilderAliasListShouldNotBeNil             = fmt.Errorf("EntityIndexMethodBuilder: `AliasList` should not be nil.")
	ErrEntityIndexMethodBuilderEntityIndexMethodListShouldNotBeNil = fmt.Errorf("EntityIndexMethodBuilder: `EntityIndexMethodList` should not be nil.")
	ErrEntityIndexMethodBuilderEntityIndexMethodAlreadyBuilt       = fmt.Errorf("EntityIndexMethodBuilder: `EntityIndexMethod` already built.")
)

// EIMB ...
type EIMB struct {
	al     *proton.AL
	teimml *proton.EIMML
	eiml   *proton.EIML
	id     string
	built  bool
}

// NewEntityIndexMethodBuilder ...
func NewEntityIndexMethodBuilder(al *proton.AL, eiml *proton.EIML) *EIMB {
	if al == nil {
		panic(ErrEntityIndexMethodBuilderAliasListShouldNotBeNil)
	}
	if eiml == nil {
		panic(ErrEntityIndexMethodBuilderEntityIndexMethodListShouldNotBeNil)
	}
	return &EIMB{
		al:     al,
		eiml:   eiml,
		teimml: proton.NewEntityIndexMethodMemberList(),
	}
}

// SetID ...
func (eimb *EIMB) SetID(id string) *EIMB {
	eimb.id = id
	return eimb
}

// NewMember ...
func (eimb *EIMB) NewMember() *EIMMB {
	return NewEntityIndexMethodMemberBuilder(eimb.al, eimb.teimml)
}

// Build ...
func (eimb *EIMB) Build() error {
	if eimb.built {
		return ErrEntityIndexMethodBuilderEntityIndexMethodAlreadyBuilt
	}
	eim, err := proton.NewEntityIndexMethod(eimb.id, eimb.teimml)
	if err != nil {
		return err
	}
	err = eimb.eiml.AddEntityIndexMethod(eim)
	if err != nil {
		return err
	}
	eimb.built = true
	return nil
}
