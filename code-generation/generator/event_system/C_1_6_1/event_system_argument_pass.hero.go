// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\event_system\C_1_6_1\event_system_argument_pass.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func EventSystemArgumentPass_1_6_1(cp *proton.Component, b *bytes.Buffer) string {
	ml := cp.MemberSlice()
	if len(ml) > 0 && cp.EventType() == proton.EventTypeAdded {
		b.WriteString(", ")
		for i, m := range ml {
			b.WriteString("component.")
			b.WriteString(m.ID().ToLowerFirst().String())
			if i != len(ml)-1 {
				b.WriteString(", ")
			}
		}
	}

	return b.String()

}
