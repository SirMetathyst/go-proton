package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"
	"github.com/SirMetathyst/proton/model"
)

// FeatureGenerator_C_1_4_2 ...
func FeatureGenerator_C_1_4_2(m *model.M) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	slice = append(slice, proton.NewFileInfo("Feature.cs", Feature_C_1_4_2(new(bytes.Buffer)), "Feature_C_1_4_2"))
	return slice, nil
}
