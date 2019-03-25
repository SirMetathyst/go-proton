package ast

type ASTAlias struct {
	ID    string
	Value string
}

func NewASTAliasWith(ID, Value string) ASTAlias {
	return ASTAlias{ID, Value}
}
