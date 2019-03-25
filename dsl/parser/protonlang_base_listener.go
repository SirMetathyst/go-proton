// Code generated from ProtonLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ProtonLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProtonLangListener is a complete listener for a parse tree produced by ProtonLangParser.
type BaseProtonLangListener struct{}

var _ ProtonLangListener = &BaseProtonLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProtonLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProtonLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProtonLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProtonLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseProtonLangListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseProtonLangListener) ExitDocument(ctx *DocumentContext) {}

// EnterIncludeStatement is called when production includeStatement is entered.
func (s *BaseProtonLangListener) EnterIncludeStatement(ctx *IncludeStatementContext) {}

// ExitIncludeStatement is called when production includeStatement is exited.
func (s *BaseProtonLangListener) ExitIncludeStatement(ctx *IncludeStatementContext) {}

// EnterTargetStatement is called when production targetStatement is entered.
func (s *BaseProtonLangListener) EnterTargetStatement(ctx *TargetStatementContext) {}

// ExitTargetStatement is called when production targetStatement is exited.
func (s *BaseProtonLangListener) ExitTargetStatement(ctx *TargetStatementContext) {}

// EnterNamespaceStatement is called when production namespaceStatement is entered.
func (s *BaseProtonLangListener) EnterNamespaceStatement(ctx *NamespaceStatementContext) {}

// ExitNamespaceStatement is called when production namespaceStatement is exited.
func (s *BaseProtonLangListener) ExitNamespaceStatement(ctx *NamespaceStatementContext) {}

// EnterContextStatement is called when production contextStatement is entered.
func (s *BaseProtonLangListener) EnterContextStatement(ctx *ContextStatementContext) {}

// ExitContextStatement is called when production contextStatement is exited.
func (s *BaseProtonLangListener) ExitContextStatement(ctx *ContextStatementContext) {}

// EnterContext is called when production context is entered.
func (s *BaseProtonLangListener) EnterContext(ctx *ContextContext) {}

// ExitContext is called when production context is exited.
func (s *BaseProtonLangListener) ExitContext(ctx *ContextContext) {}

// EnterAliasStatement is called when production aliasStatement is entered.
func (s *BaseProtonLangListener) EnterAliasStatement(ctx *AliasStatementContext) {}

// ExitAliasStatement is called when production aliasStatement is exited.
func (s *BaseProtonLangListener) ExitAliasStatement(ctx *AliasStatementContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseProtonLangListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseProtonLangListener) ExitAlias(ctx *AliasContext) {}

// EnterComponentStatement is called when production componentStatement is entered.
func (s *BaseProtonLangListener) EnterComponentStatement(ctx *ComponentStatementContext) {}

// ExitComponentStatement is called when production componentStatement is exited.
func (s *BaseProtonLangListener) ExitComponentStatement(ctx *ComponentStatementContext) {}

// EnterAsStatement is called when production asStatement is entered.
func (s *BaseProtonLangListener) EnterAsStatement(ctx *AsStatementContext) {}

// ExitAsStatement is called when production asStatement is exited.
func (s *BaseProtonLangListener) ExitAsStatement(ctx *AsStatementContext) {}

// EnterInStatement is called when production inStatement is entered.
func (s *BaseProtonLangListener) EnterInStatement(ctx *InStatementContext) {}

// ExitInStatement is called when production inStatement is exited.
func (s *BaseProtonLangListener) ExitInStatement(ctx *InStatementContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseProtonLangListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseProtonLangListener) ExitComponent(ctx *ComponentContext) {}

// EnterEntityIndexStatement is called when production entityIndexStatement is entered.
func (s *BaseProtonLangListener) EnterEntityIndexStatement(ctx *EntityIndexStatementContext) {}

// ExitEntityIndexStatement is called when production entityIndexStatement is exited.
func (s *BaseProtonLangListener) ExitEntityIndexStatement(ctx *EntityIndexStatementContext) {}

// EnterEntityIndex is called when production entityIndex is entered.
func (s *BaseProtonLangListener) EnterEntityIndex(ctx *EntityIndexContext) {}

// ExitEntityIndex is called when production entityIndex is exited.
func (s *BaseProtonLangListener) ExitEntityIndex(ctx *EntityIndexContext) {}

// EnterEntityIndexMethodStatementList is called when production entityIndexMethodStatementList is entered.
func (s *BaseProtonLangListener) EnterEntityIndexMethodStatementList(ctx *EntityIndexMethodStatementListContext) {
}

// ExitEntityIndexMethodStatementList is called when production entityIndexMethodStatementList is exited.
func (s *BaseProtonLangListener) ExitEntityIndexMethodStatementList(ctx *EntityIndexMethodStatementListContext) {
}

// EnterEntityIndexMethodStatement is called when production entityIndexMethodStatement is entered.
func (s *BaseProtonLangListener) EnterEntityIndexMethodStatement(ctx *EntityIndexMethodStatementContext) {
}

// ExitEntityIndexMethodStatement is called when production entityIndexMethodStatement is exited.
func (s *BaseProtonLangListener) ExitEntityIndexMethodStatement(ctx *EntityIndexMethodStatementContext) {
}

// EnterMemberStatementList is called when production memberStatementList is entered.
func (s *BaseProtonLangListener) EnterMemberStatementList(ctx *MemberStatementListContext) {}

// ExitMemberStatementList is called when production memberStatementList is exited.
func (s *BaseProtonLangListener) ExitMemberStatementList(ctx *MemberStatementListContext) {}

// EnterMemberStatement is called when production memberStatement is entered.
func (s *BaseProtonLangListener) EnterMemberStatement(ctx *MemberStatementContext) {}

// ExitMemberStatement is called when production memberStatement is exited.
func (s *BaseProtonLangListener) ExitMemberStatement(ctx *MemberStatementContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseProtonLangListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseProtonLangListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseProtonLangListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseProtonLangListener) ExitProperty(ctx *PropertyContext) {}

// EnterKeyValueList is called when production keyValueList is entered.
func (s *BaseProtonLangListener) EnterKeyValueList(ctx *KeyValueListContext) {}

// ExitKeyValueList is called when production keyValueList is exited.
func (s *BaseProtonLangListener) ExitKeyValueList(ctx *KeyValueListContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseProtonLangListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseProtonLangListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseProtonLangListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseProtonLangListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseProtonLangListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseProtonLangListener) ExitIdentifier(ctx *IdentifierContext) {}
