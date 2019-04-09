package proton

import "errors"

var (
	// ErrContextListTriedToAddNilContext ...
	ErrContextListTriedToAddNilContext = errors.New("proton: context list: tried to add nil context")
	// ErrContextListTriedToAddDuplicateContextID ...
	ErrContextListTriedToAddDuplicateContextID = errors.New("proton: context list: tried to add context with duplicate id")
)

// ContextList ...
type ContextList struct {
	contextSlice []*Context
}

// NewContextList ...
func NewContextList() *ContextList {
	return &ContextList{}
}

// AddContext ...
func (cl *ContextList) AddContext(context *Context) error {
	if context == nil {
		panic(ErrContextListTriedToAddNilContext)
	}
	if cl.ContextWithID(context.ID().String()) != nil {
		return ErrContextListTriedToAddDuplicateContextID
	}

	cl.contextSlice = append(cl.contextSlice, context)
	return nil
}

// ContextWithID ...
func (cl *ContextList) ContextWithID(id string) *Context {
	for _, context := range cl.contextSlice {
		if context.ID().EqualTo(id) {
			return context
		}
	}
	return nil
}

// ContextSlice ...
func (cl *ContextList) ContextSlice() []*Context {
	return cl.contextSlice
}
