package modelprovider

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/proton/model"
)

// Chain ...
func Chain(provider ...func(*blackboard.BB) (*model.MD, error)) func(*blackboard.BB) (*model.MD, error) {
	return func(bb *blackboard.BB) (md *model.MD, err error) {
		for _, p := range provider {
			md, err = p(bb)
			if err == nil {
				return
			}
		}
		return
	}
}
