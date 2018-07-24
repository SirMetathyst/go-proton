<%: func ComponentLookup_C_1_4_2(c *entitas.C, cp []*entitas.CP, b *bytes.Buffer) string %>
public static class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup 
{
<% for i, ccp := range cp {
    b.WriteString("\tpublic const int ")
    b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
    b.WriteString(" = ")
    b.WriteString(strconv.Itoa(i))
    b.WriteString(";\n")
}%>
    public const int TotalComponents = <%==s strconv.Itoa(len(cp)) %>;
	
	public static readonly string[] componentNames = 
    {
<% 
for i, ccp := range cp {
    b.WriteString("\t\t\"")
    b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
    b.WriteString("\"")
    if i != len(cp)-1 {
        b.WriteString(",\n")
    }
} 
%>
	};
	
	public static readonly System.Type[] componentTypes = 
    {
<% 
for i, ccp := range cp {
    b.WriteString("\t\ttypeof(")
    b.WriteString(ccp.ID().WithComponentSuffix().ToUpperFirst().String())
    b.WriteRune(')')
    if i != len(cp)-1 {
        b.WriteString(",\n")
    }
} 
%>
	};
}
<% return b.String() %>