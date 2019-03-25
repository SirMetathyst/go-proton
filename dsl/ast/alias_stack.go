package ast

type ASTAliasStack []ASTAlias

func (s *ASTAliasStack) PushAlias(ID, Value string) {
	*s = append(*s, NewASTAliasWith(ID, Value))
}

func (s *ASTAliasStack) Push(v ASTAlias) {
	*s = append(*s, v)
}

func (s *ASTAliasStack) Pop() ASTAlias {
	var x ASTAlias
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTAliasStack) Drain() []ASTAlias {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTAliasStack) Copy() []ASTAlias {
	return append([]ASTAlias{}, *s...)
}

func (s *ASTAliasStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
