// Code generated from ProtonLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ProtonLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 278,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11,
	2, 3, 2, 5, 2, 58, 10, 2, 3, 2, 5, 2, 61, 10, 2, 3, 2, 7, 2, 64, 10, 2,
	12, 2, 14, 2, 67, 11, 2, 3, 2, 7, 2, 70, 10, 2, 12, 2, 14, 2, 73, 11, 2,
	3, 2, 7, 2, 76, 10, 2, 12, 2, 14, 2, 79, 11, 2, 3, 2, 7, 2, 82, 10, 2,
	12, 2, 14, 2, 85, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 103, 10, 6, 12, 6,
	14, 6, 106, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 113, 10, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 119, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 125,
	10, 8, 12, 8, 14, 8, 128, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 7, 10, 140, 10, 10, 12, 10, 14, 10, 143, 11, 10,
	3, 10, 5, 10, 146, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 155, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	163, 10, 11, 3, 11, 3, 11, 5, 11, 167, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 177, 10, 13, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 187, 10, 15, 3, 15, 3, 15, 5,
	15, 191, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 6, 16, 197, 10, 16, 13, 16,
	14, 16, 198, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 7,
	18, 209, 10, 18, 12, 18, 14, 18, 212, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 223, 10, 19, 3, 19, 5, 19, 226,
	10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 233, 10, 20, 12, 20,
	14, 20, 236, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 243, 10,
	21, 3, 22, 3, 22, 3, 22, 7, 22, 248, 10, 22, 12, 22, 14, 22, 251, 11, 22,
	3, 23, 3, 23, 3, 23, 5, 23, 256, 10, 23, 3, 24, 6, 24, 259, 10, 24, 13,
	24, 14, 24, 260, 3, 24, 3, 24, 7, 24, 265, 10, 24, 12, 24, 14, 24, 268,
	11, 24, 3, 25, 5, 25, 271, 10, 25, 3, 25, 6, 25, 274, 10, 25, 13, 25, 14,
	25, 275, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 3, 4, 2, 16, 16, 27, 28,
	2, 288, 2, 53, 3, 2, 2, 2, 4, 86, 3, 2, 2, 2, 6, 90, 3, 2, 2, 2, 8, 94,
	3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12, 109, 3, 2, 2, 2, 14, 120, 3, 2, 2,
	2, 16, 131, 3, 2, 2, 2, 18, 135, 3, 2, 2, 2, 20, 156, 3, 2, 2, 2, 22, 168,
	3, 2, 2, 2, 24, 171, 3, 2, 2, 2, 26, 178, 3, 2, 2, 2, 28, 181, 3, 2, 2,
	2, 30, 194, 3, 2, 2, 2, 32, 202, 3, 2, 2, 2, 34, 206, 3, 2, 2, 2, 36, 215,
	3, 2, 2, 2, 38, 229, 3, 2, 2, 2, 40, 237, 3, 2, 2, 2, 42, 244, 3, 2, 2,
	2, 44, 252, 3, 2, 2, 2, 46, 258, 3, 2, 2, 2, 48, 270, 3, 2, 2, 2, 50, 52,
	5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	53, 54, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 58, 5,
	6, 4, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59,
	61, 5, 8, 5, 2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 65, 3, 2, 2,
	2, 62, 64, 5, 10, 6, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 71, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	68, 70, 5, 14, 8, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3,
	2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 77, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74,
	76, 5, 18, 10, 2, 75, 74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2,
	2, 2, 77, 78, 3, 2, 2, 2, 78, 83, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 82,
	5, 26, 14, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 83, 84, 3, 2, 2, 2, 84, 3, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7,
	4, 2, 2, 87, 88, 7, 3, 2, 2, 88, 89, 7, 15, 2, 2, 89, 5, 3, 2, 2, 2, 90,
	91, 7, 5, 2, 2, 91, 92, 5, 48, 25, 2, 92, 93, 7, 15, 2, 2, 93, 7, 3, 2,
	2, 2, 94, 95, 7, 6, 2, 2, 95, 96, 5, 46, 24, 2, 96, 97, 7, 15, 2, 2, 97,
	9, 3, 2, 2, 2, 98, 99, 7, 7, 2, 2, 99, 104, 5, 12, 7, 2, 100, 101, 7, 26,
	2, 2, 101, 103, 5, 12, 7, 2, 102, 100, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2,
	104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106,
	104, 3, 2, 2, 2, 107, 108, 7, 15, 2, 2, 108, 11, 3, 2, 2, 2, 109, 112,
	5, 48, 25, 2, 110, 111, 7, 17, 2, 2, 111, 113, 5, 48, 25, 2, 112, 110,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 118, 3, 2, 2, 2, 114, 115, 7, 24,
	2, 2, 115, 116, 5, 42, 22, 2, 116, 117, 7, 25, 2, 2, 117, 119, 3, 2, 2,
	2, 118, 114, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 13, 3, 2, 2, 2, 120,
	121, 7, 8, 2, 2, 121, 126, 5, 16, 9, 2, 122, 123, 7, 26, 2, 2, 123, 125,
	5, 16, 9, 2, 124, 122, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2,
	2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2,
	129, 130, 7, 15, 2, 2, 130, 15, 3, 2, 2, 2, 131, 132, 5, 48, 25, 2, 132,
	133, 7, 14, 2, 2, 133, 134, 7, 3, 2, 2, 134, 17, 3, 2, 2, 2, 135, 136,
	7, 13, 2, 2, 136, 141, 5, 24, 13, 2, 137, 138, 7, 26, 2, 2, 138, 140, 5,
	24, 13, 2, 139, 137, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2,
	2, 2, 141, 142, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2,
	144, 146, 5, 22, 12, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	154, 3, 2, 2, 2, 147, 155, 7, 15, 2, 2, 148, 155, 3, 2, 2, 2, 149, 150,
	5, 20, 11, 2, 150, 151, 7, 15, 2, 2, 151, 155, 3, 2, 2, 2, 152, 155, 3,
	2, 2, 2, 153, 155, 5, 34, 18, 2, 154, 147, 3, 2, 2, 2, 154, 148, 3, 2,
	2, 2, 154, 149, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2,
	155, 19, 3, 2, 2, 2, 156, 166, 7, 12, 2, 2, 157, 162, 5, 48, 25, 2, 158,
	159, 7, 22, 2, 2, 159, 160, 5, 42, 22, 2, 160, 161, 7, 23, 2, 2, 161, 163,
	3, 2, 2, 2, 162, 158, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 167, 3, 2,
	2, 2, 164, 167, 3, 2, 2, 2, 165, 167, 7, 3, 2, 2, 166, 157, 3, 2, 2, 2,
	166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 21, 3, 2, 2, 2, 168, 169,
	7, 11, 2, 2, 169, 170, 5, 42, 22, 2, 170, 23, 3, 2, 2, 2, 171, 176, 5,
	48, 25, 2, 172, 173, 7, 24, 2, 2, 173, 174, 5, 42, 22, 2, 174, 175, 7,
	25, 2, 2, 175, 177, 3, 2, 2, 2, 176, 172, 3, 2, 2, 2, 176, 177, 3, 2, 2,
	2, 177, 25, 3, 2, 2, 2, 178, 179, 7, 9, 2, 2, 179, 180, 5, 28, 15, 2, 180,
	27, 3, 2, 2, 2, 181, 186, 5, 48, 25, 2, 182, 183, 7, 24, 2, 2, 183, 184,
	5, 42, 22, 2, 184, 185, 7, 25, 2, 2, 185, 187, 3, 2, 2, 2, 186, 182, 3,
	2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 189, 7, 11, 2,
	2, 189, 191, 5, 48, 25, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2,
	191, 192, 3, 2, 2, 2, 192, 193, 5, 30, 16, 2, 193, 29, 3, 2, 2, 2, 194,
	196, 7, 19, 2, 2, 195, 197, 5, 32, 17, 2, 196, 195, 3, 2, 2, 2, 197, 198,
	3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2,
	2, 2, 200, 201, 7, 20, 2, 2, 201, 31, 3, 2, 2, 2, 202, 203, 7, 10, 2, 2,
	203, 204, 5, 48, 25, 2, 204, 205, 5, 34, 18, 2, 205, 33, 3, 2, 2, 2, 206,
	210, 7, 19, 2, 2, 207, 209, 5, 36, 19, 2, 208, 207, 3, 2, 2, 2, 209, 212,
	3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2,
	2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 7, 20, 2, 2, 214, 35, 3, 2, 2, 2,
	215, 216, 5, 38, 20, 2, 216, 225, 7, 12, 2, 2, 217, 222, 5, 48, 25, 2,
	218, 219, 7, 22, 2, 2, 219, 220, 5, 42, 22, 2, 220, 221, 7, 23, 2, 2, 221,
	223, 3, 2, 2, 2, 222, 218, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 226,
	3, 2, 2, 2, 224, 226, 7, 3, 2, 2, 225, 217, 3, 2, 2, 2, 225, 224, 3, 2,
	2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 7, 15, 2, 2, 228, 37, 3, 2, 2, 2,
	229, 234, 5, 40, 21, 2, 230, 231, 7, 26, 2, 2, 231, 233, 5, 40, 21, 2,
	232, 230, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 39, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 242, 5,
	48, 25, 2, 238, 239, 7, 24, 2, 2, 239, 240, 5, 42, 22, 2, 240, 241, 7,
	25, 2, 2, 241, 243, 3, 2, 2, 2, 242, 238, 3, 2, 2, 2, 242, 243, 3, 2, 2,
	2, 243, 41, 3, 2, 2, 2, 244, 249, 5, 44, 23, 2, 245, 246, 7, 26, 2, 2,
	246, 248, 5, 44, 23, 2, 247, 245, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2, 249,
	247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 43, 3, 2, 2, 2, 251, 249, 3,
	2, 2, 2, 252, 255, 5, 48, 25, 2, 253, 254, 7, 17, 2, 2, 254, 256, 5, 48,
	25, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 45, 3, 2, 2, 2,
	257, 259, 7, 27, 2, 2, 258, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260,
	258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 266, 3, 2, 2, 2, 262, 263,
	7, 21, 2, 2, 263, 265, 7, 27, 2, 2, 264, 262, 3, 2, 2, 2, 265, 268, 3,
	2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 47, 3, 2, 2,
	2, 268, 266, 3, 2, 2, 2, 269, 271, 7, 16, 2, 2, 270, 269, 3, 2, 2, 2, 270,
	271, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 274, 9, 2, 2, 2, 273, 272,
	3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2,
	2, 2, 276, 49, 3, 2, 2, 2, 33, 53, 57, 60, 65, 71, 77, 83, 104, 112, 118,
	126, 141, 145, 154, 162, 166, 176, 186, 190, 198, 210, 222, 225, 234, 242,
	249, 255, 260, 266, 270, 275,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'#include'", "'target'", "'namespace'", "'context'", "'alias'",
	"'index'", "'method'", "'in'", "'as'", "'component'", "'='", "';'", "'_'",
	"':'", "'\"'", "'{'", "'}'", "'.'", "'<'", "'>'", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "StringLiteral", "INCLUDE", "TARGET", "NAMESPACE", "CONTEXT", "ALIAS",
	"INDEX", "METHOD", "IN", "AS", "COMPONENT", "EQUALS", "SEMICOLON", "UNDERSCORE",
	"COLON", "DOUBLE_QUOTE", "LCURLEY", "RCURLEY", "PERIOD", "LESS_THAN", "GREATER_THAN",
	"LPAREN", "RPAREN", "COMMA", "WORD", "NUMBER", "COMMENT", "WS",
}

var ruleNames = []string{
	"document", "includeStatement", "targetStatement", "namespaceStatement",
	"contextStatement", "context", "aliasStatement", "alias", "componentStatement",
	"asStatement", "inStatement", "component", "entityIndexStatement", "entityIndex",
	"entityIndexMethodStatementList", "entityIndexMethodStatement", "memberStatementList",
	"memberStatement", "propertyList", "property", "keyValueList", "keyValue",
	"namespace", "identifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ProtonLangParser struct {
	*antlr.BaseParser
}

func NewProtonLangParser(input antlr.TokenStream) *ProtonLangParser {
	this := new(ProtonLangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ProtonLang.g4"

	return this
}

// ProtonLangParser tokens.
const (
	ProtonLangParserEOF           = antlr.TokenEOF
	ProtonLangParserStringLiteral = 1
	ProtonLangParserINCLUDE       = 2
	ProtonLangParserTARGET        = 3
	ProtonLangParserNAMESPACE     = 4
	ProtonLangParserCONTEXT       = 5
	ProtonLangParserALIAS         = 6
	ProtonLangParserINDEX         = 7
	ProtonLangParserMETHOD        = 8
	ProtonLangParserIN            = 9
	ProtonLangParserAS            = 10
	ProtonLangParserCOMPONENT     = 11
	ProtonLangParserEQUALS        = 12
	ProtonLangParserSEMICOLON     = 13
	ProtonLangParserUNDERSCORE    = 14
	ProtonLangParserCOLON         = 15
	ProtonLangParserDOUBLE_QUOTE  = 16
	ProtonLangParserLCURLEY       = 17
	ProtonLangParserRCURLEY       = 18
	ProtonLangParserPERIOD        = 19
	ProtonLangParserLESS_THAN     = 20
	ProtonLangParserGREATER_THAN  = 21
	ProtonLangParserLPAREN        = 22
	ProtonLangParserRPAREN        = 23
	ProtonLangParserCOMMA         = 24
	ProtonLangParserWORD          = 25
	ProtonLangParserNUMBER        = 26
	ProtonLangParserCOMMENT       = 27
	ProtonLangParserWS            = 28
)

// ProtonLangParser rules.
const (
	ProtonLangParserRULE_document                       = 0
	ProtonLangParserRULE_includeStatement               = 1
	ProtonLangParserRULE_targetStatement                = 2
	ProtonLangParserRULE_namespaceStatement             = 3
	ProtonLangParserRULE_contextStatement               = 4
	ProtonLangParserRULE_context                        = 5
	ProtonLangParserRULE_aliasStatement                 = 6
	ProtonLangParserRULE_alias                          = 7
	ProtonLangParserRULE_componentStatement             = 8
	ProtonLangParserRULE_asStatement                    = 9
	ProtonLangParserRULE_inStatement                    = 10
	ProtonLangParserRULE_component                      = 11
	ProtonLangParserRULE_entityIndexStatement           = 12
	ProtonLangParserRULE_entityIndex                    = 13
	ProtonLangParserRULE_entityIndexMethodStatementList = 14
	ProtonLangParserRULE_entityIndexMethodStatement     = 15
	ProtonLangParserRULE_memberStatementList            = 16
	ProtonLangParserRULE_memberStatement                = 17
	ProtonLangParserRULE_propertyList                   = 18
	ProtonLangParserRULE_property                       = 19
	ProtonLangParserRULE_keyValueList                   = 20
	ProtonLangParserRULE_keyValue                       = 21
	ProtonLangParserRULE_namespace                      = 22
	ProtonLangParserRULE_identifier                     = 23
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllIncludeStatement() []IIncludeStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIncludeStatementContext)(nil)).Elem())
	var tst = make([]IIncludeStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIncludeStatementContext)
		}
	}

	return tst
}

func (s *DocumentContext) IncludeStatement(i int) IIncludeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIncludeStatementContext)
}

func (s *DocumentContext) TargetStatement() ITargetStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetStatementContext)
}

func (s *DocumentContext) NamespaceStatement() INamespaceStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceStatementContext)
}

func (s *DocumentContext) AllContextStatement() []IContextStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContextStatementContext)(nil)).Elem())
	var tst = make([]IContextStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContextStatementContext)
		}
	}

	return tst
}

func (s *DocumentContext) ContextStatement(i int) IContextStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContextStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IContextStatementContext)
}

func (s *DocumentContext) AllAliasStatement() []IAliasStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasStatementContext)(nil)).Elem())
	var tst = make([]IAliasStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasStatementContext)
		}
	}

	return tst
}

func (s *DocumentContext) AliasStatement(i int) IAliasStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasStatementContext)
}

func (s *DocumentContext) AllComponentStatement() []IComponentStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentStatementContext)(nil)).Elem())
	var tst = make([]IComponentStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentStatementContext)
		}
	}

	return tst
}

func (s *DocumentContext) ComponentStatement(i int) IComponentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentStatementContext)
}

func (s *DocumentContext) AllEntityIndexStatement() []IEntityIndexStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntityIndexStatementContext)(nil)).Elem())
	var tst = make([]IEntityIndexStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntityIndexStatementContext)
		}
	}

	return tst
}

func (s *DocumentContext) EntityIndexStatement(i int) IEntityIndexStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityIndexStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEntityIndexStatementContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *ProtonLangParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProtonLangParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserINCLUDE {
		{
			p.SetState(48)
			p.IncludeStatement()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserTARGET {
		{
			p.SetState(54)
			p.TargetStatement()
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserNAMESPACE {
		{
			p.SetState(57)
			p.NamespaceStatement()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCONTEXT {
		{
			p.SetState(60)
			p.ContextStatement()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserALIAS {
		{
			p.SetState(66)
			p.AliasStatement()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMPONENT {
		{
			p.SetState(72)
			p.ComponentStatement()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserINDEX {
		{
			p.SetState(78)
			p.EntityIndexStatement()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIncludeStatementContext is an interface to support dynamic dispatch.
type IIncludeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeStatementContext differentiates from other interfaces.
	IsIncludeStatementContext()
}

type IncludeStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeStatementContext() *IncludeStatementContext {
	var p = new(IncludeStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_includeStatement
	return p
}

func (*IncludeStatementContext) IsIncludeStatementContext() {}

func NewIncludeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeStatementContext {
	var p = new(IncludeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_includeStatement

	return p
}

func (s *IncludeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeStatementContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserINCLUDE, 0)
}

func (s *IncludeStatementContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserStringLiteral, 0)
}

func (s *IncludeStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *IncludeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterIncludeStatement(s)
	}
}

func (s *IncludeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitIncludeStatement(s)
	}
}

func (p *ProtonLangParser) IncludeStatement() (localctx IIncludeStatementContext) {
	localctx = NewIncludeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProtonLangParserRULE_includeStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(ProtonLangParserINCLUDE)
	}
	{
		p.SetState(85)
		p.Match(ProtonLangParserStringLiteral)
	}
	{
		p.SetState(86)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// ITargetStatementContext is an interface to support dynamic dispatch.
type ITargetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetStatementContext differentiates from other interfaces.
	IsTargetStatementContext()
}

type TargetStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetStatementContext() *TargetStatementContext {
	var p = new(TargetStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_targetStatement
	return p
}

func (*TargetStatementContext) IsTargetStatementContext() {}

func NewTargetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetStatementContext {
	var p = new(TargetStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_targetStatement

	return p
}

func (s *TargetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetStatementContext) TARGET() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserTARGET, 0)
}

func (s *TargetStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TargetStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *TargetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterTargetStatement(s)
	}
}

func (s *TargetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitTargetStatement(s)
	}
}

func (p *ProtonLangParser) TargetStatement() (localctx ITargetStatementContext) {
	localctx = NewTargetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProtonLangParserRULE_targetStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ProtonLangParserTARGET)
	}
	{
		p.SetState(89)
		p.Identifier()
	}
	{
		p.SetState(90)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// INamespaceStatementContext is an interface to support dynamic dispatch.
type INamespaceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceStatementContext differentiates from other interfaces.
	IsNamespaceStatementContext()
}

type NamespaceStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceStatementContext() *NamespaceStatementContext {
	var p = new(NamespaceStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_namespaceStatement
	return p
}

func (*NamespaceStatementContext) IsNamespaceStatementContext() {}

func NewNamespaceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceStatementContext {
	var p = new(NamespaceStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_namespaceStatement

	return p
}

func (s *NamespaceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceStatementContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserNAMESPACE, 0)
}

func (s *NamespaceStatementContext) Namespace() INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *NamespaceStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *NamespaceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterNamespaceStatement(s)
	}
}

func (s *NamespaceStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitNamespaceStatement(s)
	}
}

func (p *ProtonLangParser) NamespaceStatement() (localctx INamespaceStatementContext) {
	localctx = NewNamespaceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProtonLangParserRULE_namespaceStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(ProtonLangParserNAMESPACE)
	}
	{
		p.SetState(93)
		p.Namespace()
	}
	{
		p.SetState(94)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// IContextStatementContext is an interface to support dynamic dispatch.
type IContextStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContextStatementContext differentiates from other interfaces.
	IsContextStatementContext()
}

type ContextStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContextStatementContext() *ContextStatementContext {
	var p = new(ContextStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_contextStatement
	return p
}

func (*ContextStatementContext) IsContextStatementContext() {}

func NewContextStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextStatementContext {
	var p = new(ContextStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_contextStatement

	return p
}

func (s *ContextStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContextStatementContext) CONTEXT() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCONTEXT, 0)
}

func (s *ContextStatementContext) AllContext() []IContextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IContextContext)(nil)).Elem())
	var tst = make([]IContextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IContextContext)
		}
	}

	return tst
}

func (s *ContextStatementContext) Context(i int) IContextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IContextContext)
}

func (s *ContextStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *ContextStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserCOMMA)
}

func (s *ContextStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMMA, i)
}

func (s *ContextStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContextStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterContextStatement(s)
	}
}

func (s *ContextStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitContextStatement(s)
	}
}

func (p *ProtonLangParser) ContextStatement() (localctx IContextStatementContext) {
	localctx = NewContextStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProtonLangParserRULE_contextStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(ProtonLangParserCONTEXT)
	}
	{
		p.SetState(97)
		p.Context()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMMA {
		{
			p.SetState(98)
			p.Match(ProtonLangParserCOMMA)
		}
		{
			p.SetState(99)
			p.Context()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// IContextContext is an interface to support dynamic dispatch.
type IContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContextContext differentiates from other interfaces.
	IsContextContext()
}

type ContextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContextContext() *ContextContext {
	var p = new(ContextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_context
	return p
}

func (*ContextContext) IsContextContext() {}

func NewContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextContext {
	var p = new(ContextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_context

	return p
}

func (s *ContextContext) GetParser() antlr.Parser { return s.parser }

func (s *ContextContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ContextContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ContextContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOLON, 0)
}

func (s *ContextContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLPAREN, 0)
}

func (s *ContextContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *ContextContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRPAREN, 0)
}

func (s *ContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterContext(s)
	}
}

func (s *ContextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitContext(s)
	}
}

func (p *ProtonLangParser) Context() (localctx IContextContext) {
	localctx = NewContextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProtonLangParserRULE_context)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Identifier()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserCOLON {
		{
			p.SetState(108)
			p.Match(ProtonLangParserCOLON)
		}
		{
			p.SetState(109)
			p.Identifier()
		}

	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserLPAREN {
		{
			p.SetState(112)
			p.Match(ProtonLangParserLPAREN)
		}
		{
			p.SetState(113)
			p.KeyValueList()
		}
		{
			p.SetState(114)
			p.Match(ProtonLangParserRPAREN)
		}

	}

	return localctx
}

// IAliasStatementContext is an interface to support dynamic dispatch.
type IAliasStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasStatementContext differentiates from other interfaces.
	IsAliasStatementContext()
}

type AliasStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasStatementContext() *AliasStatementContext {
	var p = new(AliasStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_aliasStatement
	return p
}

func (*AliasStatementContext) IsAliasStatementContext() {}

func NewAliasStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasStatementContext {
	var p = new(AliasStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_aliasStatement

	return p
}

func (s *AliasStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasStatementContext) ALIAS() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserALIAS, 0)
}

func (s *AliasStatementContext) AllAlias() []IAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasContext)(nil)).Elem())
	var tst = make([]IAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasContext)
		}
	}

	return tst
}

func (s *AliasStatementContext) Alias(i int) IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *AliasStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *AliasStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserCOMMA)
}

func (s *AliasStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMMA, i)
}

func (s *AliasStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterAliasStatement(s)
	}
}

func (s *AliasStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitAliasStatement(s)
	}
}

func (p *ProtonLangParser) AliasStatement() (localctx IAliasStatementContext) {
	localctx = NewAliasStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProtonLangParserRULE_aliasStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(ProtonLangParserALIAS)
	}
	{
		p.SetState(119)
		p.Alias()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMMA {
		{
			p.SetState(120)
			p.Match(ProtonLangParserCOMMA)
		}
		{
			p.SetState(121)
			p.Alias()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AliasContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserEQUALS, 0)
}

func (s *AliasContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserStringLiteral, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *ProtonLangParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProtonLangParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Identifier()
	}
	{
		p.SetState(130)
		p.Match(ProtonLangParserEQUALS)
	}
	{
		p.SetState(131)
		p.Match(ProtonLangParserStringLiteral)
	}

	return localctx
}

// IComponentStatementContext is an interface to support dynamic dispatch.
type IComponentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentStatementContext differentiates from other interfaces.
	IsComponentStatementContext()
}

type ComponentStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentStatementContext() *ComponentStatementContext {
	var p = new(ComponentStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_componentStatement
	return p
}

func (*ComponentStatementContext) IsComponentStatementContext() {}

func NewComponentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentStatementContext {
	var p = new(ComponentStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_componentStatement

	return p
}

func (s *ComponentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentStatementContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMPONENT, 0)
}

func (s *ComponentStatementContext) AllComponent() []IComponentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentContext)(nil)).Elem())
	var tst = make([]IComponentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentContext)
		}
	}

	return tst
}

func (s *ComponentStatementContext) Component(i int) IComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *ComponentStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *ComponentStatementContext) AsStatement() IAsStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsStatementContext)
}

func (s *ComponentStatementContext) MemberStatementList() IMemberStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberStatementListContext)
}

func (s *ComponentStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserCOMMA)
}

func (s *ComponentStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMMA, i)
}

func (s *ComponentStatementContext) InStatement() IInStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInStatementContext)
}

func (s *ComponentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterComponentStatement(s)
	}
}

func (s *ComponentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitComponentStatement(s)
	}
}

func (p *ProtonLangParser) ComponentStatement() (localctx IComponentStatementContext) {
	localctx = NewComponentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProtonLangParserRULE_componentStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(ProtonLangParserCOMPONENT)
	}
	{
		p.SetState(134)
		p.Component()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMMA {
		{
			p.SetState(135)
			p.Match(ProtonLangParserCOMMA)
		}
		{
			p.SetState(136)
			p.Component()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserIN {
		{
			p.SetState(142)
			p.InStatement()
		}

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(145)
			p.Match(ProtonLangParserSEMICOLON)
		}

	case 2:

	case 3:
		{
			p.SetState(147)
			p.AsStatement()
		}
		{
			p.SetState(148)
			p.Match(ProtonLangParserSEMICOLON)
		}

	case 4:

	case 5:
		{
			p.SetState(151)
			p.MemberStatementList()
		}

	}

	return localctx
}

// IAsStatementContext is an interface to support dynamic dispatch.
type IAsStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsStatementContext differentiates from other interfaces.
	IsAsStatementContext()
}

type AsStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsStatementContext() *AsStatementContext {
	var p = new(AsStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_asStatement
	return p
}

func (*AsStatementContext) IsAsStatementContext() {}

func NewAsStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsStatementContext {
	var p = new(AsStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_asStatement

	return p
}

func (s *AsStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AsStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserAS, 0)
}

func (s *AsStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AsStatementContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserStringLiteral, 0)
}

func (s *AsStatementContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLESS_THAN, 0)
}

func (s *AsStatementContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *AsStatementContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserGREATER_THAN, 0)
}

func (s *AsStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterAsStatement(s)
	}
}

func (s *AsStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitAsStatement(s)
	}
}

func (p *ProtonLangParser) AsStatement() (localctx IAsStatementContext) {
	localctx = NewAsStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProtonLangParserRULE_asStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(ProtonLangParserAS)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProtonLangParserUNDERSCORE, ProtonLangParserWORD, ProtonLangParserNUMBER:
		{
			p.SetState(155)
			p.Identifier()
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProtonLangParserLESS_THAN {
			{
				p.SetState(156)
				p.Match(ProtonLangParserLESS_THAN)
			}
			{
				p.SetState(157)
				p.KeyValueList()
			}
			{
				p.SetState(158)
				p.Match(ProtonLangParserGREATER_THAN)
			}

		}

	case ProtonLangParserSEMICOLON:

	case ProtonLangParserStringLiteral:
		{
			p.SetState(163)
			p.Match(ProtonLangParserStringLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInStatementContext is an interface to support dynamic dispatch.
type IInStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInStatementContext differentiates from other interfaces.
	IsInStatementContext()
}

type InStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInStatementContext() *InStatementContext {
	var p = new(InStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_inStatement
	return p
}

func (*InStatementContext) IsInStatementContext() {}

func NewInStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InStatementContext {
	var p = new(InStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_inStatement

	return p
}

func (s *InStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InStatementContext) IN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserIN, 0)
}

func (s *InStatementContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *InStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterInStatement(s)
	}
}

func (s *InStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitInStatement(s)
	}
}

func (p *ProtonLangParser) InStatement() (localctx IInStatementContext) {
	localctx = NewInStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProtonLangParserRULE_inStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(ProtonLangParserIN)
	}
	{
		p.SetState(167)
		p.KeyValueList()
	}

	return localctx
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_component
	return p
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ComponentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLPAREN, 0)
}

func (s *ComponentContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *ComponentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRPAREN, 0)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterComponent(s)
	}
}

func (s *ComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitComponent(s)
	}
}

func (p *ProtonLangParser) Component() (localctx IComponentContext) {
	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProtonLangParserRULE_component)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Identifier()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserLPAREN {
		{
			p.SetState(170)
			p.Match(ProtonLangParserLPAREN)
		}
		{
			p.SetState(171)
			p.KeyValueList()
		}
		{
			p.SetState(172)
			p.Match(ProtonLangParserRPAREN)
		}

	}

	return localctx
}

// IEntityIndexStatementContext is an interface to support dynamic dispatch.
type IEntityIndexStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityIndexStatementContext differentiates from other interfaces.
	IsEntityIndexStatementContext()
}

type EntityIndexStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityIndexStatementContext() *EntityIndexStatementContext {
	var p = new(EntityIndexStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_entityIndexStatement
	return p
}

func (*EntityIndexStatementContext) IsEntityIndexStatementContext() {}

func NewEntityIndexStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityIndexStatementContext {
	var p = new(EntityIndexStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_entityIndexStatement

	return p
}

func (s *EntityIndexStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityIndexStatementContext) INDEX() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserINDEX, 0)
}

func (s *EntityIndexStatementContext) EntityIndex() IEntityIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityIndexContext)
}

func (s *EntityIndexStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityIndexStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityIndexStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterEntityIndexStatement(s)
	}
}

func (s *EntityIndexStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitEntityIndexStatement(s)
	}
}

func (p *ProtonLangParser) EntityIndexStatement() (localctx IEntityIndexStatementContext) {
	localctx = NewEntityIndexStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ProtonLangParserRULE_entityIndexStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(ProtonLangParserINDEX)
	}
	{
		p.SetState(177)
		p.EntityIndex()
	}

	return localctx
}

// IEntityIndexContext is an interface to support dynamic dispatch.
type IEntityIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityIndexContext differentiates from other interfaces.
	IsEntityIndexContext()
}

type EntityIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityIndexContext() *EntityIndexContext {
	var p = new(EntityIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_entityIndex
	return p
}

func (*EntityIndexContext) IsEntityIndexContext() {}

func NewEntityIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityIndexContext {
	var p = new(EntityIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_entityIndex

	return p
}

func (s *EntityIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityIndexContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *EntityIndexContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EntityIndexContext) EntityIndexMethodStatementList() IEntityIndexMethodStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityIndexMethodStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityIndexMethodStatementListContext)
}

func (s *EntityIndexContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLPAREN, 0)
}

func (s *EntityIndexContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *EntityIndexContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRPAREN, 0)
}

func (s *EntityIndexContext) IN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserIN, 0)
}

func (s *EntityIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterEntityIndex(s)
	}
}

func (s *EntityIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitEntityIndex(s)
	}
}

func (p *ProtonLangParser) EntityIndex() (localctx IEntityIndexContext) {
	localctx = NewEntityIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ProtonLangParserRULE_entityIndex)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Identifier()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserLPAREN {
		{
			p.SetState(180)
			p.Match(ProtonLangParserLPAREN)
		}
		{
			p.SetState(181)
			p.KeyValueList()
		}
		{
			p.SetState(182)
			p.Match(ProtonLangParserRPAREN)
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserIN {
		{
			p.SetState(186)
			p.Match(ProtonLangParserIN)
		}
		{
			p.SetState(187)
			p.Identifier()
		}

	}
	{
		p.SetState(190)
		p.EntityIndexMethodStatementList()
	}

	return localctx
}

// IEntityIndexMethodStatementListContext is an interface to support dynamic dispatch.
type IEntityIndexMethodStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityIndexMethodStatementListContext differentiates from other interfaces.
	IsEntityIndexMethodStatementListContext()
}

type EntityIndexMethodStatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityIndexMethodStatementListContext() *EntityIndexMethodStatementListContext {
	var p = new(EntityIndexMethodStatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_entityIndexMethodStatementList
	return p
}

func (*EntityIndexMethodStatementListContext) IsEntityIndexMethodStatementListContext() {}

func NewEntityIndexMethodStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityIndexMethodStatementListContext {
	var p = new(EntityIndexMethodStatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_entityIndexMethodStatementList

	return p
}

func (s *EntityIndexMethodStatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityIndexMethodStatementListContext) LCURLEY() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLCURLEY, 0)
}

func (s *EntityIndexMethodStatementListContext) RCURLEY() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRCURLEY, 0)
}

func (s *EntityIndexMethodStatementListContext) AllEntityIndexMethodStatement() []IEntityIndexMethodStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntityIndexMethodStatementContext)(nil)).Elem())
	var tst = make([]IEntityIndexMethodStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntityIndexMethodStatementContext)
		}
	}

	return tst
}

func (s *EntityIndexMethodStatementListContext) EntityIndexMethodStatement(i int) IEntityIndexMethodStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityIndexMethodStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEntityIndexMethodStatementContext)
}

func (s *EntityIndexMethodStatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityIndexMethodStatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityIndexMethodStatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterEntityIndexMethodStatementList(s)
	}
}

func (s *EntityIndexMethodStatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitEntityIndexMethodStatementList(s)
	}
}

func (p *ProtonLangParser) EntityIndexMethodStatementList() (localctx IEntityIndexMethodStatementListContext) {
	localctx = NewEntityIndexMethodStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ProtonLangParserRULE_entityIndexMethodStatementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(ProtonLangParserLCURLEY)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ProtonLangParserMETHOD {
		{
			p.SetState(193)
			p.EntityIndexMethodStatement()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(198)
		p.Match(ProtonLangParserRCURLEY)
	}

	return localctx
}

// IEntityIndexMethodStatementContext is an interface to support dynamic dispatch.
type IEntityIndexMethodStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityIndexMethodStatementContext differentiates from other interfaces.
	IsEntityIndexMethodStatementContext()
}

type EntityIndexMethodStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityIndexMethodStatementContext() *EntityIndexMethodStatementContext {
	var p = new(EntityIndexMethodStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_entityIndexMethodStatement
	return p
}

func (*EntityIndexMethodStatementContext) IsEntityIndexMethodStatementContext() {}

func NewEntityIndexMethodStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityIndexMethodStatementContext {
	var p = new(EntityIndexMethodStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_entityIndexMethodStatement

	return p
}

func (s *EntityIndexMethodStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityIndexMethodStatementContext) METHOD() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserMETHOD, 0)
}

func (s *EntityIndexMethodStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EntityIndexMethodStatementContext) MemberStatementList() IMemberStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberStatementListContext)
}

func (s *EntityIndexMethodStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityIndexMethodStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityIndexMethodStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterEntityIndexMethodStatement(s)
	}
}

func (s *EntityIndexMethodStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitEntityIndexMethodStatement(s)
	}
}

func (p *ProtonLangParser) EntityIndexMethodStatement() (localctx IEntityIndexMethodStatementContext) {
	localctx = NewEntityIndexMethodStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ProtonLangParserRULE_entityIndexMethodStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(ProtonLangParserMETHOD)
	}
	{
		p.SetState(201)
		p.Identifier()
	}
	{
		p.SetState(202)
		p.MemberStatementList()
	}

	return localctx
}

// IMemberStatementListContext is an interface to support dynamic dispatch.
type IMemberStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberStatementListContext differentiates from other interfaces.
	IsMemberStatementListContext()
}

type MemberStatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberStatementListContext() *MemberStatementListContext {
	var p = new(MemberStatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_memberStatementList
	return p
}

func (*MemberStatementListContext) IsMemberStatementListContext() {}

func NewMemberStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberStatementListContext {
	var p = new(MemberStatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_memberStatementList

	return p
}

func (s *MemberStatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberStatementListContext) LCURLEY() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLCURLEY, 0)
}

func (s *MemberStatementListContext) RCURLEY() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRCURLEY, 0)
}

func (s *MemberStatementListContext) AllMemberStatement() []IMemberStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberStatementContext)(nil)).Elem())
	var tst = make([]IMemberStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberStatementContext)
		}
	}

	return tst
}

func (s *MemberStatementListContext) MemberStatement(i int) IMemberStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberStatementContext)
}

func (s *MemberStatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberStatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberStatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterMemberStatementList(s)
	}
}

func (s *MemberStatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitMemberStatementList(s)
	}
}

func (p *ProtonLangParser) MemberStatementList() (localctx IMemberStatementListContext) {
	localctx = NewMemberStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ProtonLangParserRULE_memberStatementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(ProtonLangParserLCURLEY)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtonLangParserUNDERSCORE)|(1<<ProtonLangParserWORD)|(1<<ProtonLangParserNUMBER))) != 0 {
		{
			p.SetState(205)
			p.MemberStatement()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(211)
		p.Match(ProtonLangParserRCURLEY)
	}

	return localctx
}

// IMemberStatementContext is an interface to support dynamic dispatch.
type IMemberStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberStatementContext differentiates from other interfaces.
	IsMemberStatementContext()
}

type MemberStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberStatementContext() *MemberStatementContext {
	var p = new(MemberStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_memberStatement
	return p
}

func (*MemberStatementContext) IsMemberStatementContext() {}

func NewMemberStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberStatementContext {
	var p = new(MemberStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_memberStatement

	return p
}

func (s *MemberStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberStatementContext) PropertyList() IPropertyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyListContext)
}

func (s *MemberStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserAS, 0)
}

func (s *MemberStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserSEMICOLON, 0)
}

func (s *MemberStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberStatementContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserStringLiteral, 0)
}

func (s *MemberStatementContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLESS_THAN, 0)
}

func (s *MemberStatementContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *MemberStatementContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserGREATER_THAN, 0)
}

func (s *MemberStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterMemberStatement(s)
	}
}

func (s *MemberStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitMemberStatement(s)
	}
}

func (p *ProtonLangParser) MemberStatement() (localctx IMemberStatementContext) {
	localctx = NewMemberStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ProtonLangParserRULE_memberStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.PropertyList()
	}
	{
		p.SetState(214)
		p.Match(ProtonLangParserAS)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProtonLangParserUNDERSCORE, ProtonLangParserWORD, ProtonLangParserNUMBER:
		{
			p.SetState(215)
			p.Identifier()
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProtonLangParserLESS_THAN {
			{
				p.SetState(216)
				p.Match(ProtonLangParserLESS_THAN)
			}
			{
				p.SetState(217)
				p.KeyValueList()
			}
			{
				p.SetState(218)
				p.Match(ProtonLangParserGREATER_THAN)
			}

		}

	case ProtonLangParserStringLiteral:
		{
			p.SetState(222)
			p.Match(ProtonLangParserStringLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(225)
		p.Match(ProtonLangParserSEMICOLON)
	}

	return localctx
}

// IPropertyListContext is an interface to support dynamic dispatch.
type IPropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyListContext differentiates from other interfaces.
	IsPropertyListContext()
}

type PropertyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyListContext() *PropertyListContext {
	var p = new(PropertyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_propertyList
	return p
}

func (*PropertyListContext) IsPropertyListContext() {}

func NewPropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListContext {
	var p = new(PropertyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_propertyList

	return p
}

func (s *PropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *PropertyListContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertyListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserCOMMA)
}

func (s *PropertyListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMMA, i)
}

func (s *PropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterPropertyList(s)
	}
}

func (s *PropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitPropertyList(s)
	}
}

func (p *ProtonLangParser) PropertyList() (localctx IPropertyListContext) {
	localctx = NewPropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ProtonLangParserRULE_propertyList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Property()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMMA {
		{
			p.SetState(228)
			p.Match(ProtonLangParserCOMMA)
		}
		{
			p.SetState(229)
			p.Property()
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserLPAREN, 0)
}

func (s *PropertyContext) KeyValueList() IKeyValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueListContext)
}

func (s *PropertyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserRPAREN, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *ProtonLangParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ProtonLangParserRULE_property)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Identifier()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserLPAREN {
		{
			p.SetState(236)
			p.Match(ProtonLangParserLPAREN)
		}
		{
			p.SetState(237)
			p.KeyValueList()
		}
		{
			p.SetState(238)
			p.Match(ProtonLangParserRPAREN)
		}

	}

	return localctx
}

// IKeyValueListContext is an interface to support dynamic dispatch.
type IKeyValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueListContext differentiates from other interfaces.
	IsKeyValueListContext()
}

type KeyValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueListContext() *KeyValueListContext {
	var p = new(KeyValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_keyValueList
	return p
}

func (*KeyValueListContext) IsKeyValueListContext() {}

func NewKeyValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueListContext {
	var p = new(KeyValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_keyValueList

	return p
}

func (s *KeyValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueListContext) AllKeyValue() []IKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValueContext)(nil)).Elem())
	var tst = make([]IKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValueContext)
		}
	}

	return tst
}

func (s *KeyValueListContext) KeyValue(i int) IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *KeyValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserCOMMA)
}

func (s *KeyValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOMMA, i)
}

func (s *KeyValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterKeyValueList(s)
	}
}

func (s *KeyValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitKeyValueList(s)
	}
}

func (p *ProtonLangParser) KeyValueList() (localctx IKeyValueListContext) {
	localctx = NewKeyValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ProtonLangParserRULE_keyValueList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.KeyValue()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserCOMMA {
		{
			p.SetState(243)
			p.Match(ProtonLangParserCOMMA)
		}
		{
			p.SetState(244)
			p.KeyValue()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *KeyValueContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *KeyValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProtonLangParserCOLON, 0)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (p *ProtonLangParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ProtonLangParserRULE_keyValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Identifier()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtonLangParserCOLON {
		{
			p.SetState(251)
			p.Match(ProtonLangParserCOLON)
		}
		{
			p.SetState(252)
			p.Identifier()
		}

	}

	return localctx
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_namespace
	return p
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserWORD)
}

func (s *NamespaceContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserWORD, i)
}

func (s *NamespaceContext) AllPERIOD() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserPERIOD)
}

func (s *NamespaceContext) PERIOD(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserPERIOD, i)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *ProtonLangParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ProtonLangParserRULE_namespace)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ProtonLangParserWORD {
		{
			p.SetState(255)
			p.Match(ProtonLangParserWORD)
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtonLangParserPERIOD {
		{
			p.SetState(260)
			p.Match(ProtonLangParserPERIOD)
		}
		{
			p.SetState(261)
			p.Match(ProtonLangParserWORD)
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtonLangParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtonLangParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllUNDERSCORE() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserUNDERSCORE)
}

func (s *IdentifierContext) UNDERSCORE(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserUNDERSCORE, i)
}

func (s *IdentifierContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserWORD)
}

func (s *IdentifierContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserWORD, i)
}

func (s *IdentifierContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ProtonLangParserNUMBER)
}

func (s *IdentifierContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ProtonLangParserNUMBER, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtonLangListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *ProtonLangParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ProtonLangParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(267)
			p.Match(ProtonLangParserUNDERSCORE)
		}

	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtonLangParserUNDERSCORE)|(1<<ProtonLangParserWORD)|(1<<ProtonLangParserNUMBER))) != 0) {
		{
			p.SetState(270)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtonLangParserUNDERSCORE)|(1<<ProtonLangParserWORD)|(1<<ProtonLangParserNUMBER))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
