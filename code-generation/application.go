package codegeneration

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	proton "github.com/SirMetathyst/go-proton"

	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/clean_target_directory/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/file_header/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/merge_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/write_to_disk/C_1_4_2"

	"github.com/fsnotify/fsnotify"
)

var (
	ProjectDirectory = flag.String("p", "./", "Project Directory")
	DaemonMode       = flag.Bool("d", false, "Daemon Mode")
)

var (
	// ErrNoProtonFilesFound is returned when no *.proton files were found
	ErrNoProtonFilesFound = fmt.Errorf("proton: no *.proton files were found")
)

// Must ...
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Usage ...
func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "[Flags]\n")
	c := 0
	flag.VisitAll(func(f *flag.Flag) {
		l := len(fmt.Sprintf("  --%s", f.Name))
		if l > c {
			c = l
		}
	})
	flag.VisitAll(func(f *flag.Flag) {
		s := fmt.Sprintf("  --%s", f.Name)
		name, _ := flag.UnquoteUsage(f)
		if len(name) > 0 {
			s += " " + name
		}
		diff := c - len(s)
		if diff < 0 {
			diff = 0
		}
		s += fmt.Sprintf(strings.Repeat(" ", diff)+"     (default %v)", f.DefValue)
		fmt.Fprint(flag.CommandLine.Output(), s, "\n")
	})
}

// SetupProton ...
func SetupProton() error {

	/* PostProcessor(s). */
	AddPostProcessor("MergeContentPostProcessor_C_1_4_2", MergeContentPostProcessor_C_1_4_2, true)
	AddPostProcessor("FileHeaderPostProcessor_C_1_4_2", FileHeaderPostProcessor_C_1_4_2, true)
	AddPostProcessor("PrintFilePostProcessor_C_1_4_2", PrintFilePostProcessor_C_1_4_2, false)
	AddPostProcessor("PrintFileContentPostProcessor_C_1_4_2", PrintFileContentPostProcessor_C_1_4_2, false)
	AddPostProcessor("CleanTargetDirectoryPostProcessor_C_1_4_2", CleanTargetDirectoryPostProcessor_C_1_4_2, true)
	AddPostProcessor("WriteToDiskPostProcessor_C_1_4_2", WriteToDiskPostProcessor_C_1_4_2, true)

	return nil
}

// SetupConfiguration ...
func SetupConfiguration() error {
	for _, generatorInfo := range GeneratorInfo() {
		flag.BoolVar(&generatorInfo.Enabled, generatorInfo.GeneratorVersion, generatorInfo.Enabled, "")
	}
	for _, postProcessorInfo := range PostProcessorInfo() {
		flag.BoolVar(&postProcessorInfo.Enabled, postProcessorInfo.PostProcessorVersion, postProcessorInfo.Enabled, "")
	}
	return nil
}

// Setup ...
func Setup() error {
	err := SetupProton()
	if err != nil {
		return err
	}
	err = SetupConfiguration()
	if err != nil {
		return err
	}
	flag.Usage = Usage
	flag.Parse()
	return nil
}

// RunProton ...
func RunProton(parser func(file string) (*proton.MD, error)) error {
	md, err := parser("proton.proton")
	if err != nil {
		return err
	}
	_, err = Run(md)
	if err != nil {
		return err
	}
	log.Println("proton: generated")
	return nil
}

// ProtonFiles ...
func ProtonFiles(root string) ([]string, error) {
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
	w, err := fsnotify.NewWatcher()
	defer w.Close()
	Must(err)

	files, err := ProtonFiles(*ProjectDirectory)
	Must(err)

	if len(files) == 0 {
		return ErrNoProtonFilesFound
	}

	for _, file := range files {
		log.Printf("proton: %q: WATCH", file)
		Must(w.Add(file))
	}

	go func() {
		for {
			select {
			case ev := <-w.Events:
				{
					log.Printf("proton: %s\n", ev)
					RunProton(parser)
				}
			}
		}
	}()

	Done := make(chan os.Signal)
	signal.Notify(Done, os.Interrupt, syscall.SIGTERM)
	<-Done
	return nil
}

// RunApplication ...
func RunApplication(parser func(file string) (*proton.MD, error)) {

	Must(Setup())
	Must(RunProton(parser))

	if *DaemonMode == true {
		Must(Daemon(parser))
		return
	}
}
