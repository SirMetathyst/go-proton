package main

import (
	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-proton-lang"
)

// ProtonLang ...
func ProtonLang(bb *blackboard.BB) (*entitas.MD, error) {
	md, err := protonlang.Parse(suffix(File(bb), ".proton"))
	if err != nil {
		return nil, err
	}
	return md, nil
}
