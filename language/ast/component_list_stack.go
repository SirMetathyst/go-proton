package ast

type ASTComponentListStack [][]ASTComponent

func (s *ASTComponentListStack) Push(v []ASTComponent) {
	*s = append(*s, v)
}

func (s *ASTComponentListStack) Pop() []ASTComponent {
	var x []ASTComponent
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTComponentListStack) Drain() [][]ASTComponent {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTComponentListStack) Copy() [][]ASTComponent {
	return append([][]ASTComponent{}, *s...)
}

func (s *ASTComponentListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
