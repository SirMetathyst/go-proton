package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context_matcher/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpContextMatcherGenerator_C_1_4_2", true)
}
