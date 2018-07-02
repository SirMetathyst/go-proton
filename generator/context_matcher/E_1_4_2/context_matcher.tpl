<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextMatcher_E_1_4_2(ctx *model.Context, b *bytes.Buffer) string %>public sealed partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher {

    public static Entitas.IAllOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params int[] Index) => Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(Index);
    public static Entitas.IAllOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] Matcher) => Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(Matcher);
    public static Entitas.IAnyOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params int[] Index) => Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(Index);
    public static Entitas.IAnyOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] Matcher) => Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(Matcher);

}
<% return b.String() %>