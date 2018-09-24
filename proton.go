package proton

import (
	"fmt"

	entitas "github.com/SirMetathyst/go-entitas"
)

var (
	ErrProtonModelUndefined    = fmt.Errorf("Proton: `Model` Undefined.")
	ErrProtonFileInfoUndefined = fmt.Errorf("Proton: `FileInfo` Undefined.")
)

// P ...
type P struct {
	generatorInfo     []*GI
	postProcessorInfo []*PPI
}

// NewProton ...
func NewProton() *P {
	return new(P)
}

var (
	p = NewProton()
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
	return p.generatorInfo
}

// PostProcessorInfo ...
func (p *P) PostProcessorInfo() []*PPI {
	return p.postProcessorInfo
}

// PostProcessorInfo ...
func PostProcessorInfo() []*PPI {
	return p.postProcessorInfo
}

// AddGenerator ...
func (p *P) AddGenerator(generatorVersion string, generator G, enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(generatorVersion, generator, enabled))
}

// AddGenerator ...
func AddGenerator(generatorVersion string, generator G, enabled bool) {
	p.AddGenerator(generatorVersion, generator, enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(postProcessorVersion, postProcessor, enabled))
}

// AddPostProcessor ...
func AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	p.AddPostProcessor(postProcessorVersion, postProcessor, enabled)
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
	p.EnableGenerator(generatorVersion, enabled)
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
	p.EnablePostProcessor(postProcessorVersion, enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(md *entitas.MD) ([]entitas.FI, error) {
	if md == nil {
		return nil, ErrProtonModelUndefined
	}
	r := make([]entitas.FI, 0)
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
func RunGenerator(md *entitas.MD) ([]entitas.FI, error) {
	return p.RunGenerator(md)
}

// RunPostProcessor ...
func (p *P) RunPostProcessor(fi []entitas.FI) ([]entitas.FI, error) {
	if fi == nil {
		return nil, ErrProtonFileInfoUndefined
	}
	r := fi
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
func RunPostProcessor(fi []entitas.FI) ([]entitas.FI, error) {
	return p.RunPostProcessor(fi)
}

// Run ...
func (p *P) Run(md *entitas.MD) ([]entitas.FI, error) {
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
func Run(md *entitas.MD) ([]entitas.FI, error) {
	return p.Run(md)
}
