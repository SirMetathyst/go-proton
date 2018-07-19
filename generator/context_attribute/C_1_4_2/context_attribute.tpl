<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextAttribute_C_1_4_2(c *model.C, b *bytes.Buffer) string %>
public sealed class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Attribute : Entitas.CodeGeneration.Attributes.ContextAttribute 
{
    public <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Attribute() : base("<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>") 
    {
    }
}
<% return b.String() %>