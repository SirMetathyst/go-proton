package proton

// FII ...
type FII interface {
	File() string
	FileContent() string
	Generator() string
}

// FI ...
type FI struct {
	f, fc, g string
}

// File ...
func (fi FI) File() string {
	return fi.f
}

// FileContent ...
func (fi FI) FileContent() string {
	return fi.fc
}

// Generator ...
func (fi FI) Generator() string {
	return fi.g
}

// NewFileInfo ...
func NewFileInfo(file, fileContent, generator string) FI {
	return FI{file, fileContent, generator}
}
