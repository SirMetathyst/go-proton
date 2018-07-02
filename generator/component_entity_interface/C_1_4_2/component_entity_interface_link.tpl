<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityInterfaceLink_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity : I<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>Entity { }
<% return b.String() %>