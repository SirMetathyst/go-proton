<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntity_C_1_4_2(c *model.C, cp *model.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity
{
    public <%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%><% b.WriteRune(' ')%><%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%> 
    { 
        get 
        { 
            return (<%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%>)GetComponent(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>); 
        } 
    }

    public bool has<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> 
    { 
        get 
        { 
            return HasComponent(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>); 
        } 
    }

    public void Add<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentEntityArgument_C_1_4_2(cp, b) %>) 
    {
        var index = <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        var component = CreateComponent<<%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%>>(index);
<% ComponentEntityAssignment_C_1_4_2(cp, b) %>
        AddComponent(index, component);
    }

    public void Replace<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentEntityArgument_C_1_4_2(cp, b) %>)
    {
        var index = <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        var component = CreateComponent<<%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%>>(index);
<% ComponentEntityAssignment_C_1_4_2(cp, b) %>
        ReplaceComponent(index, component);
    }

    public void Remove<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>()
    {
        RemoveComponent(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>);
    }
}
<% return b.String() %>