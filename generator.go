package proton

// G ...
type G func(interface{}) ([]interface{}, error)

// GI ...
type GI struct {
	GeneratorVersion string
	Generator        G
	Enabled          bool
}

// NewGeneratorInfo ...
func NewGeneratorInfo(generatorVersion string, generator G, enabled bool) *GI {
	return &GI{generatorVersion, generator, enabled}
}
