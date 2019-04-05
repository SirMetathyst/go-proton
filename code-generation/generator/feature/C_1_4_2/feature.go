package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddGenerator("CSharpFeatureGenerator_C_1_4_2", FeatureGenerator_C_1_4_2, false)
}

// FeatureGenerator_C_1_4_2 ...
func FeatureGenerator_C_1_4_2(md *proton.MD) ([]proton.FI, error) {
	slice := make([]proton.FI, 0)
	slice = append(slice, proton.NewFileInfo("Feature.cs", Feature_C_1_4_2(new(bytes.Buffer)), "Feature_C_1_4_2"))
	return slice, nil
}
