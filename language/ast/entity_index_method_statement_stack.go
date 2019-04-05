package ast

type ASTEntityIndexMethodStatementStack []ASTMethodStatement

func (s *ASTEntityIndexMethodStatementStack) PushEntityIndexMethodStatement(ID string, MemberStatementList []ASTMemberStatement) {
	*s = append(*s, NewASTMethodStatementWith(ID, MemberStatementList))
}

func (s *ASTEntityIndexMethodStatementStack) Push(v ASTMethodStatement) {
	*s = append(*s, v)
}

func (s *ASTEntityIndexMethodStatementStack) Pop() ASTMethodStatement {
	var x ASTMethodStatement
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTEntityIndexMethodStatementStack) Drain() []ASTMethodStatement {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTEntityIndexMethodStatementStack) Copy() []ASTMethodStatement {
	return append([]ASTMethodStatement{}, *s...)
}

func (s *ASTEntityIndexMethodStatementStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
