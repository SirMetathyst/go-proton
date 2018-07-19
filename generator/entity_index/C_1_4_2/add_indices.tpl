<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexAddIndices_C_1_4_2(cp []*model.CP, b *bytes.Buffer) string %>
<% for _, ccp := range cp { 
    for _, c := range ccp.ContextList() { 
    for _, m := range ccp.MemberList() {
        if m.EntityIndex() > 0 { 

    ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
    Type := m.Value().String()
    IndexType := ""
    if m.EntityIndex() == 1 {
        IndexType = "Entitas.PrimaryEntityIndex"
    } else if m.EntityIndex() > 1 {
        IndexType = "Entitas.EntityIndex"
    }
%><% b.WriteString("\t\t")%><%==s c.ID().WithoutContextSuffix().ToLowerFirst().String()%>.AddEntityIndex(new <%==s IndexType %><<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>(
<% if len(ccp.MembersWithEntityIndex()) == 1 { %><% b.WriteString("\t\t\t")%><%==s ID%><% } else if len(ccp.MembersWithEntityIndex()) > 1 { %><% b.WriteString("\t\t\t")%><%==s ID%><%==s m.ID().ToUpperFirst().String()%><% 
} %>,<%==s c.ID().WithoutContextSuffix().ToLowerFirst().String()%>.GetGroup(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s ID %>), (e, c) => ((<%==s ccp.ID().WithComponentSuffix().ToUpperFirst().String()%>)c).<%==s m.ID().ToLowerFirst().String()%>));<% b.WriteString("\n\n")%><% }}}} %>
<% return b.String() %>