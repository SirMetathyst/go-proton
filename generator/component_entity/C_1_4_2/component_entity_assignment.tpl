<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityAssignment_C_1_4_2(c *model.Component, b *bytes.Buffer) string %>
<% ms := c.GetMember()
    for i, m := range ms {
        b.WriteString("\t\tcomponent.")
        b.WriteString(m.GetID().ToLowerFirst().String())
        b.WriteString(" = ")
        b.WriteString("new")
        b.WriteString(m.GetID().ToUpperFirst().String())
        b.WriteRune(';')
        if i != len(ms)-1 {
             b.WriteString("\n")
        }
    }
%><% return b.String() %>