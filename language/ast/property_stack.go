package ast

type ASTPropertyStack []ASTProperty

func (s *ASTPropertyStack) PushProperty(ID string, Attribute []ASTKeyValue) {
	*s = append(*s, NewASTProperty(ID, Attribute))
}

func (s *ASTPropertyStack) Push(v ASTProperty) {
	*s = append(*s, v)
}

func (s *ASTPropertyStack) Pop() ASTProperty {
	var x ASTProperty
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTPropertyStack) Drain() []ASTProperty {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTPropertyStack) Copy() []ASTProperty {
	return append([]ASTProperty{}, *s...)
}

func (s *ASTPropertyStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
