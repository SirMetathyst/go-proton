package ast

// ASTEntityIndex ...
type ASTEntityIndex struct {
	ID      string
	Option  []ASTKeyValue
	Context string
	Method  []ASTMethodStatement
}

// NewASTEntityIndexWith ...
func NewASTEntityIndexWith(ID string, Option []ASTKeyValue, Context string, Method []ASTMethodStatement) ASTEntityIndex {
	return ASTEntityIndex{ID, Option, Context, Method}
}
