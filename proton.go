package proton

import (
	"fmt"

	"github.com/SirMetathyst/proton/model"
)

var (
	ErrProtonModelUndefined = fmt.Errorf("Proton: Model Undefined.")
)

// Generator ...
type Generator func(*model.MD) ([]FileInfo, error)

// PostProcessor ...
type PostProcessor func([]FileInfo) ([]FileInfo, error)

// GI ...
type GI struct {
	GeneratorVersion string
	Generator        Generator
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(GeneratorVersion string, Generator Generator, Enabled bool) *GI {
	return &GI{GeneratorVersion, Generator, Enabled}
}

// PI ...
type PI struct {
	PostProcessorVersion string
	PostProcessor        PostProcessor
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) *PI {
	return &PI{PostProcessorVersion, PostProcessor, Enabled}
}

// FileInfo ...
type FileInfo struct {
	File        string
	FileContent string
	Generator   string
}

// NewFileInfo ...
func NewFileInfo(File, FileContent, Generator string) FileInfo {
	return FileInfo{File, FileContent, Generator}
}

// P ...
type P struct {
	generatorInfo     []*GI
	postProcessorInfo []*PI
}

// NewProton ...
func NewProton() *P {
	return new(P)
}

var (
	proton = NewProton()
)

// Singleton ...
func Singleton() *P {
	return proton
}

// GeneratorInfo ...
func (p *P) GeneratorInfo() []*GI {
	return p.generatorInfo
}

// GeneratorInfo ...
func GeneratorInfo() []*GI {
	return proton.generatorInfo
}

// PostProcessorInfo ...
func (p *P) PostProcessorInfo() []*PI {
	return p.postProcessorInfo
}

// PostProcessorInfo ...
func PostProcessorInfo() []*PI {
	return proton.postProcessorInfo
}

// AddGenerator ...
func (p *P) AddGenerator(GeneratorVersion string, Generator Generator, Enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(GeneratorVersion, Generator, Enabled))
}

// AddGenerator ...
func AddGenerator(GeneratorVersion string, Generator Generator, Enabled bool) {
	proton.AddGenerator(GeneratorVersion, Generator, Enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(PostProcessorVersion, PostProcessor, Enabled))
}

// AddPostProcessor ...
func AddPostProcessor(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) {
	proton.AddPostProcessor(PostProcessorVersion, PostProcessor, Enabled)
}

// EnableGenerator ...
func (p *P) EnableGenerator(GeneratorVersion string, Enabled bool) {
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.GeneratorVersion == GeneratorVersion {
			generatorInfo.Enabled = Enabled
		}
	}
}

// EnableGenerator ...
func EnableGenerator(GeneratorVersion string, Enabled bool) {
	proton.EnableGenerator(GeneratorVersion, Enabled)
}

// EnablePostProcessor ...
func (p *P) EnablePostProcessor(PostProcessorVersion string, Enabled bool) {
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.PostProcessorVersion == PostProcessorVersion {
			postProcessorInfo.Enabled = Enabled
		}
	}
}

// EnablePostProcessor ...
func EnablePostProcessor(PostProcessorVersion string, Enabled bool) {
	proton.EnablePostProcessor(PostProcessorVersion, Enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(m *model.MD) ([]FileInfo, error) {
	r := make([]FileInfo, 0)
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.Enabled {
			gv, err := generatorInfo.Generator(m)
			if err != nil {
				return nil, err
			}
			r = append(r, gv...)
		}
	}
	return r, nil
}

// RunGenerator ...
func RunGenerator(m *model.MD) ([]FileInfo, error) {
	return proton.RunGenerator(m)
}

// RunPostProcessor ...
func (p *P) RunPostProcessor(fileInfo []FileInfo) ([]FileInfo, error) {
	r := fileInfo
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.Enabled {
			pv, err := postProcessorInfo.PostProcessor(r)
			if err != nil {
				return nil, err
			}
			r = pv
		}
	}
	return r, nil
}

// RunPostProcessor ...
func RunPostProcessor(fileInfo []FileInfo) ([]FileInfo, error) {
	return proton.RunPostProcessor(fileInfo)
}

// Run ...
func (p *P) Run(m *model.MD) ([]FileInfo, error) {
	if m == nil {
		return nil, ErrProtonModelUndefined
	}

	gv, err := p.RunGenerator(m)
	if err != nil {
		return nil, err
	}
	pv, err := p.RunPostProcessor(gv)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

// Run ...
func Run(m *model.MD) ([]FileInfo, error) {
	return proton.Run(m)
}
