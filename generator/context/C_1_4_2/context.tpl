<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Context_C_1_4_2(ctx *model.Context, b *bytes.Buffer) string %>public sealed partial class <%==s ctx.GetID().WithContextSuffix().ToUpperFirst().String()%> : Entitas.Context<<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> {

    public <%==s ctx.GetID().WithContextSuffix().ToUpperFirst().String()%>()
        : base(
            <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.TotalComponents,
            0,
            new Entitas.ContextInfo(
                "<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>",
                <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentNames,
                <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentTypes
            ),
            (entity) =>
#if (ENTITAS_FAST_AND_UNSAFE)
                new Entitas.UnsafeAERC()
#else
                new Entitas.SafeAERC(entity)
#endif
        ) {
    }
}
<% return b.String() %>