package proton

import "errors"

var (
	// ErrAliasBuilderAliasListShouldNotBeNil ...
	ErrAliasBuilderAliasListShouldNotBeNil = errors.New("proton: alias builder: alias list should not be nil")
	// ErrAliasBuilderAliasAlreadyBuilt ...
	ErrAliasBuilderAliasAlreadyBuilt = errors.New("proton: alias builder: alias is already built")
)

// AliasBuilder ...
type AliasBuilder struct {
	aliasList *AliasList
	built     bool
	id, value string
}

// NewAliasBuilder ...
func NewAliasBuilder(aliasList *AliasList) *AliasBuilder {
	if aliasList == nil {
		panic(ErrAliasBuilderAliasListShouldNotBeNil)
	}
	return &AliasBuilder{aliasList: aliasList}
}

// SetID ...
func (ab *AliasBuilder) SetID(id string) *AliasBuilder {
	ab.id = id
	return ab
}

// SetValue ...
func (ab *AliasBuilder) SetValue(value string) *AliasBuilder {
	ab.value = value
	return ab
}

// Build ...
func (ab *AliasBuilder) Build() error {
	if ab.built {
		return ErrAliasBuilderAliasAlreadyBuilt
	}
	a, err := NewAlias(ab.id, ab.value)
	if err != nil {
		return err
	}
	err = ab.aliasList.AddAlias(a)
	if err != nil {
		return err
	}
	ab.built = true
	return nil
}
