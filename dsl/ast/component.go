package ast

// ASTComponent ...
type ASTComponent struct {
	ID            string
	AttributeList []ASTKeyValue
}

// NewASTComponent ...
func NewASTComponent(ID string, Attribute []ASTKeyValue) ASTComponent {
	return ASTComponent{ID, Attribute}
}
