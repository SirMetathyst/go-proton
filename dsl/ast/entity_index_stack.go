package ast

type ASTEntityIndexStack []ASTEntityIndex

func (s *ASTEntityIndexStack) PushEntityIndex(ID string, Option []ASTKeyValue, Context string, Method []ASTMethodStatement) {
	*s = append(*s, NewASTEntityIndexWith(ID, Option, Context, Method))
}

func (s *ASTEntityIndexStack) Push(v ASTEntityIndex) {
	*s = append(*s, v)
}

func (s *ASTEntityIndexStack) Pop() ASTEntityIndex {
	var x ASTEntityIndex
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTEntityIndexStack) Drain() []ASTEntityIndex {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTEntityIndexStack) Copy() []ASTEntityIndex {
	return append([]ASTEntityIndex{}, *s...)
}

func (s *ASTEntityIndexStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
