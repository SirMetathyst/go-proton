<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexGetCustomIndices_C_1_4_2(ei []*model.EI, b *bytes.Buffer) string %>
<% for _, cei := range ei { 

    IndexType := cei.ID().String()
    
    ReturnType := ""
    if !cei.IsPrimary() {
       ReturnType = "System.Collections.Generic.HashSet<" + cei.Context().ID().WithoutContextSuffix().ToUpperFirst().String() + "Entity>" 
    } else {
        ReturnType = cei.Context().ID().WithoutContextSuffix().ToUpperFirst().String() + "Entity"
    }
    
    for _, eim := range cei.EntityIndexMethodList() {
%>
public static <%==s ReturnType%><% b.WriteRune(' ')%><%==s eim.ID().String()%>(this <%==s cei.Context().ID().WithContextSuffix().ToUpperFirst().String()%> context, <% EntityIndexArgument_C_1_4_2(eim, b) %>) 
{
    return ((<%==s IndexType%>)(context.GetEntityIndex(Contexts.<%==s IndexType%>))).<%==s eim.ID().String()%>(<% EntityIndexArgumentPass_C_1_4_2(eim, b) %>);
}

<% }} %>
<% return b.String() %>