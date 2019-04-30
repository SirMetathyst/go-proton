package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/composite_system/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpCompositeSystemGenerator_C_1_6_1", true)
}
