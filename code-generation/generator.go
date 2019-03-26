package codegeneration

import (
	"fmt"
	"path/filepath"

	proton "github.com/SirMetathyst/go-proton"
)

var (
	// ErrProtonModelIsNil is returned when the a nil model is given
	ErrProtonModelIsNil = fmt.Errorf("proton: model is nil")
	// ErrProtonFileInfoIsNil is returned when a nil FileInfo is given
	ErrProtonFileInfoIsNil = fmt.Errorf("proton: fileInfo is nil")
)

// P ...
type P struct {
	generatorInfo     []*GI
	postProcessorInfo []*PPI
	projectPath       string
	outputFolder      string
}

// NewProton ...
func NewProton(projectPath, outputFolder string) *P {
	proton := new(P)
	proton.outputFolder = outputFolder
	proton.SetProjectPath(projectPath)
	return proton
}

var (
	p = NewProton("./", "src-gen")
)

// Singleton ...
func Singleton() *P {
	return p
}

// GeneratorInfo ...
func (p *P) GeneratorInfo() []*GI {
	return p.generatorInfo
}

// GeneratorInfo ...
func GeneratorInfo() []*GI {
	return Singleton().GeneratorInfo()
}

// PostProcessorInfo ...
func (p *P) PostProcessorInfo() []*PPI {
	return p.postProcessorInfo
}

// PostProcessorInfo ...
func PostProcessorInfo() []*PPI {
	return Singleton().PostProcessorInfo()
}

// ProjectPath ...
func (p *P) ProjectPath() string {
	return p.projectPath
}

// ProjectPath ...
func ProjectPath() string {
	return Singleton().ProjectPath()
}

// OutputFolder ...
func (p *P) OutputFolder() string {
	return p.outputFolder
}

// OutputFolder ...
func OutputFolder() string {
	return Singleton().OutputFolder()
}

// SetProjectPath ...
func (p *P) SetProjectPath(path string) {
	p.projectPath = path
	outputFolder := filepath.Join(p.projectPath, p.outputFolder)
	p.outputFolder = outputFolder
}

// SetProjectPath ...
func SetProjectPath(path string) {
	Singleton().SetProjectPath(path)
}

// SetOutputFolder ...
func (p *P) SetOutputFolder(folder string) {
	outputFolder := filepath.Join(p.projectPath, folder)
	p.outputFolder = outputFolder
}

// SetOutputFolder ...
func SetOutputFolder(folder string) {
	Singleton().SetOutputFolder(folder)
}

// AddGenerator ...
func (p *P) AddGenerator(generatorVersion string, generator G, enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(generatorVersion, generator, enabled))
}

// AddGenerator ...
func AddGenerator(generatorVersion string, generator G, enabled bool) {
	Singleton().AddGenerator(generatorVersion, generator, enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(postProcessorVersion, postProcessor, enabled))
}

// AddPostProcessor ...
func AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	Singleton().AddPostProcessor(postProcessorVersion, postProcessor, enabled)
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
	Singleton().EnableGenerator(generatorVersion, enabled)
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
	Singleton().EnablePostProcessor(postProcessorVersion, enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(md *proton.MD) ([]proton.FI, error) {
	if md == nil {
		return nil, ErrProtonModelIsNil
	}
	r := make([]proton.FI, 0)
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
func RunGenerator(md *proton.MD) ([]proton.FI, error) {
	return Singleton().RunGenerator(md)
}

// RunPostProcessor ...
func (p *P) RunPostProcessor(fi []proton.FI) ([]proton.FI, error) {
	if fi == nil {
		return nil, ErrProtonFileInfoIsNil
	}
	r := fi
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.Enabled {
			pv, err := postProcessorInfo.PostProcessor(p, r)
			if err != nil {
				return nil, err
			}
			r = pv
		}
	}
	return r, nil
}

// RunPostProcessor ...
func RunPostProcessor(fi []proton.FI) ([]proton.FI, error) {
	return Singleton().RunPostProcessor(fi)
}

// Run ...
func (p *P) Run(md *proton.MD) ([]proton.FI, error) {
	if md == nil {
		return nil, ErrProtonModelIsNil
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
func Run(md *proton.MD) ([]proton.FI, error) {
	return Singleton().Run(md)
}
