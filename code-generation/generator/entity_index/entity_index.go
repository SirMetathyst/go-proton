package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/entity_index/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpEntityIndexGenerator_C_1_4_2", true)
}
