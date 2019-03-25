package proton

import (
	"fmt"
)

var (
	ErrModelBuilderModelAlreadyBuilt = fmt.Errorf("ModelBuilder: `Model` already built.")
)

// MDB ...
type MDB struct {
	target, ns string
	cl         *CL
	al         *AL
	cpl        *CPL
	eil        *EIL
	dctx       string
	built      bool
}

// NewModelBuilder ...
func NewModelBuilder() *MDB {
	return &MDB{
		cl:  NewContextList(),
		al:  NewAliasList(),
		cpl: NewComponentList(),
		eil: NewEntityIndexList(),
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
	mdb.cl = NewContextList()
	mdb.al = NewAliasList()
	mdb.cpl = NewComponentList()
	mdb.eil = NewEntityIndexList()
	return mdb
}

// Build ...
func (mdb *MDB) Build() (*MD, error) {
	if mdb.built {
		return nil, ErrModelBuilderModelAlreadyBuilt
	}
	md, err := NewModel(mdb.target, mdb.ns, mdb.cl, mdb.al, mdb.cpl, mdb.eil)
	mdb.built = true
	return md, err
}
