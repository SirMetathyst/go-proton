package builder

import (
	"fmt"

	"github.com/SirMetathyst/go-proton"
)

var (
	ErrAliasBuilderAliasListShouldNotBeNil = fmt.Errorf("AliasBuilder: `AliasList` should not be nil.")
	ErrAliasBuilderAliasAlreadyBuilt       = fmt.Errorf("AliasBuilder: `Alias` already built.")
)

// AB ...
type AB struct {
	al        *proton.AL
	built     bool
	id, value string
}

// NewAliasBuilder ...
func NewAliasBuilder(al *proton.AL) *AB {
	if al == nil {
		panic(ErrAliasBuilderAliasListShouldNotBeNil)
	}
	return &AB{al: al}
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
	a, err := proton.NewAlias(ab.id, ab.value)
	if err != nil {
		return err
	}
	err = ab.al.AddAlias(a)
	if err != nil {
		return err
	}
	ab.built = true
	return nil
}
