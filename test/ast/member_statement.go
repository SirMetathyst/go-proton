package ast

// ASTMemberStatement ...
type ASTMemberStatement struct {
	PropertyList  []ASTProperty
	ID            string
	StringLiteral string
}

// NewASTMemberStatement ...
func NewASTMemberStatement(PropertyList []ASTProperty, ID, StringLiteral string) ASTMemberStatement {
	return ASTMemberStatement{PropertyList, ID, StringLiteral}
}
