// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\event_listener_interface\C_1_6_1\event_listener_interface.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func EventListenerInterface_C_1_6_1(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string {
	b.WriteString(`
public interface I`)
	b.WriteString(componentID(c, cp).ToUpperFirst().String())
	b.WriteString(` 
{
    void On`)
	b.WriteString(methodID(cp).ToUpperFirst().String())
	b.WriteString(`(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity entity`)
	EventListenerArgumentPass_1_6_1(cp, b)
	b.WriteString(`);
}
`)
	return b.String()

}