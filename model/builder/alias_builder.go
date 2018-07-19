package modelbuilder

import (
	"fmt"

	"github.com/SirMetathyst/proton/model"
)

var (
	ErrAliasBuilderAliasAlreadyBuilt = fmt.Errorf("AliasBuilder: Alias already built.")
)

// AB ...
type AB struct {
	outList   *model.AL
	built     bool
	id, value string
}

// NewAliasBuilder ...
func NewAliasBuilder(outList *model.AL) *AB {
	return &AB{outList: outList}
}

// SetID ...
func (ab *AB) SetID(id string) *AB {
	ab.id = id
	return ab
}

// SetValue ...
func (ab *AB) SetValue(value string) *AB {
	ab.value = value
	return ab
}

// Build ...
func (ab *AB) Build() error {
	if ab.built {
		return ErrAliasBuilderAliasAlreadyBuilt
	}
	a, err := model.NewAlias(ab.id, ab.value)
	if err != nil {
		return err
	}
	err = ab.outList.AddReplaceAlias(a)
	if err != nil {
		return err
	}
	return nil
}
