package codegeneration

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	proton "github.com/SirMetathyst/go-proton"
	"github.com/fsnotify/fsnotify"
)

var (
	// ErrProtonModelIsNil is returned when the a nil model is given
	ErrProtonModelIsNil = fmt.Errorf("proton: model is nil")
	// ErrProtonFileInfoIsNil is returned when a nil FileInfo is given
	ErrProtonFileInfoIsNil = fmt.Errorf("proton: fileInfo is nil")
	// ErrNoProtonFilesFound is returned when no *.proton files were found
	ErrNoProtonFilesFound = fmt.Errorf("proton: no *.proton files were found")
)

// P ...
type P struct {
	generatorInfo     []*GI
	postProcessorInfo []*PPI
	Options           PO
}

// PO ...
type PO struct {
	ProjectPath  string
	OutputFolder string
}

// NewProton ...
func NewProton(options PO) *P {
	proton := new(P)
	proton.Options = options
	return proton
}

var (
	p = NewProton(PO{"./", "src-gen"})
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
	return p.Options.ProjectPath
}

// ProjectPath ...
func ProjectPath() string {
	return Singleton().ProjectPath()
}

// OutputFolder ...
func (p *P) OutputFolder() string {
	return p.Options.OutputFolder
}

// OutputFolder ...
func OutputFolder() string {
	return Singleton().OutputFolder()
}

// SetOptions ...
func (p *P) SetOptions(options PO) {
	p.Options.ProjectPath = options.ProjectPath
	outputFolder := filepath.Join(p.Options.ProjectPath, options.OutputFolder)
	p.Options.OutputFolder = outputFolder
}

// SetOptions ...
func SetOptions(options PO) {
	Singleton().SetOptions(options)
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

// RunGenerators ...
func (p *P) RunGenerators(md *proton.MD) ([]proton.FI, error) {
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

// RunGenerators ...
func RunGenerators(md *proton.MD) ([]proton.FI, error) {
	return Singleton().RunGenerators(md)
}

// RunPostProcessors ...
func (p *P) RunPostProcessors(fi []proton.FI) ([]proton.FI, error) {
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

// RunPostProcessors ...
func RunPostProcessors(fi []proton.FI) ([]proton.FI, error) {
	return Singleton().RunPostProcessors(fi)
}

// Generate ...
func (p *P) Generate(md *proton.MD) error {
	if md == nil {
		return ErrProtonModelIsNil
	}
	gv, err := p.RunGenerators(md)
	if err != nil {
		return err
	}
	_, err = p.RunPostProcessors(gv)
	if err != nil {
		return err
	}
	return nil
}

// Generate ...
func Generate(md *proton.MD) error {
	return Singleton().Generate(md)
}

// ParseGenerate ...
func (p *P) ParseGenerate(parser func(file string) (*proton.MD, error)) error {
	md, err := parser("proton.proton")
	if err != nil {
		return err
	}
	return p.Generate(md)
}

// ParseGenerate ...
func ParseGenerate(parser func(file string) (*proton.MD, error)) error {
	return Singleton().ParseGenerate(parser)
}

// ProtonFilesList ...
func ProtonFilesList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".proton" {
			log.Printf("proton: %q: FOUND", path)
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Daemon ...
func Daemon(parser func(file string) (*proton.MD, error)) error {
	err := ParseGenerate(parser)
	if err != nil {
		return err
	}
	w, err := fsnotify.NewWatcher()
	defer w.Close()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ProtonFilesList(ProjectPath())
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		return ErrNoProtonFilesFound
	}

	for _, file := range files {
		log.Printf("proton: %q: WATCH", file)
		err := w.Add(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	go func() {
		for {
			select {
			case ev := <-w.Events:
				{
					log.Printf("proton: %s\n", ev)
					err := ParseGenerate(parser)
					if err != nil {
						log.Printf("proton: %s\n", err)
					}
					log.Println("proton: generated")
				}
			}
		}
	}()

	Done := make(chan os.Signal)
	signal.Notify(Done, os.Interrupt, syscall.SIGTERM)
	<-Done
	return nil
}
