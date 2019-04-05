// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\context_matcher\C_1_4_2\context_matcher.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ContextMatcher_C_1_4_2(c *proton.C, b *bytes.Buffer) string {
	b.WriteString(`
public sealed partial class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Matcher 
{
    public static Entitas.IAllOfMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> AllOf(params int[] Index) => Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>.AllOf(Index);
    public static Entitas.IAllOfMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> AllOf(params Entitas.IMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>[] Matcher) => Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>.AllOf(Matcher);
    public static Entitas.IAnyOfMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> AnyOf(params int[] Index) => Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>.AnyOf(Index);
    public static Entitas.IAnyOfMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> AnyOf(params Entitas.IMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>[] Matcher) => Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>.AnyOf(Matcher);
}
`)
	return b.String()

}
