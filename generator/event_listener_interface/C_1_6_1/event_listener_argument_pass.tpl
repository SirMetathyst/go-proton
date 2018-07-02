<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerArgumentPass_1_6_1(c *model.Component, b *bytes.Buffer) string %>
<% ms := c.GetMember()
if len(ms) > 0 && c.GetEventType() != 1 {
    b.WriteString(", ")
    for i, m := range ms {
        b.WriteString(m.GetValue().String())
        b.WriteRune(' ')
        b.WriteString(m.GetID().ToLowerFirst().String())
        if i != len(ms)-1 {
             b.WriteString(", ")
        }
    }
}
%><% return b.String() %>