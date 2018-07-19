package modelbuilder

import "github.com/SirMetathyst/proton/model"

// MDB ...
type MDB struct {
	ns        string
	ctxList   model.CL
	dctx      string
	aliasList model.AL
	cpList    model.CPL
	eiList    model.EIL
}

// NewModelBuilder ...
func NewModelBuilder() *MDB {
	return &MDB{}
}

// SetNamespace ...
func (mdb *MDB) SetNamespace(ns string) *MDB {
	mdb.ns = ns
	return mdb
}

// NewContext ...
func (mdb *MDB) NewContext() *CB {
	return NewContextBuilder(&mdb.ctxList)
}

// SetDefaultContext ...
func (mdb *MDB) SetDefaultContext(id string) *MDB {
	mdb.dctx = id
	return mdb
}

// NewAlias ...
func (mdb *MDB) NewAlias() *AB {
	return NewAliasBuilder(&mdb.aliasList)
}

// NewComponent ...
func (mdb *MDB) NewComponent() *CPB {
	return NewComponentBuilder(&mdb.cpList, &mdb.ctxList, mdb.dctx, &mdb.aliasList)
}

// Reset ...
func (mdb *MDB) Reset() *MDB {
	mdb.ns = ""
	mdb.dctx = ""
	mdb.ctxList = model.NewContextList()
	mdb.aliasList = model.NewAliasList()
	mdb.cpList = model.NewComponentList()
	mdb.eiList = model.NewEntityIndexList()
	return mdb
}

// Build ...
func (mdb *MDB) Build() (*model.MD, error) {
	md, err := model.NewModel(mdb.ns, mdb.ctxList, mdb.aliasList, mdb.cpList, mdb.eiList)
	mdb.Reset()
	return md, err
}
