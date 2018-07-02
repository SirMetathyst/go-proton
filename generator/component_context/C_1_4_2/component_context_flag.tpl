<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentContextFlag_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public partial class <%==s ctx.GetID().WithContextSuffix().ToUpperFirst().String()%> {

    public <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity <%==s c.GetID().ToLowerFirst().String()%>Entity { get { return GetGroup(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s c.GetID().ToUpperFirst().String()%>).GetSingleEntity(); } }

    public bool <%==s c.GetPrefixOrDefault().ToLowerFirst().String() %><%==s c.GetID().ToUpperFirst().String()%> {
        get { return <%==s c.GetID().ToLowerFirst().String()%>Entity != null; }
        set {
            var entity = <%==s c.GetID().ToLowerFirst().String()%>Entity;
            if (value != (entity != null)) {
                if (value) {
                    CreateEntity().<%==s c.GetPrefixOrDefault().ToLowerFirst().String()  %><%==s c.GetID().ToUpperFirst().String()%> = true;
                } else {
                    entity.Destroy();
                }
            }
        }
    }
}
<% return b.String() %>
