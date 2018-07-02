<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EntityIndex_C_1_4_2(c []*model.Component, ei []*model.EntityIndex, b *bytes.Buffer) string %>
public partial class Contexts {

<% EntityIndexConstants_C_1_4_2(c, b) %>
<% EntityIndexCustomConstants_C_1_4_2(ei, b) %>

    [Entitas.CodeGeneration.Attributes.PostConstructor]
    public void InitializeEntityIndices() {
<% EntityIndexAddIndices_C_1_4_2(c, b) %>
<% EntityIndexAddCustomIndices_C_1_4_2(ei, b) %>
    }
}

public static class ContextsExtensions {
<% EntityIndexGetIndices_C_1_4_2(c, b) %>
<% EntityIndexGetCustomIndices_C_1_4_2(ei, b) %>
}
<% return b.String() %>