package ast

type ASTKeyValue struct {
	Key   string
	Value string
}

func NewASTKeyValueWith(k, v string) ASTKeyValue {
	return ASTKeyValue{k, v}
}
