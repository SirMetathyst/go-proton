<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityInterface_C_1_4_2(c *model.Component, b *bytes.Buffer) string %>public interface I<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %>Entity {

    <%==s c.GetID().WithComponentSuffix().ToUpperFirst().String() %><% b.WriteRune(' ')%><%==s c.GetID().WithoutComponentSuffix().ToLowerFirst().String() %> { get; }
    bool has<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %> { get; }

    void Add<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %>(<% ComponentEntityInterfaceArgument_C_1_4_2(c, b) %>);
    void Replace<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %>(<% ComponentEntityInterfaceArgument_C_1_4_2(c, b) %>);
    void Remove<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %>();
}
<% return b.String() %>