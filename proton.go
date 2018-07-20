package proton

import (
	"fmt"

	entitas "github.com/SirMetathyst/go-entitas"
)

var (
	ErrProtonModelUndefined    = fmt.Errorf("Proton: `Model` Undefined.")
	ErrProtonFileInfoUndefined = fmt.Errorf("Proton: `FileInfo` Undefined.")
)

// Generator ...
type Generator func(*entitas.MD) ([]FileInfo, error)

// PostProcessor ...
type PostProcessor func([]FileInfo) ([]FileInfo, error)

// GI ...
type GI struct {
	GeneratorVersion string
	Generator        Generator
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(generatorVersion string, generator Generator, enabled bool) *GI {
	return &GI{generatorVersion, generator, enabled}
}

// PI ...
type PI struct {
	PostProcessorVersion string
	PostProcessor        PostProcessor
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(postProcessorVersion string, postProcessor PostProcessor, enabled bool) *PI {
	return &PI{postProcessorVersion, postProcessor, enabled}
}

// FileInfo ...
type FileInfo struct {
	File        string
	FileContent string
	Generator   string
}

// NewFileInfo ...
func NewFileInfo(file, fileContent, generator string) FileInfo {
	return FileInfo{file, fileContent, generator}
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
func (p *P) AddGenerator(generatorVersion string, generator Generator, enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(generatorVersion, generator, enabled))
}

// AddGenerator ...
func AddGenerator(generatorVersion string, generator Generator, enabled bool) {
	proton.AddGenerator(generatorVersion, generator, enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(postProcessorVersion string, postProcessor PostProcessor, enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(postProcessorVersion, postProcessor, enabled))
}

// AddPostProcessor ...
func AddPostProcessor(postProcessorVersion string, postProcessor PostProcessor, enabled bool) {
	proton.AddPostProcessor(postProcessorVersion, postProcessor, enabled)
}

// EnableGenerator ...
func (p *P) EnableGenerator(generatorVersion string, enabled bool) {
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.GeneratorVersion == generatorVersion {
			generatorInfo.Enabled = enabled
		}
	}
}

// EnableGenerator ...
func EnableGenerator(generatorVersion string, enabled bool) {
	proton.EnableGenerator(generatorVersion, enabled)
}

// EnablePostProcessor ...
func (p *P) EnablePostProcessor(postProcessorVersion string, enabled bool) {
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.PostProcessorVersion == postProcessorVersion {
			postProcessorInfo.Enabled = enabled
		}
	}
}

// EnablePostProcessor ...
func EnablePostProcessor(postProcessorVersion string, enabled bool) {
	proton.EnablePostProcessor(postProcessorVersion, enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(md *entitas.MD) ([]FileInfo, error) {
	if md == nil {
		return nil, ErrProtonModelUndefined
	}
	r := make([]FileInfo, 0)
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.Enabled {
			gv, err := generatorInfo.Generator(md)
			if err != nil {
				return nil, err
			}
			r = append(r, gv...)
		}
	}
	return r, nil
}

// RunGenerator ...
func RunGenerator(md *entitas.MD) ([]FileInfo, error) {
	return proton.RunGenerator(md)
}

// RunPostProcessor ...
func (p *P) RunPostProcessor(fileInfo []FileInfo) ([]FileInfo, error) {
	if fileInfo == nil {
		return nil, ErrProtonFileInfoUndefined
	}
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
func (p *P) Run(md *entitas.MD) ([]FileInfo, error) {
	if md == nil {
		return nil, ErrProtonModelUndefined
	}
	gv, err := p.RunGenerator(md)
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
func Run(md *entitas.MD) ([]FileInfo, error) {
	return proton.Run(md)
}
