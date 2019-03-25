// Code generated from ProtonLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ProtonLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProtonLangListener is a complete listener for a parse tree produced by ProtonLangParser.
type ProtonLangListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterIncludeStatement is called when entering the includeStatement production.
	EnterIncludeStatement(c *IncludeStatementContext)

	// EnterTargetStatement is called when entering the targetStatement production.
	EnterTargetStatement(c *TargetStatementContext)

	// EnterNamespaceStatement is called when entering the namespaceStatement production.
	EnterNamespaceStatement(c *NamespaceStatementContext)

	// EnterContextStatement is called when entering the contextStatement production.
	EnterContextStatement(c *ContextStatementContext)

	// EnterContext is called when entering the context production.
	EnterContext(c *ContextContext)

	// EnterAliasStatement is called when entering the aliasStatement production.
	EnterAliasStatement(c *AliasStatementContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterComponentStatement is called when entering the componentStatement production.
	EnterComponentStatement(c *ComponentStatementContext)

	// EnterAsStatement is called when entering the asStatement production.
	EnterAsStatement(c *AsStatementContext)

	// EnterInStatement is called when entering the inStatement production.
	EnterInStatement(c *InStatementContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterEntityIndexStatement is called when entering the entityIndexStatement production.
	EnterEntityIndexStatement(c *EntityIndexStatementContext)

	// EnterEntityIndex is called when entering the entityIndex production.
	EnterEntityIndex(c *EntityIndexContext)

	// EnterEntityIndexMethodStatementList is called when entering the entityIndexMethodStatementList production.
	EnterEntityIndexMethodStatementList(c *EntityIndexMethodStatementListContext)

	// EnterEntityIndexMethodStatement is called when entering the entityIndexMethodStatement production.
	EnterEntityIndexMethodStatement(c *EntityIndexMethodStatementContext)

	// EnterMemberStatementList is called when entering the memberStatementList production.
	EnterMemberStatementList(c *MemberStatementListContext)

	// EnterMemberStatement is called when entering the memberStatement production.
	EnterMemberStatement(c *MemberStatementContext)

	// EnterPropertyList is called when entering the propertyList production.
	EnterPropertyList(c *PropertyListContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterKeyValueList is called when entering the keyValueList production.
	EnterKeyValueList(c *KeyValueListContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitIncludeStatement is called when exiting the includeStatement production.
	ExitIncludeStatement(c *IncludeStatementContext)

	// ExitTargetStatement is called when exiting the targetStatement production.
	ExitTargetStatement(c *TargetStatementContext)

	// ExitNamespaceStatement is called when exiting the namespaceStatement production.
	ExitNamespaceStatement(c *NamespaceStatementContext)

	// ExitContextStatement is called when exiting the contextStatement production.
	ExitContextStatement(c *ContextStatementContext)

	// ExitContext is called when exiting the context production.
	ExitContext(c *ContextContext)

	// ExitAliasStatement is called when exiting the aliasStatement production.
	ExitAliasStatement(c *AliasStatementContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitComponentStatement is called when exiting the componentStatement production.
	ExitComponentStatement(c *ComponentStatementContext)

	// ExitAsStatement is called when exiting the asStatement production.
	ExitAsStatement(c *AsStatementContext)

	// ExitInStatement is called when exiting the inStatement production.
	ExitInStatement(c *InStatementContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitEntityIndexStatement is called when exiting the entityIndexStatement production.
	ExitEntityIndexStatement(c *EntityIndexStatementContext)

	// ExitEntityIndex is called when exiting the entityIndex production.
	ExitEntityIndex(c *EntityIndexContext)

	// ExitEntityIndexMethodStatementList is called when exiting the entityIndexMethodStatementList production.
	ExitEntityIndexMethodStatementList(c *EntityIndexMethodStatementListContext)

	// ExitEntityIndexMethodStatement is called when exiting the entityIndexMethodStatement production.
	ExitEntityIndexMethodStatement(c *EntityIndexMethodStatementContext)

	// ExitMemberStatementList is called when exiting the memberStatementList production.
	ExitMemberStatementList(c *MemberStatementListContext)

	// ExitMemberStatement is called when exiting the memberStatement production.
	ExitMemberStatement(c *MemberStatementContext)

	// ExitPropertyList is called when exiting the propertyList production.
	ExitPropertyList(c *PropertyListContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitKeyValueList is called when exiting the keyValueList production.
	ExitKeyValueList(c *KeyValueListContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
