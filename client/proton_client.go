package client

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-proton"

	"github.com/fsnotify/fsnotify"

	. "github.com/SirMetathyst/go-proton/postprocessor/clean_target_directory/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/file_header/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/merge_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/print_file/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/print_file_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/write_to_disk/C_1_4_2"
)

// Must ...
func Must(err error) {
	if err != nil {
		log.Fatalf("Electron: %s", err)
	}
}

// SetupProton ...
func SetupProton(p *proton.P) error {

	CSharpEntitasGenerator(p)

	/* PostProcessor(s). */
	p.AddPostProcessor("MergeContentPostProcessor_C_1_4_2", MergeContentPostProcessor_C_1_4_2, true)
	p.AddPostProcessor("FileHeaderPostProcessor_C_1_4_2", FileHeaderPostProcessor_C_1_4_2, true)
	p.AddPostProcessor("PrintFilePostProcessor_C_1_4_2", PrintFilePostProcessor_C_1_4_2, true)
	p.AddPostProcessor("PrintFileContentPostProcessor_C_1_4_2", PrintFileContentPostProcessor_C_1_4_2, false)
	p.AddPostProcessor("CleanTargetDirectoryPostProcessor_C_1_4_2", CleanTargetDirectoryPostProcessor_C_1_4_2, true)
	p.AddPostProcessor("WriteToDiskPostProcessor_C_1_4_2", WriteToDiskPostProcessor_C_1_4_2, true)

	return nil
}

// SetupConfiguration ...
func SetupConfiguration(bb *blackboard.BB, p *proton.P) error {

	for _, generatorInfo := range p.GeneratorInfo() {
		bb.SetBoolP(generatorInfo.GeneratorVersion, &generatorInfo.Enabled)
	}

	for _, postProcessorInfo := range p.PostProcessorInfo() {
		bb.SetBoolP(postProcessorInfo.PostProcessorVersion, &postProcessorInfo.Enabled)
	}

	CleanTargetDirectory = OutputDirectory(bb)
	WriteToDiskDirectory = OutputDirectory(bb)

	return nil
}

// Setup ...
func Setup(bb *blackboard.BB, p *proton.P) error {
	Must(SetupBlackboard(bb))
	Must(SetupProton(p))
	Must(SetupConfiguration(bb, p))
	Must(Flag(bb))
	return nil
}

// Run ...
func Run(bb *blackboard.BB, p *proton.P) error {
	md, err := JSON(bb)
	if err != nil {
		return err
	}
	_, err = p.Run(md)
	return err
}

// Watcher ...
func Watcher(bb *blackboard.BB, p *proton.P, File ...string) error {
	w, err := fsnotify.NewWatcher()
	defer w.Close()
	Must(err)

	for _, File := range File {
		Must(w.Add(File))
	}

	go func() {
		for {
			select {
			case <-w.Events:
				{
					err := Run(bb, p)
					if err != nil {
						log.Println(err)
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

// PC ...
type PC struct {
}

// NewProtonClient ...
func NewProtonClient() *PC {
	return &PC{}
}

var (
	pc = NewProtonClient()
)

// Run ...
func (pc *PC) Run() {
	bb := blackboard.NewBlackboard()
	p := proton.NewProton()

	Must(Setup(bb, p))

	if WatchFileEnable(bb) && len(WatchFile(bb)) >= 1 {
		err := Run(bb, p)
		if err != nil {
			log.Println(err)
		}
		Must(Watcher(bb, p, WatchFile(bb)...))
		return
	}
	err := Run(bb, p)
	if err != nil {
		log.Println(err)
	}
}
