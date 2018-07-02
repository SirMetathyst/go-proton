<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextAttribute_C_1_4_2(ctx *model.Context, b *bytes.Buffer) string %>public sealed class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Attribute : Entitas.CodeGeneration.Attributes.ContextAttribute {
    
    public <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Attribute() : base("<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>") {
    }
}
<% return b.String() %>