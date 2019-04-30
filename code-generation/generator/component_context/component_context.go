package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_context/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpComponentContextGenerator_C_1_4_2", true)
}
