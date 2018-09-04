package main

import (

	//protonlang "github.com/SirMetathyst/proton-lang"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-entitas/builder"
)

// ProtonLang ...
func ProtonLang(bb *blackboard.BB) (*entitas.MD, error) {
	mdb := builder.NewModelBuilder()

	//file, err := os.Open(suffix(*c.StringP("File"), ".proton"))
	//defer file.Close()

	//bytes, err := ioutil.ReadAll(file)
	//e.Must(err)

	//m, err = protonlang.Parse(string(bytes))
	//e.Must(err)
	return mdb.Build()
}
