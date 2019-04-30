package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/component/E_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpComponentGenerator_E_1_4_2", true)
}
