package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context/C_1_4_2"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/context/C_1_9_0"
)

func init() {
	codegeneration.EnableGenerator("CSharpContextGenerator_C_1_9_0", true)
}
