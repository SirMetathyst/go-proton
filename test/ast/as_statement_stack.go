package ast

type ASTAsStatementStack []ASTAsStatement

func (s *ASTAsStatementStack) PushAsStatement(ID, StringLiteral string) {
	*s = append(*s, NewASTAsStatement(ID, StringLiteral))
}

func (s *ASTAsStatementStack) Push(v ASTAsStatement) {
	*s = append(*s, v)
}

func (s *ASTAsStatementStack) Pop() ASTAsStatement {
	var x ASTAsStatement
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTAsStatementStack) Drain() []ASTAsStatement {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTAsStatementStack) Copy() []ASTAsStatement {
	return append([]ASTAsStatement{}, *s...)
}

func (s *ASTAsStatementStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
