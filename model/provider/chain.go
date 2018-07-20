package modelprovider

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
)

// Chain ...
func Chain(provider ...func(*blackboard.BB) (*entitas.MD, error)) func(*blackboard.BB) (*entitas.MD, error) {
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
