<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityFlag_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity {

    static readonly <%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%><% b.WriteRune(' ') %><%==s c.GetID().WithComponentSuffix().ToLowerFirst().String()%> = new <%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%>();

    public bool <%==s c.GetPrefixOrDefault().ToLowerFirst().String() %><%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> {
        get { return HasComponent(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>); }
        set {
            if (value != <%==s c.GetPrefixOrDefault().ToLowerFirst().String() %><%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>) {
                var index = <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>;
                if (value) {
                    var componentPool = GetComponentPool(index);
                    var component = componentPool.Count > 0
                            ? componentPool.Pop()
                            : <%==s c.GetID().WithComponentSuffix().ToLowerFirst().String()%>;

                    AddComponent(index, component);
                } else {
                    RemoveComponent(index);
                }
            }
        }
    }
}
<% return b.String() %>