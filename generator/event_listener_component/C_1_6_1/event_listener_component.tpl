<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerComponent_C_1_6_1(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>
[Entitas.CodeGeneration.Attributes.DontGenerate(false)]
public sealed class <%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>ListenerComponent : Entitas.IComponent {
    public System.Collections.Generic.List<I<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>Listener> value;
}
<% return b.String() %>