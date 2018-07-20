package configuration

import "github.com/SirMetathyst/go-blackboard"

// Provider ...
type Provider func(*blackboard.BB) error
