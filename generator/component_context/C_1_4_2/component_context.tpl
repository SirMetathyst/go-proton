<%: func ComponentContext_C_1_4_2(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%> 
{
    public <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity 
    { 
        get 
        { 
            return GetGroup(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>).GetSingleEntity(); 
        } 
    }

    public <%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%><% b.WriteRune(' ')%><%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%> 
    { 
        get 
        { 
            return <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity.<%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>; 
        } 
    }

    public bool has<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> 
    { 
        get 
        { 
            return <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity != null; 
        } 
    }

    public <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity Set<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentContextArgument_1_4_2(cp,b) %>) 
    {
        if (has<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>) 
        {
            throw new Entitas.EntitasException("Could not set <%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>!\n" + this + " already has an entity with <%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().WithComponentSuffix().String()%>!",
                "You should check if the context already has a <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity before setting it or use context.Replace<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>().");
        }
        var entity = CreateEntity();
        entity.Add<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentContextArgumentPass_1_4_2(cp, b) %>);
        return entity;
    }

    public void Replace<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentContextArgument_1_4_2(cp, b) %>) 
    {
        var entity = <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity;
        if (entity == null) 
        {
            entity = Set<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentContextArgumentPass_1_4_2(cp, b) %>);
        } 
        else 
        {
            entity.Replace<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentContextArgumentPass_1_4_2(cp, b) %>);
        }
    }

    public void Remove<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>() 
    {
        <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity.Destroy();
    }
}
<% return b.String() %>
