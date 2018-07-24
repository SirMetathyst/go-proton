package postprocessor

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	WriteToDiskDirectory = "Generated"
)

// file ...
type file interface {
	File() string
	FileContent() string
}

// WriteToDiskPostProcessor_C_1_4_2 ...
func WriteToDiskPostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	for _, cv := range v {
		f := cv.(file)
		WriteFile(WriteToDiskDirectory+"/"+f.File(), []byte(f.FileContent()))
	}
	return v, nil
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
