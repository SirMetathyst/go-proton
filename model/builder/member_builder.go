package modelbuilder

import (
	"fmt"

	"github.com/SirMetathyst/proton/model"
)

var (
	ErrMemberBuilderMemberAlreadyBuilt = fmt.Errorf("MemberBuilder: Member already built.")
)

// MB ...
type MB struct {
	outList     *model.ML
	aliasList   *model.AL
	built       bool
	id, value   string
	entityIndex model.EntityIndex
	alias       *model.A
}

// NewMemberBuilder ...
func NewMemberBuilder(outList *model.ML, aliasList *model.AL) *MB {
	return &MB{outList: outList, aliasList: aliasList}
}

// SetID ...
func (mb *MB) SetID(id string) *MB {
	mb.id = id
	return mb
}

// SetValue ...
func (mb *MB) SetValue(value string) *MB {
	mb.value = value
	return mb
}

// SetEntityIndex ...
func (mb *MB) SetEntityIndex(value model.EntityIndex) *MB {
	mb.entityIndex = value
	return mb
}

// SetAlias ...
func (mb *MB) SetAlias(id string) *MB {
	mb.alias = mb.aliasList.AliasWithID(id)
	return mb
}

// Build ...
func (mb *MB) Build() error {
	if mb.built {
		return ErrMemberBuilderMemberAlreadyBuilt
	}
	if mb.alias != nil {
		m, err := model.NewMemberAlias(mb.id, mb.alias, mb.entityIndex)
		if err != nil {
			return err
		}
		mb.outList.AddMember(m)
		return nil
	}
	m, err := model.NewMember(mb.id, mb.value, mb.entityIndex)
	if err != nil {
		return err
	}
	err = mb.outList.AddMember(m)
	if err != nil {
		return err
	}
	return nil
}
