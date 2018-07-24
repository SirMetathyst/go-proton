package postprocessor

// fII ...
type fII interface {
	File() string
	FileContent() string
	Generator() string
}

// fI ...
type fI struct {
	f, fc, g string
}

// File ...
func (fi fI) File() string {
	return fi.f
}

// FileContent ...
func (fi fI) FileContent() string {
	return fi.fc
}

// Generator ...
func (fi fI) Generator() string {
	return fi.g
}

// newFileInfo ...
func newFileInfo(file, fileContent, generator string) fI {
	return fI{file, fileContent, generator}
}
