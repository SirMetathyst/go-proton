<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerInterface_C_1_6_1(c *model.C, cp *model.CP, b *bytes.Buffer) string %>
public interface I<%==s ComponentID(c, cp).ToUpperFirst().String()%> 
{
    void On<%==s MethodID(cp).ToUpperFirst().String()%>(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity entity<% EventListenerArgumentPass_1_6_1(cp, b)%>);
}
<% return b.String() %>