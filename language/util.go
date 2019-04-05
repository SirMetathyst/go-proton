package language

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Required ...
func Required(t antlr.ParseTree, e error) string {
	if t != nil {
		if strings.HasPrefix(t.GetText(), "\"") {
			return strings.Trim(t.GetText(), "\"")
		}
		return t.GetText()
	}
	panic(e)
}

// Optional ...
func Optional(t antlr.ParseTree, v string) string {
	if t != nil {
		if strings.HasPrefix(t.GetText(), "\"") {
			return strings.Trim(t.GetText(), "\"")
		}
		return t.GetText()
	}
	return v
}
