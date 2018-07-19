package model

// EIM ...
type EIM struct {
	id string
	ml ML
}

// NewEntityIndexMethod ...
func NewEntityIndexMethod(id string, ml ML) *EIM {
	return &EIM{id, ml}
}

// ID ...
func (eim *EIM) ID() String {
	return String(eim.id)
}

// MemberList ...
func (eim *EIM) MemberList() []*M {
	return eim.ml.MemberList()
}
