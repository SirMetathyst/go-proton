
<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Entity_C_1_4_2(ctx *model.Context, b *bytes.Buffer) string %>public sealed partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity : Entitas.Entity {
}
<% return b.String() %>