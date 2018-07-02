package postprocessor

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/SirMetathyst/proton"
)

var (
	Directory = "Generated"
)

// WriteToDiskPostProcessor_C_1_4_2 ...
func WriteToDiskPostProcessor_C_1_4_2(f []proton.FileInfo) ([]proton.FileInfo, error) {
	for _, cf := range f {
		WriteFile(Directory+"/"+cf.File, []byte(cf.FileContent))
	}
	return f, nil
}

// CreateDirectory ...
func CreateDirectory(Path string) error {
	return os.MkdirAll(Path, 0644)
}

// WriteFile ...
func WriteFile(Path string, Data []byte) error {
	err := CreateDirectory(filepath.Dir(Path))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Path, Data, 0666)
	return err
}
