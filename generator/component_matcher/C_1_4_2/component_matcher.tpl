<%: func ComponentMatcher_C_1_4_2(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string %>
public sealed partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher
{
    static Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> _matcher<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>;

    public static Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity><% b.WriteRune(' ')%><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> 
    {
        get 
        {
            if (_matcher<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> == null) 
            {
                var matcher = (Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>)Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>);
                matcher.componentNames = <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentNames;
                _matcher<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> = matcher;
            }

            return _matcher<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        }
    }
}
<% return b.String() %>