<%: func EntityIndexGetIndices_C_1_4_2(cp []*entitas.CP, b *bytes.Buffer) string %>
   <% for _, ccp := range cp { 
        for _, c := range ccp.ContextList() { 
        for _, m := range ccp.MemberList() {
            if m.EntityIndex() > 0 { 
                
     
             ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
             Type := m.Value().String()
             MemberName := m.ID().ToLowerFirst().String()
             IndexName := ""
                IndexType := ""
            if m.EntityIndex() == 1 {
                IndexType = "Entitas.PrimaryEntityIndex"
            } else if m.EntityIndex() > 1 {
                IndexType = "Entitas.EntityIndex"
            }
            entityIndexCount := len(ccp.MembersWithEntityIndex())
              if entityIndexCount == 1 {
                  IndexName = ID
            } else if entityIndexCount > 1 {
                IndexName = ID + m.ID().ToUpperFirst().String() 
        }
if m.EntityIndex() == 1 { b.WriteRune('\t')%>public static <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity GetEntityWith<%==s IndexName %>(this <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%> context, <%==s Type%><% b.WriteRune(' ')%><%==s MemberName %>)
    {
        return ((<%==s IndexType %><<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>)context.GetEntityIndex(Contexts.<%==s IndexName%>)).GetEntity(<%==s MemberName %>); 
    }<% b.WriteRune('\n')%><%
} else if m.EntityIndex() > 1 { b.WriteRune('\t')%>public static System.Collections.Generic.HashSet<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> GetEntitiesWith<%==s IndexName %>(this <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%> context, <%==s Type%><% b.WriteRune(' ')%><%==s MemberName %>)
    {
        return ((<%==s IndexType %><<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity, <%==s Type%>>)context.GetEntityIndex(Contexts.<%==s IndexName%>)).GetEntities(<%==s MemberName %>); 
    }<% }}}}} %>
<% return b.String() %>