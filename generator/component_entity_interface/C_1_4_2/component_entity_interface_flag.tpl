<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentEntityInterfaceFlag_C_1_4_2(c *model.Component, b *bytes.Buffer) string %>public interface I<%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %>Entity {
    bool <%==s c.GetPrefixOrDefault().ToLowerFirst().String()  %><%==s c.GetID().WithoutComponentSuffix().ToUpperFirst().String() %> { get; set; }
}
<% return b.String() %>