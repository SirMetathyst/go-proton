package builder

import (
	"fmt"

	"github.com/SirMetathyst/go-proton"
)

var (
	ErrEntityIndexBuilderContextListShouldNotBeNil     = fmt.Errorf("EntityIndexBuilder: `ContextList` should not be nil.")
	ErrEntityIndexBuilderAliasListShouldNotBeNil       = fmt.Errorf("EntityIndexBuilder: `AliasList` should not be nil.")
	ErrEntityIndexBuilderEntityIndexListShouldNotBeNil = fmt.Errorf("EntityIndexBuilder: `EntityIndexList` should not be nil.")
	ErrEntityIndexBuilderEntityIndexAlreadyBuilt       = fmt.Errorf("EntityIndexBuilder: `EntityIndex` already built.")
	ErrEntityIndexBuilderEntityIndexMustHaveContext    = fmt.Errorf("EntityIndexBuilder: `EntityIndex` does not have a `Context`.")
	ErrEntityIndexBuilderEntityIndexMustHaveMethod     = fmt.Errorf("EntityIndexBuilder: `EntityIndex` does not have a `Method`.")
)

// EIB ...
type EIB struct {
	cl        *proton.CL
	al        *proton.AL
	eil       *proton.EIL
	ctx       *proton.C
	teiml     *proton.EIML
	id        string
	isPrimary bool
	built     bool
}

// NewEntityIndexBuilder ...
func NewEntityIndexBuilder(cl *proton.CL, al *proton.AL, eil *proton.EIL) *EIB {
	if cl == nil {
		panic(ErrEntityIndexBuilderContextListShouldNotBeNil)
	}
	if al == nil {
		panic(ErrEntityIndexBuilderAliasListShouldNotBeNil)
	}
	if eil == nil {
		panic(ErrEntityIndexBuilderEntityIndexListShouldNotBeNil)
	}
	return &EIB{
		cl:    cl,
		al:    al,
		eil:   eil,
		teiml: proton.NewEntityIndexMethodList(),
	}
}

// SetID ...
func (eib *EIB) SetID(id string) *EIB {
	eib.id = id
	return eib
}

// SetPrimary ...
func (eib *EIB) SetPrimary(isPrimary bool) *EIB {
	eib.isPrimary = isPrimary
	return eib
}

// AddContext ...
func (eib *EIB) AddContext(id string) *EIB {
	ctx := eib.cl.ContextWithID(id)
	eib.ctx = ctx
	return eib
}

// NewMethod ...
func (eib *EIB) NewMethod() *EIMB {
	return NewEntityIndexMethodBuilder(eib.al, eib.teiml)
}

// Build ...
func (eib *EIB) Build() error {
	if eib.built {
		return ErrEntityIndexBuilderEntityIndexAlreadyBuilt
	}
	if eib.ctx == nil {
		return ErrEntityIndexBuilderEntityIndexMustHaveContext
	}
	if len(eib.teiml.EntityIndexMethodList()) == 0 {
		return ErrEntityIndexBuilderEntityIndexMustHaveMethod
	}
	ei, err := proton.NewEntityIndex(eib.id, eib.isPrimary, eib.ctx, eib.teiml)
	if err != nil {
		return err
	}
	err = eib.eil.AddEntityIndex(ei)
	if err != nil {
		return err
	}
	eib.built = true
	return nil
}
