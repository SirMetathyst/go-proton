package ast

type ASTKeyValueStack []ASTKeyValue

func (s *ASTKeyValueStack) PushKeyValue(k, v string) {
	*s = append(*s, NewASTKeyValueWith(k, v))
}

func (s *ASTKeyValueStack) Push(v ASTKeyValue) {
	*s = append(*s, v)
}

func (s *ASTKeyValueStack) Pop() ASTKeyValue {
	var x ASTKeyValue
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTKeyValueStack) Drain() []ASTKeyValue {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTKeyValueStack) Copy() []ASTKeyValue {
	return append([]ASTKeyValue{}, *s...)
}

func (s *ASTKeyValueStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
