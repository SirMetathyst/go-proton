package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/SirMetathyst/go-blackboard"
	entitas "github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-proton"
	protonlang "github.com/SirMetathyst/go-proton-lang"

	. "github.com/SirMetathyst/go-proton/generator/component/E_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_context/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_entity/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_entity_interface/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_lookup/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/component_matcher/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/composite_system/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context/C_1_9_0"
	. "github.com/SirMetathyst/go-proton/generator/context_matcher/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_matcher/E_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/context_observer/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/contexts/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/entity/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/entity_index/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/generator/event_listener_component_entity/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_listener_interface/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_system/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/event_systems/C_1_6_1"
	. "github.com/SirMetathyst/go-proton/generator/feature/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/clean_target_directory/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/file_header/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/merge_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/print_file/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/print_file_content/C_1_4_2"
	. "github.com/SirMetathyst/go-proton/postprocessor/write_to_disk/C_1_4_2"
	"github.com/fsnotify/fsnotify"
)

// Must ...
func Must(err error) {
	if err != nil {
		log.Fatalf("Proton: %s", err)
	}
}

// SetupProton ...
func SetupProton(p *proton.P) error {

	/* Generator(s).*/
	p.AddGenerator("CSharpComponentGenerator_E_1_4_2", ComponentGenerator_E_1_4_2, true)
	p.AddGenerator("CSharpComponentContextGenerator_C_1_4_2", ComponentContextGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpComponentEntityGenerator_C_1_4_2", ComponentEntityGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpComponentEntityInterfaceGenerator_C_1_4_2", ComponentEntityInterfaceGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpComponentLookupGenerator_C_1_4_2", ComponentLookupGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpComponentMatcherGenerator_C_1_4_2", ComponentMatcherGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpContextGenerator_C_1_4_2", ContextGenerator_C_1_4_2, false)
	p.AddGenerator("CSharpContextGenerator_C_1_9_0", ContextGenerator_C_1_9_0, true)
	p.AddGenerator("CSharpContextMatcherGenerator_C_1_4_2", ContextMatcherGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpContextMatcherGenerator_E_1_4_2", ContextMatcherGenerator_E_1_4_2, false)
	p.AddGenerator("CSharpContextsGenerator_C_1_4_2", ContextsGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpContextObserverGenerator_C_1_4_2", ContextObserverGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpEntityGenerator_C_1_4_2", EntityGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpEntityIndexGenerator_C_1_4_2", EntityIndexGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpFeatureGenerator_C_1_4_2", FeatureGenerator_C_1_4_2, true)
	p.AddGenerator("CSharpEventListenerComponentEntityGenerator_C_1_6_1", EventListenerComponentEntityGenerator_C_1_6_1, true)
	p.AddGenerator("CSharpEventListenerInterfaceGenerator_C_1_6_1", EventListenerInterfaceGenerator_C_1_6_1, true)
	p.AddGenerator("CSharpEventSystemGenerator_C_1_6_1", EventSystemGenerator_C_1_6_1, true)
	p.AddGenerator("CSharpEventSystemsGenerator_C_1_6_1", EventSystemsGenerator_C_1_6_1, true)
	p.AddGenerator("CSharpCompositeSystemGenerator_C_1_6_1", CompositeSystemGenerator_C_1_4_2, false)

	/* PostProcessor(s). */
	p.AddPostProcessor("MergeContentPostProcessor_C_1_4_2", MergeContentPostProcessor_C_1_4_2, true)
	p.AddPostProcessor("FileHeaderPostProcessor_C_1_4_2", FileHeaderPostProcessor_C_1_4_2, true)
	p.AddPostProcessor("PrintFilePostProcessor_C_1_4_2", PrintFilePostProcessor_C_1_4_2, false)
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
	providers := make(map[string]func(string) (*entitas.MD, error))
	providers[".json"] = JSON
	providers[".proton"] = protonlang.Parse

	f := filepath.Ext(File(bb))
	if tp, ok := providers[f]; ok {
		md, err := tp(File(bb))
		if err != nil {
			return err
		}
		_, err = p.Run(md)
		return err
	}

	return errors.New("No provider for " + File(bb))
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
			case ev := <-w.Events:
				{
					log.Printf("Proton: %s\n", ev)
					err := Run(bb, p)
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

func main() {
	bb := blackboard.NewBlackboard()
	p := proton.NewProton()

	Must(Setup(bb, p))

	err := Run(bb, p)
	if err != nil {
		log.Printf("Proton: %s", err)
	} else {
		log.Printf("Proton: Generated")
	}

	if WatchFileEnable(bb) && len(WatchFile(bb)) >= 1 {
		Must(Watcher(bb, p, WatchFile(bb)...))
		return
	}
}
