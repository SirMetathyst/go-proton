<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventSystemArgumentPass_1_6_1(cp *model.CP, b *bytes.Buffer) string %>
<% ml := cp.MemberList()
if len(ml) > 0 && cp.EventType() == model.AddedEvent {
    b.WriteString(", ")
    for i, m := range ml {
        b.WriteString("component.")
        b.WriteString(m.ID().ToLowerFirst().String())
        if i != len(ml)-1 {
             b.WriteString(", ")
        }
    }
}
%><% return b.String() %>