package modelprovider

import (
	"github.com/SirMetathyst/proton/configuration"
	"github.com/SirMetathyst/proton/model"
)

// Chain ...
func Chain(provider ...func(*configuration.C) (*model.M, error)) func(*configuration.C) (*model.M, error) {
	return func(c *configuration.C) (m *model.M, err error) {
		for _, p := range provider {
			m, err = p(c)
			if err == nil {
				return
			}
		}
		return
	}
}
