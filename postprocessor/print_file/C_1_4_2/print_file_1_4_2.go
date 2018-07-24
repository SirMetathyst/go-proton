package postprocessor

import (
	"fmt"
)

// file ...
type file interface {
	File() string
}

// PrintFilePostProcessor_C_1_4_2 ...
func PrintFilePostProcessor_C_1_4_2(v []interface{}) ([]interface{}, error) {
	for _, cv := range v {
		f := cv.(file)
		fmt.Println(f.File())
	}
	return v, nil
}
