<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Contexts_C_1_4_2(c []*model.C, b *bytes.Buffer) string %>
public partial class Contexts : Entitas.IContexts 
{
    public static Contexts sharedInstance 
    {
        get 
        {
			if (_sharedInstance == null)
            {
				_sharedInstance = new Contexts();
            }
            return _sharedInstance;
        }
        set 
        {
            _sharedInstance = value; 
        }
    }
  
    static Contexts _sharedInstance;

<% for _, cc := range c { %><% b.WriteRune('\t')%>public <%==s cc.ID().WithContextSuffix().ToUpperFirst().String()%><% b.WriteRune(' ')%><%==s cc.ID().WithoutContextSuffix().ToLowerFirst().String()%> { get; set; }<% b.WriteRune('\n')%><% } %>

    public Entitas.IContext[] allContexts 
    {
        get 
        { 
            return new Entitas.IContext [] { <% for Index, cc := range c { %><%==s cc.ID().WithoutContextSuffix().ToLowerFirst().String()%><% if Index != len(c)-1 { b.WriteString(", ") }}%> };
        } 
    }
  
	public Contexts()
    {
<% for _, cc := range c { %>
<% b.WriteString("\t\t")%><%==s cc.ID().WithoutContextSuffix().ToLowerFirst().String()%> = new <%==s cc.ID().WithContextSuffix().ToUpperFirst().String()%>();<% b.WriteRune('\n')%>
<% } %>

        var postConstructors = System.Linq.Enumerable.Where(
            GetType().GetMethods(),
            method => System.Attribute.IsDefined(method, typeof(Entitas.CodeGeneration.Attributes.PostConstructorAttribute))
        );
  
        foreach (var postConstructor in postConstructors)
        {
            postConstructor.Invoke(this, null);
        }
    }
  
    public void Reset() 
    {
        var contexts = allContexts;
        for (int i = 0; i < contexts.Length; i++)
        {
            contexts[i].Reset();
        }
    }
}
<% return b.String() %>