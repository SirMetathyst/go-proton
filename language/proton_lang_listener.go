package language

import (
	"errors"
	"log"
	"path/filepath"
	"strings"

	"github.com/SirMetathyst/go-proton/language/ast"
	"github.com/SirMetathyst/go-proton/language/parser"
)

// Error ...
var (
	ErrBlank = errors.New("")
)

// Constant ...
const (
	StringBlank = ""
)

// ASTInclude ...
type ASTInclude struct {
	File string
	AST  *AST
}

// ASTIncludeList ...
type ASTIncludeList struct {
	List       []*ASTInclude
	ListLookup map[string]*AST
}

// NewASTIncludeList ...
func NewASTIncludeList() *ASTIncludeList {
	return &ASTIncludeList{make([]*ASTInclude, 0), make(map[string]*AST, 0)}
}

// Create ASTInclude & Add ...
func (Include *ASTIncludeList) Create(File string, AST *AST) {
	Include.List = append(Include.List, &ASTInclude{File, AST})
	Include.ListLookup[File] = AST
}

// Add ASTInclude ...
func (Include *ASTIncludeList) Add(Value *ASTInclude) {
	Include.List = append(Include.List, Value)
	Include.ListLookup[Value.File] = Value.AST
}

// Exist ...
func (Include *ASTIncludeList) Exist(File string) bool {
	_, Ok := Include.ListLookup[File]
	return Ok
}

// AST ...
type AST struct {
	IncludeDir  string
	Include     string
	IncludeList *ASTIncludeList

	Target    string
	Namespace string

	/* CONTEXT. */
	ContextListStack ast.ASTContextListStack
	ContextStack     ast.ASTContextStack

	/* ALIAS. */
	AliasListStack ast.ASTAliasListStack
	AliasStack     ast.ASTAliasStack

	/* COMPONENT. */
	ComponentStatementStack ast.ASTComponentStatementStack
	ComponentListStack      ast.ASTComponentListStack
	ComponentStack          ast.ASTComponentStack
	AsStatementStack        ast.ASTAsStatementStack
	PropertyListStack       ast.ASTPropertyListStack
	PropertyStack           ast.ASTPropertyStack

	/* INDEX. */
	EntityIndexMethodStatementStack ast.ASTEntityIndexMethodStatementStack
	EntityIndexStack                ast.ASTEntityIndexStack

	/* MEMBER. */
	MemberStatementListStack ast.ASTMemberStatementListStack
	MemberStatementStack     ast.ASTMemberStatementStack

	/* KEY/VALUE. */
	KeyValueListStack ast.ASTKeyValueListStack
	KeyValueStack     ast.ASTKeyValueStack
}

// NewASTWith ...
func NewASTWith(include, includeDir string, IncludeList *ASTIncludeList) *AST {
	return &AST{Include: include, IncludeDir: includeDir, IncludeList: IncludeList}
}

// NewAST ...
func NewAST(include, includeDir string) *AST {
	return &AST{Include: include, IncludeDir: includeDir, IncludeList: NewASTIncludeList()}
}

// DefaultProtonLangListener ...
type DefaultProtonLangListener struct {
	*parser.BaseProtonLangListener
	AST *AST
}

// NewDefaultProtonLangListenerWith ...
func NewDefaultProtonLangListenerWith(Include, IncludeDir string, IncludeList *ASTIncludeList) *DefaultProtonLangListener {
	return &DefaultProtonLangListener{AST: NewASTWith(Include, IncludeDir, IncludeList)}
}

// NewDefaultProtonLangListener ...
func NewDefaultProtonLangListener(Include, IncludeDir string) *DefaultProtonLangListener {
	return &DefaultProtonLangListener{AST: NewAST(Include, IncludeDir)}
}

// ExitIncludeStatement is called when production includeStatement is exited.
func (Listener *DefaultProtonLangListener) ExitIncludeStatement(ctx *parser.IncludeStatementContext) {
	pkg := Optional(ctx.StringLiteral(), StringBlank)
	pkg = strings.Replace(pkg, "${PWD}", Listener.AST.IncludeDir, -1)
	if !strings.HasSuffix(pkg, ".proton") {
		pkg = pkg + ".proton"
	}
	pkgDir := filepath.Dir(pkg)
	pkgDirAbs, _ := filepath.Abs(pkgDir)
	if isIncluded := Listener.AST.IncludeList.Exist(pkg); !isIncluded {
		ast, err := parseInclude(pkg, Listener.AST.IncludeList)
		if err != nil {
			log.Println(err)
			return
		}
		Listener.AST.IncludeList.Create(pkgDirAbs, ast)
	}
}

// ExitTargetStatement is called when production targetStatement is exited.
func (Listener *DefaultProtonLangListener) ExitTargetStatement(ctx *parser.TargetStatementContext) {
	Listener.AST.Target = Optional(ctx.Identifier(), StringBlank)
}

// ExitNamespaceStatement is called when production namespaceStatement is exited.
func (Listener *DefaultProtonLangListener) ExitNamespaceStatement(ctx *parser.NamespaceStatementContext) {
	Listener.AST.Namespace = Optional(ctx.Namespace(), StringBlank)
}

// ExitContextStatement is called when production contextStatement is exited.
func (Listener *DefaultProtonLangListener) ExitContextStatement(ctx *parser.ContextStatementContext) {
	Listener.AST.ContextListStack.Push(Listener.AST.ContextStack.Drain())
}

// ExitContext is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitContext(ctx *parser.ContextContext) {
	Listener.AST.ContextStack.PushContext(Optional(ctx.Identifier(0), StringBlank), Optional(ctx.Identifier(1), StringBlank), Listener.AST.KeyValueListStack.Pop())
}

// ExitAliasStatement is called when production aliasStatement is exited.
func (Listener *DefaultProtonLangListener) ExitAliasStatement(ctx *parser.AliasStatementContext) {
	Listener.AST.AliasListStack.Push(Listener.AST.AliasStack.Drain())
}

// ExitAlias is called when production alias is exited.
func (Listener *DefaultProtonLangListener) ExitAlias(ctx *parser.AliasContext) {
	Listener.AST.AliasStack.PushAlias(Optional(ctx.Identifier(), StringBlank), Optional(ctx.StringLiteral(), StringBlank))
}

// ExitComponentStatement is called when production componentStatement is exited.
func (Listener *DefaultProtonLangListener) ExitComponentStatement(ctx *parser.ComponentStatementContext) {
	Listener.AST.ComponentListStack.Push(Listener.AST.ComponentStack.Drain())
	Listener.AST.ComponentStatementStack.PushComponentStatement(
		Listener.AST.ComponentListStack.Pop(),
		Listener.AST.KeyValueListStack.Pop(),
		Listener.AST.MemberStatementListStack.Pop(),
		Listener.AST.AsStatementStack.Pop())
}

// ExitComponentList is called when production context is exited.
//func (Listener *DefaultProtonLangListener) ExitComponentList(ctx *ComponentListContext) {
//	Listener.AST.ComponentListStack.Push(Listener.AST.ComponentStack.Drain())
//}

// ExitComponent is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitComponent(ctx *parser.ComponentContext) {
	Listener.AST.ComponentStack.PushComponent(Optional(ctx.Identifier(), StringBlank), Listener.AST.KeyValueListStack.Pop())
}

// ExitAsStatement is called when production asStatement is exited.
func (Listener *DefaultProtonLangListener) ExitAsStatement(ctx *parser.AsStatementContext) {
	Listener.AST.AsStatementStack.PushAsStatement(Optional(ctx.Identifier(), StringBlank), Optional(ctx.StringLiteral(), StringBlank))
}

// ExitPropertyList is called when production propertyList is exited.
func (Listener *DefaultProtonLangListener) ExitPropertyList(ctx *parser.PropertyListContext) {
	Listener.AST.PropertyListStack.Push(Listener.AST.PropertyStack.Drain())
}

// ExitProperty is called when production property is exited.
func (Listener *DefaultProtonLangListener) ExitProperty(ctx *parser.PropertyContext) {
	Listener.AST.PropertyStack.PushProperty(Optional(ctx.Identifier(), StringBlank), Listener.AST.KeyValueListStack.PopIf(ctx.KeyValueList() != nil))
}

// ExitEntityIndex is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitEntityIndex(ctx *parser.EntityIndexContext) {
	Listener.AST.EntityIndexStack.PushEntityIndex(Optional(ctx.Identifier(0), StringBlank), Listener.AST.KeyValueListStack.Pop(), Optional(ctx.Identifier(1), StringBlank), Listener.AST.EntityIndexMethodStatementStack.Drain())
}

// ExitEntityIndexMethodStatement is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitEntityIndexMethodStatement(ctx *parser.EntityIndexMethodStatementContext) {
	Listener.AST.EntityIndexMethodStatementStack.PushEntityIndexMethodStatement(Optional(ctx.Identifier(), StringBlank), Listener.AST.MemberStatementListStack.Pop())
}

// ExitMemberStatementList is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitMemberStatementList(ctx *parser.MemberStatementListContext) {
	Listener.AST.MemberStatementListStack.Push(Listener.AST.MemberStatementStack.Drain())
}

// ExitMemberStatement is called when production context is exited.
func (Listener *DefaultProtonLangListener) ExitMemberStatement(ctx *parser.MemberStatementContext) {

	// TODO: implement alias modifier

	Listener.AST.MemberStatementStack.PushMemberStatement(Listener.AST.PropertyListStack.Pop(), Optional(ctx.Identifier(), StringBlank), Optional(ctx.StringLiteral(), StringBlank))
}

// ExitKeyValueList is called when production keyValueList is exited.
func (Listener *DefaultProtonLangListener) ExitKeyValueList(ctx *parser.KeyValueListContext) {
	Listener.AST.KeyValueListStack.Push(Listener.AST.KeyValueStack.Drain())
}

// ExitKeyValue is called when production keyValue is exited.
func (Listener *DefaultProtonLangListener) ExitKeyValue(ctx *parser.KeyValueContext) {
	Listener.AST.KeyValueStack.PushKeyValue(Optional(ctx.Identifier(0), StringBlank), Optional(ctx.Identifier(1), StringBlank))
}
