package proton

import (
	"github.com/SirMetathyst/proton/model"
)

// Generator ...
type Generator func(*model.M) ([]FileInfo, error)

// PostProcessor ...
type PostProcessor func([]FileInfo) ([]FileInfo, error)

// GeneratorInfo ...
type GeneratorInfo struct {
	GeneratorVersion string
	Generator        Generator
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(GeneratorVersion string, Generator Generator, Enabled bool) *GeneratorInfo {
	return &GeneratorInfo{GeneratorVersion, Generator, Enabled}
}

// PostProcessorInfo ...
type PostProcessorInfo struct {
	PostProcessorVersion string
	PostProcessor        PostProcessor
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) *PostProcessorInfo {
	return &PostProcessorInfo{PostProcessorVersion, PostProcessor, Enabled}
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
	generatorInfo     []*GeneratorInfo
	postProcessorInfo []*PostProcessorInfo
}

// NewProton ...
func NewProton() *P {
	return new(P)
}

var (
	Proton = NewProton()
)

// GetGeneratorInfo ...
func (p *P) GetGeneratorInfo() []*GeneratorInfo {
	return p.generatorInfo
}

// GetGeneratorInfo ...
func GetGeneratorInfo() []*GeneratorInfo {
	return Proton.generatorInfo
}

// GetPostProcessorInfo ...
func (p *P) GetPostProcessorInfo() []*PostProcessorInfo {
	return p.postProcessorInfo
}

// GetPostProcessorInfo ...
func GetPostProcessorInfo() []*PostProcessorInfo {
	return Proton.postProcessorInfo
}

// AddGenerator ...
func (p *P) AddGenerator(GeneratorVersion string, Generator Generator, Enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(GeneratorVersion, Generator, Enabled))
}

// AddGenerator ...
func AddGenerator(GeneratorVersion string, Generator Generator, Enabled bool) {
	Proton.AddGenerator(GeneratorVersion, Generator, Enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(PostProcessorVersion, PostProcessor, Enabled))
}

// AddPostProcessor ...
func AddPostProcessor(PostProcessorVersion string, PostProcessor PostProcessor, Enabled bool) {
	Proton.AddPostProcessor(PostProcessorVersion, PostProcessor, Enabled)
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
	Proton.EnableGenerator(GeneratorVersion, Enabled)
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
	Proton.EnablePostProcessor(PostProcessorVersion, Enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(m *model.M) ([]FileInfo, error) {
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
func RunGenerator(m *model.M) ([]FileInfo, error) {
	return Proton.RunGenerator(m)
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
	return Proton.RunPostProcessor(fileInfo)
}

// Run ...
func (p *P) Run(m *model.M) ([]FileInfo, error) {
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
func Run(m *model.M) ([]FileInfo, error) {
	return Proton.Run(m)
}
