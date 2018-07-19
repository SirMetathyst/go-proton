package model

import (
	"github.com/SirMetathyst/proton/configuration"
)

// Provider ...
type Provider func(*configuration.C) (*MD, error)

// MD ...
type MD struct {
	ns        string
	ctxList   CL
	aliasList AL
	cpList    CPL
	eiList    EIL
}

// NewModel ...
func NewModel(ns string, ctxList CL, aliasList AL, cpList CPL, eiList EIL) (*MD, error) {
	return &MD{ns, ctxList, aliasList, cpList, eiList}, nil
}

// Namespace ...
func (md *MD) Namespace() String {
	return String(md.ns)
}

// ContextWithID ...
func (md *MD) ContextWithID(id string) *C {
	return md.ctxList.ContextWithID(id)
}

// ContextList ...
func (md *MD) ContextList() []*C {
	return md.ctxList.ContextList()
}

// AliasWithID ...
func (md *MD) AliasWithID(id string) *A {
	return md.aliasList.AliasWithID(id)
}

// AliasList ...
func (md *MD) AliasList() []*A {
	return md.aliasList.AliasList()
}

// ComponentsWithContextID ...
func (md *MD) ComponentsWithContextID(id string) []*CP {
	return md.cpList.ComponentsWithContextID(id)
}

// ComponentWithID ...
func (md *MD) ComponentWithID(id string) *CP {
	return md.cpList.ComponentWithID(id)
}

// ComponentsWithEntityIndex ...
func (md *MD) ComponentsWithEntityIndex() []*CP {
	return md.cpList.ComponentsWithEntityIndex()
}

// ComponentList ...
func (md *MD) ComponentList() []*CP {
	return md.cpList.ComponentList()
}

// EntityIndexList ...
func (md *MD) EntityIndexList() []*EI {
	return md.eiList.EntityIndexList()
}
