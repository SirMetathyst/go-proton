package ast

type ASTContextListStack [][]ASTContext

func (s *ASTContextListStack) Push(v []ASTContext) {
	*s = append(*s, v)
}

func (s *ASTContextListStack) Pop() []ASTContext {
	var x []ASTContext
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTContextListStack) Drain() [][]ASTContext {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTContextListStack) Copy() [][]ASTContext {
	return append([][]ASTContext{}, *s...)
}

func (s *ASTContextListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
