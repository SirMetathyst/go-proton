<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexAddIndices_C_1_4_2(c []*model.Component, b *bytes.Buffer) string %>
<% for _, cc := range c { 
    for _, cctx := range cc.GetContext() { 
    for _, m := range cc.GetMember() {
        if m.GetEntityIndex() > 0 { 

    ID := cc.GetID().WithoutComponentSuffix().ToUpperFirst().String()
    Type := m.GetValue().String()
    IndexType := ""
    if m.GetEntityIndex() == 1 {
        IndexType = "Entitas.PrimaryEntityIndex"
    } else if m.GetEntityIndex() > 1 {
        IndexType = "Entitas.EntityIndex"
    }
%><% b.WriteString("\t\t")%><%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%>.AddEntityIndex(new <%==s IndexType %><<%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>(
<% if cc.GetEntityIndexCount() == 1 { %><% b.WriteString("\t\t\t")%><%==s ID%><% } else if cc.GetEntityIndexCount() > 1 { %><% b.WriteString("\t\t\t")%><%==s ID%><%==s m.GetID().ToUpperFirst().String()%><% 
} %>,<%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%>.GetGroup(<%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s ID %>), (e, c) => ((<%==s cc.GetID().WithComponentSuffix().ToUpperFirst().String()%>)c).<%==s m.GetID().ToLowerFirst().String()%>));<% b.WriteString("\n")%><% }}}} %>
<% return b.String() %>