package ast

type ASTKeyValueListStack [][]ASTKeyValue

func (s *ASTKeyValueListStack) Push(v []ASTKeyValue) {
	*s = append(*s, v)
}

func (s *ASTKeyValueListStack) PopIf(v bool) []ASTKeyValue {
	if v {
		return s.Pop()
	}
	return nil
}

func (s *ASTKeyValueListStack) Pop() []ASTKeyValue {
	var x []ASTKeyValue
	if len(*s) != 0 {
		x = (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
	}
	return x
}

func (s *ASTKeyValueListStack) Drain() [][]ASTKeyValue {
	x := s.Copy()
	s.Clear()
	return x
}

func (s *ASTKeyValueListStack) Copy() [][]ASTKeyValue {
	return append([][]ASTKeyValue{}, *s...)
}

func (s *ASTKeyValueListStack) Clear() {
	*s = (*s)[0:0]
	if len(*s) != 0 {
		panic("stack not zero!")
	}
}
