package proton

import "errors"

var (
	// ErrContextBuilderContextListShouldNotBeNil ...
	ErrContextBuilderContextListShouldNotBeNil = errors.New("proton: context builder: context list should not be nil")
	// ErrContextBuilderContextAlreadyBuilt ...
	ErrContextBuilderContextAlreadyBuilt = errors.New("proton: context builder: context is already built")
)

// ContextBuilder ...
type ContextBuilder struct {
	contextList *ContextList
	built       bool
	id          string
}

// NewContextBuilder ...
func NewContextBuilder(contextList *ContextList) *ContextBuilder {
	if contextList == nil {
		panic(ErrContextBuilderContextListShouldNotBeNil)
	}
	return &ContextBuilder{contextList: contextList}
}

// SetID ...
func (cb *ContextBuilder) SetID(id string) *ContextBuilder {
	cb.id = id
	return cb
}

// Build ...
func (cb *ContextBuilder) Build() error {
	if cb.built {
		return ErrContextBuilderContextAlreadyBuilt
	}
	context, err := NewContext(cb.id)
	if err != nil {
		return err
	}
	err = cb.contextList.AddContext(context)
	if err != nil {
		return err
	}
	cb.built = true
	return nil
}
