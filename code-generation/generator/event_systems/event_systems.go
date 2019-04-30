package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_systems/C_1_6_1"
)

func init() {
	codegeneration.EnableGenerator("CSharpEventSystemsGenerator_C_1_6_1", true)
}
