<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Contexts_C_1_4_2(ctx []*model.Context, b *bytes.Buffer) string %>public partial class Contexts : Entitas.IContexts {

    public static Contexts sharedInstance {
        get {
			if (_sharedInstance == null) {
				_sharedInstance = new Contexts();
            }
            return _sharedInstance;
        }
        set { _sharedInstance = value; }
    }
  
    static Contexts _sharedInstance;

<% for _, cctx := range ctx { %><% b.WriteRune('\t')%>public <%==s cctx.GetID().WithContextSuffix().ToUpperFirst().String()%><% b.WriteRune(' ')%><%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%> { get; set; }<% b.WriteRune('\n')%><% } %>

    public Entitas.IContext[] allContexts { get { return new Entitas.IContext [] { <% for Index, cctx := range ctx { %><%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%><% if Index != len(ctx)-1 { b.WriteString(", ") }}%> }; } }
  
	public Contexts() {

<% for _, cctx := range ctx { %>
<% b.WriteString("\t\t")%><%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%> = new <%==s cctx.GetID().WithContextSuffix().ToUpperFirst().String()%>();<% b.WriteRune('\n')%>
<% } %>

        var postConstructors = System.Linq.Enumerable.Where(
            GetType().GetMethods(),
            method => System.Attribute.IsDefined(method, typeof(Entitas.CodeGeneration.Attributes.PostConstructorAttribute))
        );
  
        foreach (var postConstructor in postConstructors) {
            postConstructor.Invoke(this, null);
        }
    }
  
    public void Reset() {
        var contexts = allContexts;
        for (int i = 0; i < contexts.Length; i++) {
            contexts[i].Reset();
        }
    }
}
<% return b.String() %>