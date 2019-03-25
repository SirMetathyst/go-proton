package ast

type ASTMemberStatementListStack [][]ASTMemberStatement

func (s *ASTMemberStatementListStack) Push(v []ASTMemberStatement) {
	*s = append(*s, v)
}

func (s *ASTMemberStatementListStack) Pop() []ASTMemberStatement {
	var x []ASTMemberStatement
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTMemberStatementListStack) Drain() [][]ASTMemberStatement {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTMemberStatementListStack) Copy() [][]ASTMemberStatement {
	return append([][]ASTMemberStatement{}, *s...)
}

func (s *ASTMemberStatementListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
