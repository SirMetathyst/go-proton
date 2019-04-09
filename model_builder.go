package proton

import "errors"

var (
	// ErrModelBuilderModelAlreadyBuilt ...
	ErrModelBuilderModelAlreadyBuilt = errors.New("proton: model builder: model is already built")
)

// ModelBuilder ...
type ModelBuilder struct {
	target          string
	namespace       string
	contextList     *ContextList
	aliasList       *AliasList
	componentList   *ComponentList
	entityIndexList *EntityIndexList
	defaultContext  string
	built           bool
}

// NewModelBuilder ...
func NewModelBuilder() *ModelBuilder {
	return &ModelBuilder{
		target:          "",
		namespace:       "",
		contextList:     NewContextList(),
		aliasList:       NewAliasList(),
		componentList:   NewComponentList(),
		entityIndexList: NewEntityIndexList(),
		defaultContext:  "",
		built:           false,
	}
}

// SetTarget ...
func (mdb *ModelBuilder) SetTarget(target string) *ModelBuilder {
	mdb.target = target
	return mdb
}

// SetNamespace ...
func (mdb *ModelBuilder) SetNamespace(namespace string) *ModelBuilder {
	mdb.namespace = namespace
	return mdb
}

// NewContext ...
func (mdb *ModelBuilder) NewContext() *ContextBuilder {
	return NewContextBuilder(mdb.contextList)
}

// SetDefaultContext ...
func (mdb *ModelBuilder) SetDefaultContext(id string) *ModelBuilder {
	mdb.defaultContext = id
	return mdb
}

// NewAlias ...
func (mdb *ModelBuilder) NewAlias() *AliasBuilder {
	return NewAliasBuilder(mdb.aliasList)
}

// NewComponent ...
func (mdb *ModelBuilder) NewComponent() *ComponentBuilder {
	return NewComponentBuilder(mdb.contextList, mdb.aliasList, mdb.componentList, mdb.defaultContext)
}

// NewEntityIndex ...
func (mdb *ModelBuilder) NewEntityIndex() *EntityIndexBuilder {
	return NewEntityIndexBuilder(mdb.contextList, mdb.aliasList, mdb.entityIndexList)
}

// Reset ...
func (mdb *ModelBuilder) Reset() *ModelBuilder {
	mdb.target = ""
	mdb.namespace = ""
	mdb.defaultContext = ""
	mdb.built = false
	mdb.contextList = NewContextList()
	mdb.aliasList = NewAliasList()
	mdb.componentList = NewComponentList()
	mdb.entityIndexList = NewEntityIndexList()
	return mdb
}

// Build ...
func (mdb *ModelBuilder) Build() (*Model, error) {
	if mdb.built {
		return nil, ErrModelBuilderModelAlreadyBuilt
	}
	md, err := NewModel(
		mdb.target,
		mdb.namespace,
		mdb.contextList,
		mdb.aliasList,
		mdb.componentList,
		mdb.entityIndexList,
	)
	mdb.built = true
	return md, err
}
