// Code generated from ProtonLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 196,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 5, 2, 66, 10,
	2, 3, 2, 3, 2, 3, 3, 6, 3, 71, 10, 3, 13, 3, 14, 3, 72, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 6, 28, 173, 10, 28, 13, 28,
	14, 28, 174, 3, 29, 6, 29, 178, 10, 29, 13, 29, 14, 29, 179, 3, 30, 3,
	30, 3, 30, 3, 30, 7, 30, 186, 10, 30, 12, 30, 14, 30, 189, 11, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 2, 2, 32, 3, 3, 5, 2, 7, 2, 9, 4, 11,
	5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14,
	31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23,
	49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 3, 2, 6, 6, 2,
	12, 12, 15, 15, 36, 36, 94, 94, 4, 2, 67, 92, 99, 124, 4, 2, 12, 12, 15,
	15, 5, 2, 12, 12, 15, 15, 34, 34, 2, 198, 2, 3, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 3, 63, 3,
	2, 2, 2, 5, 70, 3, 2, 2, 2, 7, 74, 3, 2, 2, 2, 9, 76, 3, 2, 2, 2, 11, 85,
	3, 2, 2, 2, 13, 92, 3, 2, 2, 2, 15, 102, 3, 2, 2, 2, 17, 110, 3, 2, 2,
	2, 19, 116, 3, 2, 2, 2, 21, 122, 3, 2, 2, 2, 23, 129, 3, 2, 2, 2, 25, 132,
	3, 2, 2, 2, 27, 135, 3, 2, 2, 2, 29, 145, 3, 2, 2, 2, 31, 147, 3, 2, 2,
	2, 33, 149, 3, 2, 2, 2, 35, 151, 3, 2, 2, 2, 37, 153, 3, 2, 2, 2, 39, 155,
	3, 2, 2, 2, 41, 157, 3, 2, 2, 2, 43, 159, 3, 2, 2, 2, 45, 161, 3, 2, 2,
	2, 47, 163, 3, 2, 2, 2, 49, 165, 3, 2, 2, 2, 51, 167, 3, 2, 2, 2, 53, 169,
	3, 2, 2, 2, 55, 172, 3, 2, 2, 2, 57, 177, 3, 2, 2, 2, 59, 181, 3, 2, 2,
	2, 61, 192, 3, 2, 2, 2, 63, 65, 5, 37, 19, 2, 64, 66, 5, 5, 3, 2, 65, 64,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 5, 37, 19,
	2, 68, 4, 3, 2, 2, 2, 69, 71, 5, 7, 4, 2, 70, 69, 3, 2, 2, 2, 71, 72, 3,
	2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 6, 3, 2, 2, 2, 74,
	75, 10, 2, 2, 2, 75, 8, 3, 2, 2, 2, 76, 77, 7, 37, 2, 2, 77, 78, 7, 107,
	2, 2, 78, 79, 7, 112, 2, 2, 79, 80, 7, 101, 2, 2, 80, 81, 7, 110, 2, 2,
	81, 82, 7, 119, 2, 2, 82, 83, 7, 102, 2, 2, 83, 84, 7, 103, 2, 2, 84, 10,
	3, 2, 2, 2, 85, 86, 7, 118, 2, 2, 86, 87, 7, 99, 2, 2, 87, 88, 7, 116,
	2, 2, 88, 89, 7, 105, 2, 2, 89, 90, 7, 103, 2, 2, 90, 91, 7, 118, 2, 2,
	91, 12, 3, 2, 2, 2, 92, 93, 7, 112, 2, 2, 93, 94, 7, 99, 2, 2, 94, 95,
	7, 111, 2, 2, 95, 96, 7, 103, 2, 2, 96, 97, 7, 117, 2, 2, 97, 98, 7, 114,
	2, 2, 98, 99, 7, 99, 2, 2, 99, 100, 7, 101, 2, 2, 100, 101, 7, 103, 2,
	2, 101, 14, 3, 2, 2, 2, 102, 103, 7, 101, 2, 2, 103, 104, 7, 113, 2, 2,
	104, 105, 7, 112, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107, 7, 103, 2, 2,
	107, 108, 7, 122, 2, 2, 108, 109, 7, 118, 2, 2, 109, 16, 3, 2, 2, 2, 110,
	111, 7, 99, 2, 2, 111, 112, 7, 110, 2, 2, 112, 113, 7, 107, 2, 2, 113,
	114, 7, 99, 2, 2, 114, 115, 7, 117, 2, 2, 115, 18, 3, 2, 2, 2, 116, 117,
	7, 107, 2, 2, 117, 118, 7, 112, 2, 2, 118, 119, 7, 102, 2, 2, 119, 120,
	7, 103, 2, 2, 120, 121, 7, 122, 2, 2, 121, 20, 3, 2, 2, 2, 122, 123, 7,
	111, 2, 2, 123, 124, 7, 103, 2, 2, 124, 125, 7, 118, 2, 2, 125, 126, 7,
	106, 2, 2, 126, 127, 7, 113, 2, 2, 127, 128, 7, 102, 2, 2, 128, 22, 3,
	2, 2, 2, 129, 130, 7, 107, 2, 2, 130, 131, 7, 112, 2, 2, 131, 24, 3, 2,
	2, 2, 132, 133, 7, 99, 2, 2, 133, 134, 7, 117, 2, 2, 134, 26, 3, 2, 2,
	2, 135, 136, 7, 101, 2, 2, 136, 137, 7, 113, 2, 2, 137, 138, 7, 111, 2,
	2, 138, 139, 7, 114, 2, 2, 139, 140, 7, 113, 2, 2, 140, 141, 7, 112, 2,
	2, 141, 142, 7, 103, 2, 2, 142, 143, 7, 112, 2, 2, 143, 144, 7, 118, 2,
	2, 144, 28, 3, 2, 2, 2, 145, 146, 7, 63, 2, 2, 146, 30, 3, 2, 2, 2, 147,
	148, 7, 61, 2, 2, 148, 32, 3, 2, 2, 2, 149, 150, 7, 97, 2, 2, 150, 34,
	3, 2, 2, 2, 151, 152, 7, 60, 2, 2, 152, 36, 3, 2, 2, 2, 153, 154, 7, 36,
	2, 2, 154, 38, 3, 2, 2, 2, 155, 156, 7, 125, 2, 2, 156, 40, 3, 2, 2, 2,
	157, 158, 7, 127, 2, 2, 158, 42, 3, 2, 2, 2, 159, 160, 7, 48, 2, 2, 160,
	44, 3, 2, 2, 2, 161, 162, 7, 62, 2, 2, 162, 46, 3, 2, 2, 2, 163, 164, 7,
	64, 2, 2, 164, 48, 3, 2, 2, 2, 165, 166, 7, 42, 2, 2, 166, 50, 3, 2, 2,
	2, 167, 168, 7, 43, 2, 2, 168, 52, 3, 2, 2, 2, 169, 170, 7, 46, 2, 2, 170,
	54, 3, 2, 2, 2, 171, 173, 9, 3, 2, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3,
	2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 56, 3, 2, 2,
	2, 176, 178, 4, 50, 59, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2,
	179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 58, 3, 2, 2, 2, 181, 182,
	7, 49, 2, 2, 182, 183, 7, 49, 2, 2, 183, 187, 3, 2, 2, 2, 184, 186, 10,
	4, 2, 2, 185, 184, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2,
	2, 187, 188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190,
	191, 8, 30, 2, 2, 191, 60, 3, 2, 2, 2, 192, 193, 9, 5, 2, 2, 193, 194,
	3, 2, 2, 2, 194, 195, 8, 31, 2, 2, 195, 62, 3, 2, 2, 2, 8, 2, 65, 72, 174,
	179, 187, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'#include'", "'target'", "'namespace'", "'context'", "'alias'",
	"'index'", "'method'", "'in'", "'as'", "'component'", "'='", "';'", "'_'",
	"':'", "'\"'", "'{'", "'}'", "'.'", "'<'", "'>'", "'('", "')'", "','",
}

var lexerSymbolicNames = []string{
	"", "StringLiteral", "INCLUDE", "TARGET", "NAMESPACE", "CONTEXT", "ALIAS",
	"INDEX", "METHOD", "IN", "AS", "COMPONENT", "EQUALS", "SEMICOLON", "UNDERSCORE",
	"COLON", "DOUBLE_QUOTE", "LCURLEY", "RCURLEY", "PERIOD", "LESS_THAN", "GREATER_THAN",
	"LPAREN", "RPAREN", "COMMA", "WORD", "NUMBER", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"StringLiteral", "StringCharacters", "StringCharacter", "INCLUDE", "TARGET",
	"NAMESPACE", "CONTEXT", "ALIAS", "INDEX", "METHOD", "IN", "AS", "COMPONENT",
	"EQUALS", "SEMICOLON", "UNDERSCORE", "COLON", "DOUBLE_QUOTE", "LCURLEY",
	"RCURLEY", "PERIOD", "LESS_THAN", "GREATER_THAN", "LPAREN", "RPAREN", "COMMA",
	"WORD", "NUMBER", "COMMENT", "WS",
}

type ProtonLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewProtonLangLexer(input antlr.CharStream) *ProtonLangLexer {

	l := new(ProtonLangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ProtonLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ProtonLangLexer tokens.
const (
	ProtonLangLexerStringLiteral = 1
	ProtonLangLexerINCLUDE       = 2
	ProtonLangLexerTARGET        = 3
	ProtonLangLexerNAMESPACE     = 4
	ProtonLangLexerCONTEXT       = 5
	ProtonLangLexerALIAS         = 6
	ProtonLangLexerINDEX         = 7
	ProtonLangLexerMETHOD        = 8
	ProtonLangLexerIN            = 9
	ProtonLangLexerAS            = 10
	ProtonLangLexerCOMPONENT     = 11
	ProtonLangLexerEQUALS        = 12
	ProtonLangLexerSEMICOLON     = 13
	ProtonLangLexerUNDERSCORE    = 14
	ProtonLangLexerCOLON         = 15
	ProtonLangLexerDOUBLE_QUOTE  = 16
	ProtonLangLexerLCURLEY       = 17
	ProtonLangLexerRCURLEY       = 18
	ProtonLangLexerPERIOD        = 19
	ProtonLangLexerLESS_THAN     = 20
	ProtonLangLexerGREATER_THAN  = 21
	ProtonLangLexerLPAREN        = 22
	ProtonLangLexerRPAREN        = 23
	ProtonLangLexerCOMMA         = 24
	ProtonLangLexerWORD          = 25
	ProtonLangLexerNUMBER        = 26
	ProtonLangLexerCOMMENT       = 27
	ProtonLangLexerWS            = 28
)
