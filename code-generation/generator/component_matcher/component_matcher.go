package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component_matcher/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpComponentMatcherGenerator_C_1_4_2", true)
}
