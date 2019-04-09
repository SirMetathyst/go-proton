package codegeneration

import proton "github.com/SirMetathyst/go-proton"

// GeneratorFunc ...
type GeneratorFunc func(md *proton.Model) ([]proton.FileInfo, error)

// GeneratorInfo ...
type GeneratorInfo struct {
	GeneratorVersion string
	GeneratorFunc    GeneratorFunc
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(generatorVersion string, generatorFunc GeneratorFunc, enabled bool) *GeneratorInfo {
	return &GeneratorInfo{generatorVersion, generatorFunc, enabled}
}

// PostProcessorFunc ...
type PostProcessorFunc func(*P, []proton.FileInfo) ([]proton.FileInfo, error)

// PostProcessorInfo ...
type PostProcessorInfo struct {
	PostProcessorVersion string
	PostProcessorFunc    PostProcessorFunc
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(postProcessorVersion string, postProcessorFunc PostProcessorFunc, enabled bool) *PostProcessorInfo {
	return &PostProcessorInfo{postProcessorVersion, postProcessorFunc, enabled}
}
