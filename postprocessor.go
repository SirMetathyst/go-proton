package proton

// PP ...
type PP func([]interface{}) ([]interface{}, error)

// PPI ...
type PPI struct {
	PostProcessorVersion string
	PostProcessor        PP
	Enabled              bool
}

// NewPostProcessorInfo ...
func NewPostProcessorInfo(postProcessorVersion string, postProcessor PP, enabled bool) *PPI {
	return &PPI{postProcessorVersion, postProcessor, enabled}
}
