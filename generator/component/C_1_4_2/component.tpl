<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Component_C_1_4_2(c *model.Component, b *bytes.Buffer) string %>public sealed class <%==s c.GetID().ToUpperFirst().WithComponentSuffix().String()%> : Entitas.IComponent {
<%+ "component_member_1_4_2.tpl" %>
<% return b.String() %>