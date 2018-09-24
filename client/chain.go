package main

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
)

// CP ...
type CP func(*blackboard.BB) error

// MP ...
type MP func(*blackboard.BB) (*entitas.MD, error)

// ConfigurationFallback ...
func ConfigurationFallback(returnOnError bool, provider ...CP) CP {
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

// ModelFallback ...
func ModelFallback(logger func(error), provider ...MP) MP {
	return func(bb *blackboard.BB) (md *entitas.MD, err error) {
		for _, p := range provider {
			md, err = p(bb)
			if err == nil {
				return md, nil
			}
			logger(err)
		}
		return
	}
}
