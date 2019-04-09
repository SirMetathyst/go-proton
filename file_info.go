package proton

// FileInfo ...
type FileInfo struct {
	file        string
	fileContent string
	generator   string
}

// File ...
func (fi FileInfo) File() string {
	return fi.file
}

// FileContent ...
func (fi FileInfo) FileContent() string {
	return fi.fileContent
}

// Generator ...
func (fi FileInfo) Generator() string {
	return fi.generator
}

// NewFileInfo ...
func NewFileInfo(file, fileContent, generator string) FileInfo {
	return FileInfo{file, fileContent, generator}
}
