package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_observer/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpContextObserverGenerator_C_1_4_2", true)
}
