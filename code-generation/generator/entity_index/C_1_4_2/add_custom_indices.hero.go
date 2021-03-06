// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\entity_index\C_1_4_2\add_custom_indices.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func EntityIndexAddCustomIndices_C_1_4_2(ei []*proton.EntityIndex, b *bytes.Buffer) string {

	for _, cei := range ei {
		if cei.Context() != nil {
			b.WriteString("\t\t")
			b.WriteString(cei.Context().ID().WithoutContextSuffix().ToLowerFirst().String())
			b.WriteRune('.')
			b.WriteString("AddEntityIndex(new ")
			b.WriteString(cei.ID().String())
			b.WriteRune('(')
			b.WriteString(cei.Context().ID().WithoutContextSuffix().ToLowerFirst().String())
			b.WriteString("));\n")
		}
	}
	return b.String()

}
