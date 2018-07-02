<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexAddCustomIndices_C_1_4_2(ei []*model.EntityIndex, b *bytes.Buffer) string %>
<% for _, e := range ei { 
    if e.GetContext() != nil {%>
<% b.WriteString("\t\t")%><%==s e.GetContext().GetID().WithoutContextSuffix().ToLowerFirst().String() %>.AddEntityIndex(new <%==s e.GetID().String()%>(<%==s e.GetContext().GetID().WithoutContextSuffix().ToLowerFirst().String()%>));<% b.WriteRune('\n')%>

<% }} %>
<% return b.String() %>