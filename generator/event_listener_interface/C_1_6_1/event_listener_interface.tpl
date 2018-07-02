<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerInterface_C_1_6_1(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public interface I<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String() + c.GetID().WithoutComponentSuffix().ToUpperFirst().String() + c.GetEventTypeSuffix().String()%>Listener {
    void On<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() + c.GetEventTypeSuffix().String()%>(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity entity<% EventListenerArgumentPass_1_6_1(c, b)%>);
}
<% return b.String() %>