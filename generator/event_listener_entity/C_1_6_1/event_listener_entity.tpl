<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerComponent_C_1_6_1(ctx *model.Context, c *model.Component, b *bytes.Buffer) string %>
public partial class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>Entity {
    public void Add<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%>Listener(I<%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()+c.GetID().WithoutComponentSuffix().ToUpperFirst().String()%> value) {
        var listeners = has${EventListener}
            ? ${eventListener}.value
            : new System.Collections.Generic.List<I${EventListener}>();
        listeners.Add(value);
        Replace${EventListener}(listeners);
    }
    public void Remove${EventListener}(I${EventListener} value, bool removeComponentWhenEmpty = true) {
        var listeners = ${eventListener}.value;
        listeners.Remove(value);
        if (removeComponentWhenEmpty && listeners.Count == 0) {
            Remove${EventListener}();
        } else {
            Replace${EventListener}(listeners);
        }
    }
}
<% return b.String() %>