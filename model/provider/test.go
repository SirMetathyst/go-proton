package modelprovider

import (
	"github.com/SirMetathyst/proton/configuration"
	"github.com/SirMetathyst/proton/model"
)

// createTestModelFluent ...
func createTestModelFluent(m *model.M) {
	// Fluent API
	ContextA := m.CreateContext(false, "ContextA").SetDefault(true)
	ContextB := m.CreateContext(false, "ContextB")
	ContextC := m.CreateContext(false, "ContextC")
	ContextD := m.CreateContext(false, "ContextD")

	m.CreateComponent(false, "ComponentA").AddContext(false, ContextA).CreateMember(false, "ValueA").SetValue("string")
	m.CreateComponent(false, "ComponentB").AddContext(false, ContextB).CreateMember(false, "ValueB").SetValue("string")
	m.CreateComponent(false, "ComponentC").AddContext(false, ContextC).CreateMember(false, "ValueC").SetValue("string")
	m.CreateComponent(false, "ComponentD").AddContext(false, ContextD).CreateMember(false, "ValueD").SetValue("string")
}

// createTestModel ...
func createTestModel(m *model.M) {
	// Non-Fluent API
	m.AddContext(false, model.NewContextWithID("ContextA"))
	m.AddContext(false, model.NewContextWithID("ContextB"))
	m.AddContext(false, model.NewContextWithID("ContextC"))
	m.AddContext(false, model.NewContextWithID("ContextD"))
}

// Test ...
func Test(c *configuration.C) (*model.M, error) {
	m := model.NewModel()

	createTestModelFluent(m)
	//createTestModel(m)

	return m, nil
}
