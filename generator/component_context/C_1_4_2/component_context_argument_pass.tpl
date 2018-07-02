<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentContextArgumentPass_1_4_2(c *model.Component, b *bytes.Buffer) string %>
<% ms := c.GetMember()
    for i, m := range ms {
        b.WriteString("new")
        b.WriteString(m.GetID().ToUpperFirst().String())
        if i != len(ms)-1 {
             b.WriteString(", ")
        }
    }
%><% return b.String() %>
