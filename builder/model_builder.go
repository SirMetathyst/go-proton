package builder

import (
	"fmt"

	"github.com/SirMetathyst/go-proton"
)

var (
	ErrModelBuilderModelAlreadyBuilt = fmt.Errorf("ModelBuilder: `Model` already built.")
)

// MDB ...
type MDB struct {
	target, ns string
	cl         *proton.CL
	al         *proton.AL
	cpl        *proton.CPL
	eil        *proton.EIL
	dctx       string
	built      bool
}

// NewModelBuilder ...
func NewModelBuilder() *MDB {
	return &MDB{
		cl:  proton.NewContextList(),
		al:  proton.NewAliasList(),
		cpl: proton.NewComponentList(),
		eil: proton.NewEntityIndexList(),
	}
}

// SetTarget ...
func (mdb *MDB) SetTarget(v string) *MDB {
	mdb.target = v
	return mdb
}

// SetNamespace ...
func (mdb *MDB) SetNamespace(v string) *MDB {
	mdb.ns = v
	return mdb
}

// NewContext ...
func (mdb *MDB) NewContext() *CB {
	return NewContextBuilder(mdb.cl)
}

// SetDefaultContext ...
func (mdb *MDB) SetDefaultContext(id string) *MDB {
	mdb.dctx = id
	return mdb
}

// NewAlias ...
func (mdb *MDB) NewAlias() *AB {
	return NewAliasBuilder(mdb.al)
}

// NewComponent ...
func (mdb *MDB) NewComponent() *CPB {
	return NewComponentBuilder(mdb.cl, mdb.al, mdb.cpl, mdb.dctx)
}

// NewEntityIndex ...
func (mdb *MDB) NewEntityIndex() *EIB {
	return NewEntityIndexBuilder(mdb.cl, mdb.al, mdb.eil)
}

// Reset ...
func (mdb *MDB) Reset() *MDB {
	mdb.target = ""
	mdb.ns = ""
	mdb.dctx = ""
	mdb.built = false
	mdb.cl = proton.NewContextList()
	mdb.al = proton.NewAliasList()
	mdb.cpl = proton.NewComponentList()
	mdb.eil = proton.NewEntityIndexList()
	return mdb
}

// Build ...
func (mdb *MDB) Build() (*proton.MD, error) {
	if mdb.built {
		return nil, ErrModelBuilderModelAlreadyBuilt
	}
	md, err := proton.NewModel(mdb.target, mdb.ns, mdb.cl, mdb.al, mdb.cpl, mdb.eil)
	mdb.built = true
	return md, err
}
