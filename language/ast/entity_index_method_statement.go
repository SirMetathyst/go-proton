package ast

// ASTMethodStatement ...
type ASTMethodStatement struct {
	ID                  string
	MemberStatementList []ASTMemberStatement
}

// NewASTMethodStatementWith ...
func NewASTMethodStatementWith(ID string, MemberStatementList []ASTMemberStatement) ASTMethodStatement {
	return ASTMethodStatement{ID, MemberStatementList}
}
