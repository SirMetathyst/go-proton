package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpEntityGenerator_C_1_4_2", true)
}
