<%: func Component_E_1_4_2(cp *entitas.CP, b *bytes.Buffer) string %>
public sealed partial class <%==s cp.ID().WithComponentSuffix().String()%> : Entitas.IComponent 
{
<%+ "member.tpl" %>
}
<% return b.String() %>