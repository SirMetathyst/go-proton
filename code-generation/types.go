package codegeneration

import proton "github.com/SirMetathyst/go-proton"

// G ...
type G func(md *proton.MD) ([]proton.FI, error)

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

// PP ...
type PP func([]proton.FI) ([]proton.FI, error)

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
