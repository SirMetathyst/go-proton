package builder

import (
	"fmt"

	"github.com/SirMetathyst/go-proton"
)

var (
	ErrComponentMemberBuilderMemberAlreadyBuilt   = fmt.Errorf("proton(component member builder): component member already built")
	ErrComponentMemberBuilderAliasShouldNotBeNil  = fmt.Errorf("proton(component member builder): alias list should not be nil")
	ErrComponentMemberBuilderMemberShouldNotBeNil = fmt.Errorf("proton(component member builder): component member list should not be nil")
)

// CPMB ...
type CPMB struct {
	al          *proton.AL
	cml         *proton.CML
	a           *proton.A
	built       bool
	id, value   string
	entityIndex proton.EntityIndex
}

// NewComponentMemberBuilder ...
func NewComponentMemberBuilder(al *proton.AL, cml *proton.CML) *CPMB {
	if al == nil {
		panic(ErrComponentMemberBuilderAliasShouldNotBeNil)
	}
	if cml == nil {
		panic(ErrComponentMemberBuilderMemberShouldNotBeNil)
	}
	return &CPMB{al: al, cml: cml}
}

// SetID ...
func (cpmb *CPMB) SetID(id string) *CPMB {
	cpmb.id = id
	return cpmb
}

// SetValue ...
func (cpmb *CPMB) SetValue(value string) *CPMB {
	cpmb.value = value
	return cpmb
}

// SetEntityIndex ...
func (cpmb *CPMB) SetEntityIndex(value proton.EntityIndex) *CPMB {
	cpmb.entityIndex = value
	return cpmb
}

// SetAlias ...
func (cpmb *CPMB) SetAlias(id string) *CPMB {
	cpmb.a = cpmb.al.AliasWithID(id)
	return cpmb
}

// Build ...
func (cpmb *CPMB) Build() error {
	if cpmb.built {
		return ErrComponentMemberBuilderMemberAlreadyBuilt
	}
	if cpmb.a != nil {
		m, err := proton.NewComponentMemberAlias(cpmb.id, cpmb.a, cpmb.entityIndex)
		if err != nil {
			return err
		}
		cpmb.cml.AddMember(m)
		return nil
	}
	m, err := proton.NewComponentMember(cpmb.id, cpmb.value, cpmb.entityIndex)
	if err != nil {
		return err
	}
	err = cpmb.cml.AddMember(m)
	if err != nil {
		return err
	}
	cpmb.built = true
	return nil
}
