<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexArgument_C_1_4_2(eim *model.EntityIndexMethod, b *bytes.Buffer) string %>
<% ms := eim.GetMember()
    for i, m := range ms {
        b.WriteString(m.GetValue().String())
        b.WriteRune(' ')
        b.WriteString(m.GetID().ToLowerFirst().String())
        if i != len(ms)-1 {
             b.WriteString(", ")
        }
    }
%><% return b.String() %>