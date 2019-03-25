package dsl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gobuffalo/packr"

	proton "github.com/SirMetathyst/go-proton"
	"github.com/SirMetathyst/go-proton/builder"
	"github.com/SirMetathyst/go-proton/dsl/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ASTList ...
func ASTList(RootAST *AST) []*AST {
	List := make([]*AST, 0)
	if RootAST != nil {
		for _, Include := range RootAST.IncludeList.List {
			List = append(List, Include.AST)
		}
		List = append(List, RootAST)
	}
	return List
}

func getEventTarget(b string) proton.EventTarget {
	if b == "self" {
		return proton.SelfTarget
	} else if b == "any" {
		return proton.AnyTarget
	}
	return proton.NoTarget
}

func getEventType(e string) proton.EventType {
	if e == "removed" {
		return proton.RemovedEvent
	}
	return proton.AddedEvent
}

func getCleanupMode(m string) proton.CleanupMode {
	if m == "destroy_entity" {
		return proton.DestroyEntity
	} else if m == "remove_component" {
		return proton.RemoveComponent
	}
	return proton.NoCleanup
}

func getEntityIndex(e string) proton.EntityIndex {
	if e == "multiple" {
		return proton.MultipleEntityIndex
	} else if e == "single" {
		return proton.SingleEntityIndex
	}
	return proton.NoEntityIndex
}

// ASTToModel ...
func ASTToModel(RootAST *AST) (*proton.MD, error) {

	ASTList := ASTList(RootAST)

	mdb := builder.NewModelBuilder()

	mdb.SetTarget(RootAST.Target)
	mdb.SetNamespace(RootAST.Namespace)

	for _, AST := range ASTList {
		for _, contextList := range AST.ContextListStack {
			for _, context := range contextList {
				err := mdb.NewContext().SetID(context.ID).Build()
				if err != nil {
					return nil, err
				}
				for _, option := range context.Option {
					switch option.Key {
					case "default":
						mdb.SetDefaultContext(context.ID)
					}
				}
			}
		}

		for _, aliasList := range AST.AliasListStack {
			for _, alias := range aliasList {
				err := mdb.NewAlias().SetID(alias.ID).SetValue(alias.Value).Build()
				if err != nil {
					return nil, err
				}
			}
		}

		for _, componentStatement := range AST.ComponentStatementStack {
			for _, component := range componentStatement.ComponentList {
				newComponent := mdb.NewComponent()
				newComponent.SetID(component.ID)

				for _, context := range componentStatement.ContextList {
					newComponent.AddContext(context.Key)
				}

				for _, attribute := range component.AttributeList {
					switch attribute.Key {
					case "unique":
						newComponent.SetUnique(true)
					case "flagPrefix":
						newComponent.SetFlagPrefix(attribute.Value)
					case "eventTarget":
						newComponent.SetEventTarget(getEventTarget(attribute.Value))
					case "eventPriority":
						{
							i, err := strconv.Atoi(attribute.Value)
							if err == nil {
								newComponent.SetEventPriority(i)
							}
						}
					case "eventType":
						newComponent.SetEventType(getEventType(attribute.Value))
					case "cleanupMode":
						newComponent.SetCleanupMode(getCleanupMode(attribute.Value))
					}
				}

				for _, memberStatement := range componentStatement.MemberStatementList {
					for _, property := range memberStatement.PropertyList {
						if memberStatement.ID != "" {
							m := newComponent.NewMember().
								SetID(property.ID).
								SetAlias(memberStatement.ID)

							for _, attribute := range property.AttributeList {
								switch attribute.Key {
								case "entityIndex":
									m.SetEntityIndex(getEntityIndex(attribute.Value))
								}
							}

							err := m.Build()
							if err != nil {
								return nil, err
							}
						} else {
							m := newComponent.NewMember().
								SetID(property.ID).
								SetValue(memberStatement.StringLiteral)

							for _, attribute := range property.AttributeList {
								switch attribute.Key {
								case "entityIndex":
									m.SetEntityIndex(getEntityIndex(attribute.Value))
								}
							}

							err := m.Build()
							if err != nil {
								return nil, err
							}
						}
					}
				}

				if componentStatement.AsStatement.ID != "" {
					m := newComponent.NewMember().
						SetID("value").
						SetAlias(componentStatement.AsStatement.ID)
					err := m.Build()
					if err != nil {
						return nil, err
					}
				} else if componentStatement.AsStatement.StringLiteral != "" {
					m := newComponent.NewMember().
						SetID("value").
						SetValue(componentStatement.AsStatement.StringLiteral)
					err := m.Build()
					if err != nil {
						return nil, err
					}
				}
				err := newComponent.Build()
				if err != nil {
					return nil, err
				}

			}
		}

		for _, entityIndex := range AST.EntityIndexStack {
			eib := mdb.NewEntityIndex()
			eib.SetID(entityIndex.ID)
			eib.AddContext(entityIndex.Context)

			for _, option := range entityIndex.Option {
				switch option.Key {
				case "primary":
					eib.SetPrimary(true)
				}
			}

			for _, method := range entityIndex.Method {
				eimb := eib.NewMethod()
				eimb.SetID(method.ID)

				for _, memberStatement := range method.MemberStatementList {
					for _, property := range memberStatement.PropertyList {
						if memberStatement.ID != "" {
							m := eimb.NewMember().
								SetID(property.ID).
								SetAlias(memberStatement.ID)
							err := m.Build()
							if err != nil {
								return nil, err
							}
						} else {
							m := eimb.NewMember().
								SetID(property.ID).
								SetValue(memberStatement.StringLiteral)
							err := m.Build()
							if err != nil {
								return nil, err
							}
						}
					}
				}

				err := eimb.Build()
				if err != nil {
					return nil, err
				}
			}

			err := eib.Build()
			if err != nil {
				return nil, err
			}
		}
	}

	return mdb.Build()
}

// parseInclude ...
func parseInclude(include string, includeList *ASTIncludeList) (*AST, error) {
	box := packr.NewBox("./alias")
	var lexer *parser.ProtonLangLexer

	if box.Has(include) {
		inputStream := antlr.NewInputStream(box.String(include))
		lexer = parser.NewProtonLangLexer(inputStream)
	} else {
		filestream, err := antlr.NewFileStream(include)
		if err != nil {
			return nil, err
		}
		lexer = parser.NewProtonLangLexer(filestream)
	}

	pkgDir := filepath.Dir(include)
	pkgDirAbs, err := filepath.Abs(pkgDir)
	if err != nil {
		return nil, err
	}

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewProtonLangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	listener := NewDefaultProtonLangListenerWith(include, pkgDirAbs, includeList)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AST, nil
}

// Parse ...
func Parse(pkg string) (*proton.MD, error) {
	file, err := os.Open(pkg)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	pkgDir := filepath.Dir(pkg)
	pkgDirAbs, err := filepath.Abs(pkgDir)
	if err != nil {
		return nil, err
	}
	inputStream := antlr.NewInputStream(string(bytes))
	lexer := parser.NewProtonLangLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewProtonLangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	listener := NewDefaultProtonLangListener(pkgDirAbs, pkgDirAbs)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return ASTToModel(listener.AST)
}
