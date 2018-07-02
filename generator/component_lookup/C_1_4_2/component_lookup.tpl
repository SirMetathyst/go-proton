<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ComponentLookup_C_1_4_2(ctx *model.Context, c []*model.Component, b *bytes.Buffer) string %>public static class <%==s ctx.GetID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup {
<% for i, cc := range c {
    b.WriteString("\tpublic const int ")
    b.WriteString(cc.GetID().WithoutComponentSuffix().ToUpperFirst().String())
    b.WriteString(" = ")
    b.WriteString(strconv.Itoa(i))
    b.WriteString(";\n")
} 
%>
	
    public const int TotalComponents = <%==s strconv.Itoa(len(c)) %>;
	
	public static readonly string[] componentNames = {
<% 
for i, cc := range c {
    b.WriteString("\t\t\"")
    b.WriteString(cc.GetID().WithoutComponentSuffix().ToUpperFirst().String())
    b.WriteString("\"")
    if i != len(c)-1 {
        b.WriteString(",\n")
    }
} 
%>
	};
	
	public static readonly System.Type[] componentTypes = {
<% 
for i, cc := range c {
    b.WriteString("\t\ttypeof(")
    b.WriteString(cc.GetID().WithComponentSuffix().ToUpperFirst().String())
    b.WriteRune(')')
    if i != len(c)-1 {
        b.WriteString(",\n")
    }
} 
%>
	};
}
<% return b.String() %>