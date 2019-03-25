package ast

type ASTContextStack []ASTContext

func (s *ASTContextStack) PushContext(ID, Divert string, KeyValueList []ASTKeyValue) {
	*s = append(*s, NewASTContextWith(ID, Divert, KeyValueList))
}

func (s *ASTContextStack) Push(v ASTContext) {
	*s = append(*s, v)
}

func (s *ASTContextStack) Pop() ASTContext {
	var x ASTContext
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTContextStack) Drain() []ASTContext {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTContextStack) Copy() []ASTContext {
	return append([]ASTContext{}, *s...)
}

func (s *ASTContextStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
