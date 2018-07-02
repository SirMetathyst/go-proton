<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexConstants_C_1_4_2(c []*model.Component, b *bytes.Buffer) string %>
<% for _, cc := range c { 
    ID := cc.GetID().WithoutComponentSuffix().ToUpperFirst().String()
    if cc.GetEntityIndexCount() == 1 {
%><% b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%> = "<%==s ID%>";<% b.WriteString("\n")%><% } else if cc.GetEntityIndexCount() > 1 {
for _, m := range cc.GetMember() { 
%><% b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%><%==s m.GetID().ToUpperFirst().String()%> = "<%==s ID%><%==s m.GetID().ToUpperFirst().String()%>";<% b.WriteString("\n")%><%
        }
    }
} %>
<% return b.String() %>