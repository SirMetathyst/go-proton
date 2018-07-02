<% for _, m := range c.GetMember() { %>    public <%==s m.GetValue().String()%><% b.WriteRune(' ');%><%==s m.GetID().ToLowerFirst().String()%>;
<% } %>}