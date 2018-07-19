
<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Entity_C_1_4_2(c *model.C, b *bytes.Buffer) string %>
public sealed partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity : Entitas.Entity
{
}
<% return b.String() %>