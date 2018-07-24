<%: func EventSystems_C_1_6_1(c *entitas.C, cp []*entitas.CP, b *bytes.Buffer) string %>
public sealed class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>EventSystems : Feature
{
    public <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>EventSystems(Contexts contexts) 
    {
<%
    sort.Sort(byPriority(cp))
    for _, ccp := range cp {
        b.WriteString("\t\tAdd(new ")
        b.WriteString(componentID(c, ccp).ToUpperFirst().String())
        b.WriteString("EventSystem(contexts));")
        b.WriteString(" // priority: ")
        b.WriteString(strconv.Itoa(ccp.EventPriority()))
        b.WriteRune('\n')
    }
%>
    }
}
<% return b.String() %>