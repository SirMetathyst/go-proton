<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func ContextObserver_C_1_4_2(ctx []*model.Context, b *bytes.Buffer) string %>public partial class Contexts {
#if (!ENTITAS_DISABLE_VISUAL_DEBUGGING && UNITY_EDITOR)
    [Entitas.CodeGeneration.Attributes.PostConstructor]
    public void InitializeContexObservers() {
        try {
<% for _, cctx := range ctx { b.WriteString("\t\t\t")%>CreateContextObserver(<%==s cctx.GetID().WithoutContextSuffix().ToLowerFirst().String()%>);<% b.WriteRune('\n') } %>
        } catch(System.Exception) {
        }
    }
    public void CreateContextObserver(Entitas.IContext context) {
        if (UnityEngine.Application.isPlaying) {
            var observer = new Entitas.VisualDebugging.Unity.ContextObserver(context);
            UnityEngine.Object.DontDestroyOnLoad(observer.gameObject);
        }
    }
#endif
}
<% return b.String() %>