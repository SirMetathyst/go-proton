<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextMatcher_C_1_4_2(c *model.C, b *bytes.Buffer) string %>
public sealed partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher 
{
    public static Entitas.IAllOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params int[] indices) 
    {
          return Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(indices);
    }

    public static Entitas.IAllOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] matchers)
    {
          return Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(matchers);
    }

    public static Entitas.IAnyOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params int[] indices)
    {
          return Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(indices);
    }

    public static Entitas.IAnyOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] matchers)
    {
          return Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(matchers);
    }
}
<% return b.String() %>