<%: func EventListenerInterface_C_1_6_1(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string %>
public interface I<%==s componentID(c, cp).ToUpperFirst().String()%> 
{
    void On<%==s methodID(cp).ToUpperFirst().String()%>(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity entity<% EventListenerArgumentPass_1_6_1(cp, b)%>);
}
<% return b.String() %>