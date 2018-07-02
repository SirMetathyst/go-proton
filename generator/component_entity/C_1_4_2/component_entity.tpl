<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntity_C_1_4_2(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>public partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity {

    public <%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%><% b.WriteRune(' ')%><%==s c.GetID().WithoutComponentSuffix().ToLowerFirst().String()%> { get { return (<%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%>)GetComponent(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>); } }
    public bool <%==s c.GetPrefixOrDefault().ToLowerFirst().String() %><%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> { get { return HasComponent(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>); } }

    public void Add<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentEntityArgument_C_1_4_2(c, b) %>) {
        var index = <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        var component = CreateComponent<<%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%>>(index);
<% ComponentEntityAssignment_C_1_4_2(c, b) %>
        AddComponent(index, component);
    }

    public void Replace<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>(<% ComponentEntityArgument_C_1_4_2(c, b) %>) {
        var index = <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>;
        var component = CreateComponent<<%==s c.GetID().WithComponentSuffix().ToUpperFirst().String()%>>(index);
<% ComponentEntityAssignment_C_1_4_2(c, b) %>
        ReplaceComponent(index, component);
    }

    public void Remove<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>() {
        RemoveComponent(<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>);
    }
}
<% return b.String() %>