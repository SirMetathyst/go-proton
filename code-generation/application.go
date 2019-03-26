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
	// ErrNoProtonFilesFound is returned when no *.proton files were found
	ErrNoProtonFilesFound = fmt.Errorf("proton: no *.proton files were found")
)

// Must ...
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
func Daemon(path string, parser func(file string) (*proton.MD, error)) error {
	w, err := fsnotify.NewWatcher()
	defer w.Close()
	Must(err)

	files, err := ProtonFiles(path)
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
func RunApplication(root string, daemon bool, parser func(file string) (*proton.MD, error)) {
	Must(RunProton(parser))

	SetProjectPath(root)

	if daemon == true {
		Must(Daemon(root, parser))
		return
	}
}
