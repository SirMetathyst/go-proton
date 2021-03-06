// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_entity\C_1_4_2\component_entity.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ComponentEntity_C_1_4_2(c *proton.Context, cp *proton.Component, isEventComponent bool, b *bytes.Buffer) string {

	ID := proton.String("")

	if isEventComponent {
		ID = eventComponentID(c, cp)
	} else {
		ID = cp.ID()
	}

	b.WriteString(`
public partial class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity
{
    public `)
	b.WriteString(ID.WithComponentSuffix().ToUpperFirst().String())
	b.WriteRune(' ')
	b.WriteString(ID.WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`
    {
        get
        {
            return (`)
	b.WriteString(ID.WithComponentSuffix().ToUpperFirst().String())
	b.WriteString(`)GetComponent(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`);
        }
    }

    public bool has`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`
    {
        get
        {
            return HasComponent(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`);
        }
    }

    public void Add`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentEntityArgument_C_1_4_2(cp, isEventComponent, b)
	b.WriteString(`)
    {
        var index = `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`;
        var component = CreateComponent<`)
	b.WriteString(ID.WithComponentSuffix().ToUpperFirst().String())
	b.WriteString(`>(index);
`)
	ComponentEntityAssignment_C_1_4_2(cp, isEventComponent, b)
	b.WriteString(`
        AddComponent(index, component);
    }

    public void Replace`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentEntityArgument_C_1_4_2(cp, isEventComponent, b)
	b.WriteString(`)
    {
        var index = `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`;
        var component = CreateComponent<`)
	b.WriteString(ID.WithComponentSuffix().ToUpperFirst().String())
	b.WriteString(`>(index);
`)
	ComponentEntityAssignment_C_1_4_2(cp, isEventComponent, b)
	b.WriteString(`
        ReplaceComponent(index, component);
    }

    public void Remove`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`()
    {
        RemoveComponent(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`);
    }
}
`)
	return b.String()

}
