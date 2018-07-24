package client

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
)

// ConfigurationChain ...
func ConfigurationChain(returnOnError bool, provider ...CP) CP {
	return func(bb *blackboard.BB) error {
		for _, p := range provider {
			err := p(bb)
			if returnOnError && err != nil {
				return err
			}
		}
		return nil
	}
}

// ModelChain ...
func ModelChain(provider ...MP) MP {
	return func(bb *blackboard.BB) (md *entitas.MD, err error) {
		for _, p := range provider {
			md, err = p(bb)
			if err == nil {
				return
			}
		}
		return
	}
}
