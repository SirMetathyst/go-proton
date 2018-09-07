package main

import (

	//protonlang "github.com/SirMetathyst/proton-lang"

	"io/ioutil"
	"os"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-proton-lang"
)

// ProtonLang ...
func ProtonLang(bb *blackboard.BB) (*entitas.MD, error) {
	file, err := os.Open(suffix(File(bb), ".proton"))
	defer file.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	md, err := protonlang.Parse(string(bytes))
	if err != nil {
		return nil, err
	}
	return md, nil
}
