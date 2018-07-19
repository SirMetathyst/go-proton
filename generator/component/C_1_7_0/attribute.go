package generator

import (
	"bytes"
	"strings"

	"github.com/SirMetathyst/proton/model"
)

// EntityIndexAttribute_C_1_7_0 ...
func EntityIndexAttribute_C_1_7_0(m *model.M) string {
	b := new(bytes.Buffer)
	if m.EntityIndex() == model.SingleEntityIndex {
		b.WriteString("Entitas.CodeGeneration.Attributes.PrimaryEntityIndex")
	} else if m.EntityIndex() == model.MultipleEntityIndex {
		b.WriteString("Entitas.CodeGeneration.Attributes.EntityIndex")
	}
	return b.String()
}

// CleanupModeAttribute_C_1_7_0 ...
func CleanupModeAttribute_C_1_7_0(cp *model.CP) string {
	b := new(bytes.Buffer)
	if cp.IsCleanup() {
		b.WriteString("Entitas.CodeGeneration.Attributes.Cleanup(")
		b.WriteString("Entitas.CodeGeneration.Attributes.CleanupMode.")
		b.WriteString(cp.CleanupMode().String().ToUpperFirst().String())
		b.WriteRune(')')
	}
	return b.String()
}

// EventAttribute_C_1_7_0 ...
func EventAttribute_C_1_7_0(cp *model.CP) string {
	b := new(bytes.Buffer)
	if cp.IsEvent() {
		b.WriteString("Entitas.CodeGeneration.Attributes.Event(")
		b.WriteString("Entitas.CodeGeneration.Attributes.EventTarget.")
		b.WriteString(cp.EventTarget().String().ToUpperFirst().String())
		if cp.EventType() == model.RemovedEvent {
			b.WriteString(", ")
			b.WriteString("Entitas.CodeGeneration.Attributes.EventType.")
			b.WriteString(cp.EventType().String().ToUpperFirst().String())
		}
		b.WriteRune(')')
	}
	return b.String()
}

// UniqueAttribute_C_1_7_0 ...
func UniqueAttribute_C_1_7_0(cp *model.CP) string {
	b := new(bytes.Buffer)
	if cp.IsUnique() {
		b.WriteString("Entitas.CodeGeneration.Attributes.Unique")
	}
	return b.String()
}

// FlagPrefixAttribute_C_1_7_0 ...
func FlagPrefixAttribute_C_1_7_0(cp *model.CP) string {
	b := new(bytes.Buffer)
	if cp.HasFlagPrefix() {
		b.WriteString("Entitas.CodeGeneration.Attributes.FlagPrefix(\"")
		b.WriteString(cp.FlagPrefix().String())
		b.WriteString("\")")
	}
	return b.String()
}

// ContextAttribute_C_1_7_0 ...
func ContextAttribute_C_1_7_0(cl []*model.C) string {
	b := new(bytes.Buffer)
	for i, c := range cl {
		b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
		if i != len(cl)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}

// ComponentAttribute_C_1_7_0 ...
func ComponentAttribute_C_1_7_0(cp *model.CP) string {
	slice := make([]string, 0)
	out := new(bytes.Buffer)

	append := func(s string) {
		if s != "" {
			slice = append(slice, s)
		}
	}

	append(ContextAttribute_C_1_7_0(cp.ContextList()))
	append(UniqueAttribute_C_1_7_0(cp))
	append(FlagPrefixAttribute_C_1_7_0(cp))
	append(EventAttribute_C_1_7_0(cp))
	append(CleanupModeAttribute_C_1_7_0(cp))

	if len(slice) > 0 {
		out.WriteRune('[')
		out.WriteString(strings.Join(slice, ", "))
		out.WriteRune(']')
	}
	return out.String()
}

// MemberAttribute_C_1_7_0 ...
func MemberAttribute_C_1_7_0(m *model.M) string {
	slice := make([]string, 0)
	out := new(bytes.Buffer)

	append := func(s string) {
		if s != "" {
			slice = append(slice, s)
		}
	}

	append(EntityIndexAttribute_C_1_7_0(m))

	if len(slice) > 0 {
		out.WriteRune('[')
		out.WriteString(strings.Join(slice, ", "))
		out.WriteRune(']')
	}
	return out.String()
}
