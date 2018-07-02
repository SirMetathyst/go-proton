<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndexGetIndices_C_1_4_2(c []*model.Component, b *bytes.Buffer) string %>
   <% for _, cc := range c { 
        for _, cctx := range cc.GetContext() { 
        for _, m := range cc.GetMember() {
            if m.GetEntityIndex() > 0 { 
                
     
             ID := cc.GetID().WithoutComponentSuffix().ToUpperFirst().String()
             Type := m.GetValue().String()
             MemberName := m.GetID().ToLowerFirst().String()
             IndexName := ""
                IndexType := ""
            if m.GetEntityIndex() == 1 {
                IndexType = "Entitas.PrimaryEntityIndex"
            } else if m.GetEntityIndex() > 1 {
                IndexType = "Entitas.EntityIndex"
            }
            
              if cc.GetEntityIndexCount() == 1 {
                  IndexName = ID
            } else if cc.GetEntityIndexCount() > 1 {
                IndexName = ID + m.GetID().ToUpperFirst().String() 
        }
if m.GetEntityIndex() == 1 { b.WriteRune('\t')%>public static <%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity GetEntityWith<%==s IndexName %>(this <%==s cctx.GetID().WithContextSuffix().ToUpperFirst().String()%> context, <%==s Type%><% b.WriteRune(' ')%><%==s MemberName %>) {
        return ((<%==s IndexType %><<%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>)context.GetEntityIndex(Contexts.<%==s IndexName%>)).GetEntity(<%==s MemberName %>); 
    }<% b.WriteRune('\n')%><%
} else if m.GetEntityIndex() > 1 { b.WriteRune('\t')%>public static System.Collections.Generic.HashSet<<%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> GetEntitiesWith<%==s IndexName %>(this <%==s cctx.GetID().WithContextSuffix().ToUpperFirst().String()%> context, <%==s Type%><% b.WriteRune(' ')%><%==s MemberName %>) {
        return ((<%==s IndexType %><<%==s cctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>)context.GetEntityIndex(Contexts.<%==s IndexName%>)).GetEntities(<%==s MemberName %>); 
    }<% }}}}} %>
<% return b.String() %>