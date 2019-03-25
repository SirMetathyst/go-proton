package ast

type ASTComponentStack []ASTComponent

func (s *ASTComponentStack) PushComponent(ID string, Attribute []ASTKeyValue) {
	*s = append(*s, NewASTComponent(ID, Attribute))
}

func (s *ASTComponentStack) Push(v ASTComponent) {
	*s = append(*s, v)
}

func (s *ASTComponentStack) Pop() ASTComponent {
	var x ASTComponent
	if len(*s)-1 >= 0 {
		x = (*s)[len(*s)-1]
	}

	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *ASTComponentStack) Drain() []ASTComponent {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTComponentStack) Copy() []ASTComponent {
	return append([]ASTComponent{}, *s...)
}

func (s *ASTComponentStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
