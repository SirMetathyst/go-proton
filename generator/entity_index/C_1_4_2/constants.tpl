<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexConstants_C_1_4_2(cp []*model.CP, b *bytes.Buffer) string %>
<% for _, ccp := range cp { 
    ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
    if len(ccp.MembersWithEntityIndex()) == 1 {
%><% b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%> = "<%==s ID%>";<% b.WriteString("\n")%><% } else if len(ccp.MembersWithEntityIndex()) > 1 {
for _, m := range ccp.MemberList() { 
%><% b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%><%==s m.ID().ToUpperFirst().String()%> = "<%==s ID%><%==s m.ID().ToUpperFirst().String()%>";<% b.WriteString("\n")%><%
        }
    }
} %>
<% return b.String() %>