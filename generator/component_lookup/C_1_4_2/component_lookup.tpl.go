// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\component_lookup\C_1_4_2\component_lookup.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"
	"strconv"

	"github.com/SirMetathyst/proton/model"
)

func ComponentLookup_C_1_4_2(ctx *model.Context, c []*model.Component, b *bytes.Buffer) string {
	b.WriteString(`public static class `)
	b.WriteString(ctx.GetID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup {
`)
	for i, cc := range c {
		b.WriteString("\tpublic const int ")
		b.WriteString(cc.GetID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString(" = ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString(";\n")
	}

	b.WriteString(`
	
    public const int TotalComponents = `)
	b.WriteString(strconv.Itoa(len(c)))
	b.WriteString(`;
	
	public static readonly string[] componentNames = {
`)

	for i, cc := range c {
		b.WriteString("\t\t\"")
		b.WriteString(cc.GetID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString("\"")
		if i != len(c)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString(`
	};
	
	public static readonly System.Type[] componentTypes = {
`)

	for i, cc := range c {
		b.WriteString("\t\ttypeof(")
		b.WriteString(cc.GetID().WithComponentSuffix().ToUpperFirst().String())
		b.WriteRune(')')
		if i != len(c)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString(`
	};
}
`)
	return b.String()

}
