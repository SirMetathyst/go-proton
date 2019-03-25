package proton

import "fmt"

var (
	ErrContextListTriedToAddNilContext         = fmt.Errorf("entitas(context list): tried to add nil context")
	ErrContextListTriedToAddDuplicateContextID = fmt.Errorf("entitas(context list): tried to add context with duplicate id")
)

// CL ...
type CL struct {
	l []*C
}

// NewContextList ...
func NewContextList() *CL {
	return &CL{}
}

// AddContext ...
func (cl *CL) AddContext(c *C) error {
	if c == nil {
		return ErrContextListTriedToAddNilContext
	}
	if cl.ContextWithID(c.ID().String()) != nil {
		return ErrContextListTriedToAddDuplicateContextID
	}

	cl.l = append(cl.l, c)
	return nil
}

// ContextWithID ...
func (cl *CL) ContextWithID(id string) *C {
	for _, c := range cl.l {
		if c.ID().EqualTo(id) {
			return c
		}
	}
	return nil
}

// ContextSlice ...
func (cl *CL) ContextSlice() []*C {
	return cl.l
}
