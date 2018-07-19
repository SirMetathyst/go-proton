<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerComponentEntity_C_1_6_1(c *model.C, cp *model.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity
{
    public void Add<%==s ComponentID(c, cp).String()%>(I<%==s ComponentID(c, cp).String()%> value) 
    {
        var listeners = has<%==s ComponentID(c, cp).String()%>
            ? <%==s ComponentID(c, cp).ToLowerFirst().String()%>.value
            : new System.Collections.Generic.List<I<%==s ComponentID(c, cp).String()%>>();
        listeners.Add(value);
        Replace<%==s ComponentID(c, cp).String()%>(listeners);
    }
    
    public void Remove<%==s ComponentID(c, cp).String()%>(I<%==s ComponentID(c, cp).String()%> value, bool removeComponentWhenEmpty = true)
    {
        var listeners = <%==s ComponentID(c, cp).ToLowerFirst().String()%>.value;
        listeners.Remove(value);
        if (removeComponentWhenEmpty && listeners.Count == 0) 
        {
            Remove<%==s ComponentID(c, cp).String()%>();
        } 
        else 
        {
            Replace<%==s ComponentID(c, cp).String()%>(listeners);
        }
    }
}
<% return b.String() %>