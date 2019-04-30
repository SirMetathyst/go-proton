package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/contexts/C_1_4_2"
)

func init() {
	codegeneration.EnableGenerator("CSharpContextsGenerator_C_1_4_2", true)
}
