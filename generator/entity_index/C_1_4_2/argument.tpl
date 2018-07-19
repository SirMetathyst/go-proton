<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexArgument_C_1_4_2(eim *model.EIM, b *bytes.Buffer) string %>
<% ml := eim.MemberList()
    for i, m := range ml {
        b.WriteString(m.Value().String())
        b.WriteRune(' ')
        b.WriteString(m.ID().ToLowerFirst().String())
        if i != len(ml)-1 {
             b.WriteString(", ")
        }
    }
%><% return b.String() %>