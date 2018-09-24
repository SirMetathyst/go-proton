package proton

import (
	"github.com/SirMetathyst/go-entitas"
)

// G ...
type G func(md *entitas.MD) ([]entitas.FI, error)

// GI ...
type GI struct {
	GeneratorVersion string
	Generator        G
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(generatorVersion string, generator G, enabled bool) *GI {
	return &GI{generatorVersion, generator, enabled}
}
