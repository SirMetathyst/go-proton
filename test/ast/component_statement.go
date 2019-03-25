package ast

// ASTComponentStatement ...
type ASTComponentStatement struct {
	ComponentList       []ASTComponent
	ContextList         []ASTKeyValue
	MemberStatementList []ASTMemberStatement
	AsStatement         ASTAsStatement
}

// NewASTComponentStatement ...
func NewASTComponentStatement(ComponentList []ASTComponent, ContextList []ASTKeyValue, MemberStatementList []ASTMemberStatement, AsStatement ASTAsStatement) ASTComponentStatement {
	return ASTComponentStatement{ComponentList, ContextList, MemberStatementList, AsStatement}
}
