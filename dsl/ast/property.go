package ast

// ASTProperty ...
type ASTProperty struct {
	ID            string
	AttributeList []ASTKeyValue
}

// NewASTProperty ...
func NewASTProperty(ID string, Attribute []ASTKeyValue) ASTProperty {
	return ASTProperty{ID, Attribute}
}
