<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentContext_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public partial class <%==s ctx.GetID().WithContextSuffix().ToUpperFirst().String()%> {

    public <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity <%==s c.GetID().ToLowerFirst().String()%>Entity { get { return GetGroup(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s c.GetID().ToUpperFirst().String()%>).GetSingleEntity(); } }
    public <%==s c.GetID().ToUpperFirst().String()%>Component <%==s c.GetID().ToLowerFirst().String()%> { get { return <%==s c.GetID().ToLowerFirst().String()%>Entity.<%==s c.GetID().ToLowerFirst().String()%>; } }
    public bool <%==s c.GetPrefixOrDefault().ToLowerFirst().String()  %><%==s c.GetID().ToUpperFirst().String()%> { get { return <%==s c.GetID().ToLowerFirst().String()%>Entity != null; } }

    public <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity Set<%==s c.GetID().String()%>(<% ComponentContextArgument_1_4_2(c,b) %>) {
        if (<%==s c.GetPrefixOrDefault().ToLowerFirst().String()  %><%==s c.GetID().String()%>) {
            throw new Entitas.EntitasException("Could not set <%==s c.GetID().ToUpperFirst().String()%>!\n" + this + " already has an entity with <%==s c.GetID().ToUpperFirst().WithComponentSuffix().String()%>!",
                "You should check if the context already has a <%==s c.GetID().ToLowerFirst().String()%>Entity before setting it or use context.Replace<%==s c.GetID().ToUpperFirst().String()%>().");
        }
        var entity = CreateEntity();
        entity.Add<%==s c.GetID().String()%>(<% ComponentContextArgumentPass_1_4_2(c, b) %>);
        return entity;
    }

    public void Replace<%==s c.GetID().ToUpperFirst().String()%>(<% ComponentContextArgument_1_4_2(c, b) %>) {
        var entity = <%==s c.GetID().ToLowerFirst().String()%>Entity;
        if (entity == null) {
            entity = Set<%==s c.GetID().ToUpperFirst().String()%>(<% ComponentContextArgumentPass_1_4_2(c, b) %>);
        } else {
            entity.Replace<%==s c.GetID().ToUpperFirst().String()%>(<% ComponentContextArgumentPass_1_4_2(c, b) %>);
        }
    }

    public void Remove<%==s c.GetID().ToUpperFirst().String()%>() {
        <%==s c.GetID().ToLowerFirst().String()%>Entity.Destroy();
    }
}
<% return b.String() %>
