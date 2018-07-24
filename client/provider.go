package client

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
)

// CP ...
type CP func(*blackboard.BB) error

// MP ...
type MP func(*blackboard.BB) (*entitas.MD, error)
