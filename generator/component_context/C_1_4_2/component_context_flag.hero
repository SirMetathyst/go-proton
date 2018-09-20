<%: func ComponentContextFlag_C_1_4_2(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%> 
{
    public <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity 
    { 
        get 
        { 
            return GetGroup(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>).GetSingleEntity(); 
        } 
    }

    public bool <%==s cp.FlagPrefixOrDefault().ToLowerFirst().String() %><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> 
    {
        get 
        { 
            return <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity != null; 
        }
        set {
            var entity = <%==s cp.ID().WithoutComponentSuffix().ToLowerFirst().String()%>Entity;
            if (value != (entity != null))
            {
                if (value) 
                {
                    CreateEntity().<%==s cp.FlagPrefixOrDefault().WithoutComponentSuffix().ToLowerFirst().String()  %><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> = true;
                } 
                else 
                {
                    entity.Destroy();
                }
            }
        }
    }
}
<% return b.String() %>