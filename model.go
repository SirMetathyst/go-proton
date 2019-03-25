package proton

// MD ...
type MD struct {
	target, ns string
	ctxList    *CL
	aliasList  *AL
	cpList     *CPL
	eiList     *EIL
}

// NewModel ...
func NewModel(target, ns string, ctxList *CL, aliasList *AL, cpList *CPL, eiList *EIL) (*MD, error) {
	return &MD{target, ns, ctxList, aliasList, cpList, eiList}, nil
}

// Target ...
func (md *MD) Target() String {
	return String(md.target)
}

// Namespace ...
func (md *MD) Namespace() String {
	return String(md.ns)
}

// ContextWithID ...
func (md *MD) ContextWithID(id string) *C {
	return md.ctxList.ContextWithID(id)
}

// ContextSlice ...
func (md *MD) ContextSlice() []*C {
	return md.ctxList.ContextSlice()
}

// AliasWithID ...
func (md *MD) AliasWithID(id string) *A {
	return md.aliasList.AliasWithID(id)
}

// AliasSlice ...
func (md *MD) AliasSlice() []*A {
	return md.aliasList.AliasSlice()
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

// ComponentSlice ...
func (md *MD) ComponentSlice() []*CP {
	return md.cpList.ComponentSlice()
}

// EntityIndexList ...
func (md *MD) EntityIndexList() []*EI {
	return md.eiList.EntityIndexList()
}
