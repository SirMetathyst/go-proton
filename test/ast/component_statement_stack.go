package ast

type ASTComponentStatementStack []ASTComponentStatement

func (s *ASTComponentStatementStack) PushComponentStatement(ComponentList []ASTComponent, ContextList []ASTKeyValue, MemberStatementList []ASTMemberStatement, AsStatement ASTAsStatement) {
	*s = append(*s, NewASTComponentStatement(ComponentList, ContextList, MemberStatementList, AsStatement))
}

func (s *ASTComponentStatementStack) Push(v ASTComponentStatement) {
	*s = append(*s, v)
}

func (s *ASTComponentStatementStack) Pop() ASTComponentStatement {
	var x ASTComponentStatement
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTComponentStatementStack) Drain() []ASTComponentStatement {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTComponentStatementStack) Copy() []ASTComponentStatement {
	return append([]ASTComponentStatement{}, *s...)
}

func (s *ASTComponentStatementStack) Clear() {
	*s = (*s)[0:0]
}
