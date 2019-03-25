package proton

import (
	"fmt"
)

var (
	ErrEntityIndexMethodMemberBuilderMemberAlreadyBuilt   = fmt.Errorf("proton(entity index method member builder): entity index method member already built")
	ErrEntityIndexMethodMemberBuilderAliasShouldNotBeNil  = fmt.Errorf("proton(entity index method member builder): alias list should not be nil")
	ErrEntityIndexMethodMemberBuilderMemberShouldNotBeNil = fmt.Errorf("proton(entity index method member builder): entity index method member list should not be nil")
)

// EIMMB ...
type EIMMB struct {
	al        *AL
	eimml     *EIMML
	a         *A
	built     bool
	id, value string
}

// NewEntityIndexMethodMemberBuilder ...
func NewEntityIndexMethodMemberBuilder(al *AL, eimml *EIMML) *EIMMB {
	if al == nil {
		panic(ErrEntityIndexMethodMemberBuilderAliasShouldNotBeNil)
	}
	if eimml == nil {
		panic(ErrEntityIndexMethodMemberBuilderMemberShouldNotBeNil)
	}
	return &EIMMB{al: al, eimml: eimml}
}

// SetID ...
func (eimmb *EIMMB) SetID(id string) *EIMMB {
	eimmb.id = id
	return eimmb
}

// SetValue ...
func (eimmb *EIMMB) SetValue(value string) *EIMMB {
	eimmb.value = value
	return eimmb
}

// SetAlias ...
func (eimmb *EIMMB) SetAlias(id string) *EIMMB {
	eimmb.a = eimmb.al.AliasWithID(id)
	return eimmb
}

// Build ...
func (eimmb *EIMMB) Build() error {
	if eimmb.built {
		return ErrEntityIndexMethodMemberBuilderMemberAlreadyBuilt
	}
	if eimmb.a != nil {
		m, err := NewEntityIndexMethodMemberAlias(eimmb.id, eimmb.a)
		if err != nil {
			return err
		}
		eimmb.eimml.AddMember(m)
		return nil
	}
	m, err := NewEntityIndexMethodMember(eimmb.id, eimmb.value)
	if err != nil {
		return err
	}
	err = eimmb.eimml.AddMember(m)
	if err != nil {
		return err
	}
	eimmb.built = true
	return nil
}
