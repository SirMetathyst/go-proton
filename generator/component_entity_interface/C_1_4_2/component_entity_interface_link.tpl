<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityInterfaceLink_C_1_4_2(c *model.C, cp *model.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity : I<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>Entity 
{ 
}
<% return b.String() %>