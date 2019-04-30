package codegeneration

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
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
	generatorInfoSlice     []*GeneratorInfo
	postProcessorInfoSlice []*PostProcessorInfo
	Options                PO
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

// GeneratorInfoSlice ...
func (p *P) GeneratorInfoSlice() []*GeneratorInfo {
	return p.generatorInfoSlice
}

// GeneratorInfoSlice ...
func GeneratorInfoSlice() []*GeneratorInfo {
	return Singleton().GeneratorInfoSlice()
}

// PostProcessorInfoSlice ...
func (p *P) PostProcessorInfoSlice() []*PostProcessorInfo {
	return p.postProcessorInfoSlice
}

// PostProcessorInfoSlice ...
func PostProcessorInfoSlice() []*PostProcessorInfo {
	return Singleton().PostProcessorInfoSlice()
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
func (p *P) AddGenerator(generatorVersion string, generatorFunc GeneratorFunc, enabled bool) {
	p.generatorInfoSlice = append(p.generatorInfoSlice, NewGeneratorInfo(generatorVersion, generatorFunc, enabled))
}

// AddGenerator ...
func AddGenerator(generatorVersion string, generatorFunc GeneratorFunc, enabled bool) {
	Singleton().AddGenerator(generatorVersion, generatorFunc, enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(postProcessorVersion string, postProcessorFunc PostProcessorFunc, enabled bool) {
	p.postProcessorInfoSlice = append(p.postProcessorInfoSlice, NewPostProcessorInfo(postProcessorVersion, postProcessorFunc, enabled))
}

// AddPostProcessor ...
func AddPostProcessor(postProcessorVersion string, postProcessorFunc PostProcessorFunc, enabled bool) {
	Singleton().AddPostProcessor(postProcessorVersion, postProcessorFunc, enabled)
}

// EnableGenerator ...
func (p *P) EnableGenerator(generatorVersion string, enabled bool) {
	for _, generatorInfo := range p.generatorInfoSlice {
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
	for _, postProcessorInfo := range p.postProcessorInfoSlice {
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
func (p *P) RunGenerators(md *proton.Model) ([]proton.FileInfo, error) {
	if md == nil {
		return nil, ErrProtonModelIsNil
	}
	r := make([]proton.FileInfo, 0)
	for _, generatorInfo := range p.generatorInfoSlice {
		if generatorInfo.Enabled {
			gv, err := generatorInfo.GeneratorFunc(md)
			if err != nil {
				return nil, err
			}
			r = append(r, gv...)
		}
	}
	return r, nil
}

// RunGenerators ...
func RunGenerators(md *proton.Model) ([]proton.FileInfo, error) {
	return Singleton().RunGenerators(md)
}

// RunPostProcessors ...
func (p *P) RunPostProcessors(fi []proton.FileInfo) ([]proton.FileInfo, error) {
	if fi == nil {
		return nil, ErrProtonFileInfoIsNil
	}
	r := fi
	for _, postProcessorInfo := range p.postProcessorInfoSlice {
		if postProcessorInfo.Enabled {
			pv, err := postProcessorInfo.PostProcessorFunc(p, r)
			if err != nil {
				return nil, err
			}
			r = pv
		}
	}
	return r, nil
}

// RunPostProcessors ...
func RunPostProcessors(fi []proton.FileInfo) ([]proton.FileInfo, error) {
	return Singleton().RunPostProcessors(fi)
}

// Generate ...
func (p *P) Generate(md *proton.Model) error {
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
func Generate(md *proton.Model) error {
	return Singleton().Generate(md)
}

// ParseGenerate ...
func (p *P) ParseGenerate(parser func(file string) (*proton.Model, error)) error {
	file := path.Join(p.Options.ProjectPath, "proton.proton")
	md, err := parser(file)
	if err != nil {
		return err
	}
	return p.Generate(md)
}

// ParseGenerate ...
func ParseGenerate(parser func(file string) (*proton.Model, error)) error {
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
func Daemon(parser func(file string) (*proton.Model, error)) error {
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

	err = ParseGenerate(parser)
	if err != nil {
		log.Printf("proton: %s\n", err)
	} else {
		log.Println("proton: generated")
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
					} else {
						log.Println("proton: generated")
					}
				}
			}
		}
	}()

	Done := make(chan os.Signal)
	signal.Notify(Done, os.Interrupt, syscall.SIGTERM)
	<-Done
	return nil
}
