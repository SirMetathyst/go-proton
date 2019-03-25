package codegeneration

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/SirMetathyst/go-blackboard"

	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/clean_target_directory/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/file_header/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/merge_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/print_file_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/code-generation/postprocessor/write_to_disk/C_1_4_2"
	"github.com/SirMetathyst/go-proton/dsl"
	"github.com/fsnotify/fsnotify"
)

var (
	File            = flag.String("File", "proton.proton", "")
	OutputDirectory = flag.String("OutputDirectory", "Generated", "")
	WatchFileEnable = flag.Bool("WatchFileEnable", false, "")
	WatchFile       = []string{}
)

// Must ...
func Must(err error) {
	if err != nil {
		log.Fatalf("Proton: %s", err)
	}
}

// StringSlice ...
type StringSlice []string

// String ...
func (s *StringSlice) String() string {
	return fmt.Sprint(*s)
}

// Set ...
func (s *StringSlice) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("string slice flag already set")
	}
	for _, v := range strings.Split(value, ",") {
		*s = append(*s, v)
	}
	return nil
}

// NewStringSlice ...
func NewStringSlice(value []string, p *[]string) *StringSlice {
	*p = value
	return (*StringSlice)(p)
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

// Flag ...
func Flag() error {

	for _, opt := range blackboard.AllBool() {
		flag.BoolVar(opt.Value, opt.Key, *opt.Value, "")
	}

	for _, opt := range blackboard.AllStringSlice() {
		flag.Var(NewStringSlice(*opt.Value, opt.Value), opt.Key, "")
	}

	for _, opt := range blackboard.AllString() {
		flag.StringVar(opt.Value, opt.Key, *opt.Value, "")
	}

	flag.Usage = Usage
	flag.Parse()
	return nil
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
		blackboard.SetBoolP(generatorInfo.GeneratorVersion, &generatorInfo.Enabled)
	}

	for _, postProcessorInfo := range PostProcessorInfo() {
		blackboard.SetBoolP(postProcessorInfo.PostProcessorVersion, &postProcessorInfo.Enabled)
	}

	CleanTargetDirectory = *OutputDirectory
	WriteToDiskDirectory = *OutputDirectory

	return nil
}

// Setup ...
func Setup() error {
	Must(SetupProton())
	Must(SetupConfiguration())
	Must(Flag())
	return nil
}

// RunProton ...
func RunProton() error {
	md, err := dsl.Parse(*File)
	if err != nil {
		return err
	}
	_, err = Run(md)
	return err
}

// Watcher ...
func Watcher(File ...string) error {
	w, err := fsnotify.NewWatcher()
	defer w.Close()
	Must(err)

	for _, File := range File {
		Must(w.Add(File))
	}

	go func() {
		for {
			select {
			case ev := <-w.Events:
				{
					log.Printf("Proton: %s\n", ev)
					err := RunProton()
					if err != nil {
						log.Printf("Proton: %s\n", err)
					} else {
						log.Printf("Proton: Generated\n")
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

// RunApplication ...
func RunApplication() {
	flag.Var(NewStringSlice(WatchFile, &WatchFile), "WatchFile", "")
	Must(Setup())

	err := RunProton()
	if err != nil {
		log.Printf("Proton: %s", err)
	} else {
		log.Printf("Proton: Generated")
	}

	if *WatchFileEnable && len(WatchFile) >= 1 {
		Must(Watcher(WatchFile...))
		return
	}
}
