// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_lookup\C_1_4_2\component_lookup.hero
// DO NOT EDIT!
package generator

import (
	"bytes"
	"strconv"

	entitas "github.com/SirMetathyst/go-entitas"
)

func ComponentLookup_C_1_4_2(c *entitas.C, cp []*entitas.CP, b *bytes.Buffer) string {
	b.WriteString(`
public static class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup 
{
`)

	i := 0
	ci := 0
	for _, ccp := range cp {
		b.WriteString("\tpublic const int ")
		b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString(" = ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString(";\n")
		i++
		ci++

		if ccp.EventTarget() != entitas.NoTarget {
			b.WriteString("\tpublic const int ")
			b.WriteString(eventComponentID(c, ccp).String())
			b.WriteString(" = ")
			b.WriteString(strconv.Itoa(i))
			b.WriteString(";\n")
			i++
			ci++
		}
	}
	b.WriteString(`
    public const int TotalComponents = `)
	b.WriteString(strconv.Itoa(i))
	b.WriteString(`;
	
	public static readonly string[] componentNames = 
    {
`)

	i = 0
	for _, ccp := range cp {
		b.WriteString("\t\t\"")
		b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString("\"")
		if i != ci-1 {
			b.WriteString(",\n")
		}
		i++
		if ccp.EventTarget() != entitas.NoTarget {
			b.WriteString("\t\t\"")
			b.WriteString(eventComponentID(c, ccp).String())
			b.WriteString("\"")
			if i != ci-1 {
				b.WriteString(",\n")
			}
			i++
		}
	}

	b.WriteString(`
	};
	
	public static readonly System.Type[] componentTypes = 
    {
`)

	i = 0
	for _, ccp := range cp {
		b.WriteString("\t\ttypeof(")
		b.WriteString(ccp.ID().WithComponentSuffix().ToUpperFirst().String())
		b.WriteRune(')')
		if i != ci-1 {
			b.WriteString(",\n")
		}
		i++

		if ccp.EventTarget() != entitas.NoTarget {
			b.WriteString("\t\ttypeof(")
			b.WriteString(eventComponentID(c, ccp).WithComponentSuffix().String())
			b.WriteRune(')')
			if i != ci-1 {
				b.WriteString(",\n")
			}
			i++
		}
	}

	b.WriteString(`
	};
}
`)
	return b.String()

}