<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexGetCustomIndices_C_1_4_2(ei []*model.EntityIndex, b *bytes.Buffer) string %>
<% for _, e := range ei { 
    if e.GetContext() != nil {
    IndexType := e.GetID().String()
    
    ReturnType := ""
    if !e.IsPrimary() {
       ReturnType = "System.Collections.Generic.HashSet<" + e.GetContext().GetID().WithoutContextSuffix().ToUpperFirst().String() + "Entity>" 
    } else {
        ReturnType = e.GetContext().GetID().WithoutContextSuffix().ToUpperFirst().String() + "Entity"
    }
    
    for _, eim := range e.GetEntityIndexMethod() {
%>
public static <%==s ReturnType%><% b.WriteRune(' ')%><%==s eim.GetID().String()%>(this <%==s e.GetContext().GetID().WithContextSuffix().ToUpperFirst().String()%> context, <% EntityIndexArgument_C_1_4_2(eim, b) %>) {
        return ((<%==s IndexType%>)(context.GetEntityIndex(Contexts.<%==s IndexType%>))).<%==s eim.GetID().String()%>(<% EntityIndexArgumentPass_C_1_4_2(eim, b) %>);
}

<% }}} %>
<% return b.String() %>