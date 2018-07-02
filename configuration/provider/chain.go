package configurationprovider

import "github.com/SirMetathyst/proton/configuration"

// Chain ...
func Chain(returnOnError bool, provider ...func(*configuration.C) error) func(*configuration.C) error {
	return func(c *configuration.C) error {
		for _, p := range provider {
			err := p(c)
			if returnOnError && err != nil {
				return err
			}
		}
		return nil
	}
}
