package proton

import (
	"fmt"
)

var (
	ErrContextBuilderContextListShouldNotBeNil = fmt.Errorf("ContextBuilder: `ContextList` should not be nil.")
	ErrContextBuilderContextAlreadyBuilt       = fmt.Errorf("ContextBuilder: `Context` already built.")
)

// CB ...
type CB struct {
	cl    *CL
	built bool
	id    string
}

// NewContextBuilder ...
func NewContextBuilder(cl *CL) *CB {
	if cl == nil {
		panic(ErrContextBuilderContextListShouldNotBeNil)
	}
	return &CB{cl: cl}
}

// SetID ...
func (cb *CB) SetID(id string) *CB {
	cb.id = id
	return cb
}

// Build ...
func (cb *CB) Build() error {
	if cb.built {
		return ErrContextBuilderContextAlreadyBuilt
	}
	c, err := NewContext(cb.id)
	if err != nil {
		return err
	}
	err = cb.cl.AddContext(c)
	if err != nil {
		return err
	}
	cb.built = true
	return nil
}
