<%: func EntityIndexArgumentPass_C_1_4_2(eim *entitas.EIM, b *bytes.Buffer) string %>
<% ml := eim.MemberList()
    for i, m := range ml {
        b.WriteString(m.ID().ToLowerFirst().String())
        if i != len(ml)-1 {
             b.WriteString(", ")
        }
    }
%><% return b.String() %>