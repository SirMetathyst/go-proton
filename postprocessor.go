package proton

import entitas "github.com/SirMetathyst/go-entitas"

// PP ...
type PP func([]entitas.FI) ([]entitas.FI, error)

// PPI ...
type PPI struct {
	PostProcessorVersion string
	PostProcessor        PP
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(postProcessorVersion string, postProcessor PP, enabled bool) *PPI {
	return &PPI{postProcessorVersion, postProcessor, enabled}
}
