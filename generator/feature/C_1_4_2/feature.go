package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// FeatureGenerator_C_1_4_2 ...
func FeatureGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Feature.cs", Feature_C_1_4_2(new(bytes.Buffer)), "Feature_C_1_4_2"))
	return slice, nil
}
