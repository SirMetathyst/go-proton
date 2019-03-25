package ast

type ASTMemberStatementStack []ASTMemberStatement

func (s *ASTMemberStatementStack) PushMemberStatement(PropertyList []ASTProperty, ID, StringLiteral string) {
	*s = append(*s, NewASTMemberStatement(PropertyList, ID, StringLiteral))
}

func (s *ASTMemberStatementStack) Push(v ASTMemberStatement) {
	*s = append(*s, v)
}

func (s *ASTMemberStatementStack) Pop() ASTMemberStatement {
	var x ASTMemberStatement
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTMemberStatementStack) Drain() []ASTMemberStatement {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTMemberStatementStack) Copy() []ASTMemberStatement {
	return append([]ASTMemberStatement{}, *s...)
}

func (s *ASTMemberStatementStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
