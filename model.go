package proton

// Model ...
type Model struct {
	target          string
	namespace       string
	contextList     *ContextList
	aliasList       *AliasList
	componentList   *ComponentList
	entityIndexList *EntityIndexList
}

// NewModel ...
func NewModel(target string,
	namespace string,
	contextList *ContextList,
	aliasList *AliasList,
	componentList *ComponentList,
	entityIndexList *EntityIndexList) (*Model, error) {

	return &Model{
		target:          target,
		namespace:       namespace,
		contextList:     contextList,
		aliasList:       aliasList,
		componentList:   componentList,
		entityIndexList: entityIndexList}, nil
}

// Target ...
func (md *Model) Target() String {
	return String(md.target)
}

// Namespace ...
func (md *Model) Namespace() String {
	return String(md.namespace)
}

// ContextWithID ...
func (md *Model) ContextWithID(id string) *Context {
	return md.contextList.ContextWithID(id)
}

// ContextSlice ...
func (md *Model) ContextSlice() []*Context {
	return md.contextList.ContextSlice()
}

// AliasWithID ...
func (md *Model) AliasWithID(id string) *Alias {
	return md.aliasList.AliasWithID(id)
}

// AliasSlice ...
func (md *Model) AliasSlice() []*Alias {
	return md.aliasList.AliasSlice()
}

// ComponentsWithContextID ...
func (md *Model) ComponentsWithContextID(id string) []*Component {
	return md.componentList.ComponentsWithContextID(id)
}

// ComponentWithID ...
func (md *Model) ComponentWithID(id string) *Component {
	return md.componentList.ComponentWithID(id)
}

// ComponentsWithEntityIndex ...
func (md *Model) ComponentsWithEntityIndex() []*Component {
	return md.componentList.ComponentsWithEntityIndex()
}

// ComponentSlice ...
func (md *Model) ComponentSlice() []*Component {
	return md.componentList.ComponentSlice()
}

// EntityIndexSlice ...
func (md *Model) EntityIndexSlice() []*EntityIndex {
	return md.entityIndexList.EntityIndexSlice()
}
