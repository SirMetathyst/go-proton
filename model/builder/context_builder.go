package modelbuilder

import (
	"fmt"

	"github.com/SirMetathyst/proton/model"
)

var (
	ErrContextBuilderContextAlreadyBuilt = fmt.Errorf("ContextBuilder: Context already built.")
)

// CB ...
type CB struct {
	outList *model.CL
	built   bool
	id      string
}

// NewContextBuilder ...
func NewContextBuilder(outList *model.CL) *CB {
	return &CB{outList: outList}
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
	c, err := model.NewContext(cb.id)
	if err != nil {
		return err
	}
	err = cb.outList.AddContext(c)
	if err != nil {
		return err
	}

	return nil
}
