<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextMatcher_C_1_4_2(ctx *model.Context, b *bytes.Buffer) string %>public sealed partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher {

    public static Entitas.IAllOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params int[] indices) {
        return Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(indices);
    }

    public static Entitas.IAllOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] matchers) {
          return Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(matchers);
    }

    public static Entitas.IAnyOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params int[] indices) {
          return Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(indices);
    }

    public static Entitas.IAnyOfMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params Entitas.IMatcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] matchers) {
          return Entitas.Matcher<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(matchers);
    }
}
<% return b.String() %>