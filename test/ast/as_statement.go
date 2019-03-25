package ast

// ASTAsStatement ...
type ASTAsStatement struct {
	ID, StringLiteral string
}

// NewASTAsStatement ...
func NewASTAsStatement(ID, StringLiteral string) ASTAsStatement {
	return ASTAsStatement{ID, StringLiteral}
}
