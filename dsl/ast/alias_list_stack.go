package ast

type ASTAliasListStack [][]ASTAlias

func (s *ASTAliasListStack) Push(v []ASTAlias) {
	*s = append(*s, v)
}

func (s *ASTAliasListStack) Pop() []ASTAlias {
	var x []ASTAlias
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTAliasListStack) Drain() [][]ASTAlias {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTAliasListStack) Copy() [][]ASTAlias {
	return append([][]ASTAlias{}, *s...)
}

func (s *ASTAliasListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
