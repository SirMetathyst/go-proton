<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexCustomConstants_C_1_4_2(ei []*model.EntityIndex, b *bytes.Buffer) string %>
<% for _, e := range ei { 
    ID := e.GetID().String()
    b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%> = "<%==s ID%>";<% b.WriteString("\n") } %>
<% return b.String() %>