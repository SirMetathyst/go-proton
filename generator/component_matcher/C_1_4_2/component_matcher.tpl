<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentMatcher_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public sealed partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher {

    static Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> _matcher<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>;

    public static Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity><% b.WriteRune(' ')%><%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> {
        get {
            if (_matcher<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> == null) {
                var matcher = (Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>)Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>);
                matcher.componentNames = <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentNames;
                _matcher<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> = matcher;
            }

            return _matcher<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        }
    }
}
<% return b.String() %>