package ast

type ASTPropertyListStack [][]ASTProperty

func (s *ASTPropertyListStack) Push(v []ASTProperty) {
	*s = append(*s, v)
}

func (s *ASTPropertyListStack) Pop() []ASTProperty {
	var x []ASTProperty
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTPropertyListStack) Drain() [][]ASTProperty {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTPropertyListStack) Copy() [][]ASTProperty {
	return append([][]ASTProperty{}, *s...)
}

func (s *ASTPropertyListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
