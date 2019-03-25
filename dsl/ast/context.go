package ast

type ASTContext struct {
	ID, Divert string
	Option     []ASTKeyValue
}

func NewASTContextWith(ID, Divert string, KeyValueList []ASTKeyValue) ASTContext {
	return ASTContext{ID, Divert, KeyValueList}
}
