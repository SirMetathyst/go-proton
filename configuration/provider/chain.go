package configurationprovider

import "github.com/SirMetathyst/go-blackboard"

// Chain ...
func Chain(returnOnError bool, provider ...func(*blackboard.BB) error) func(*blackboard.BB) error {
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
