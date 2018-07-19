package modelprovider

import (
	"github.com/SirMetathyst/proton/configuration"
	"github.com/SirMetathyst/proton/model"
	"github.com/SirMetathyst/proton/model/builder"
)

// createTestModel ...
func createTestModel() (*model.MD, error) {
	mdb := modelbuilder.NewModelBuilder()

	mdb.NewContext().SetID("ContextA").Build()
	mdb.NewContext().SetID("ContextB").Build()
	mdb.NewContext().SetID("ContextC").Build()
	mdb.NewContext().SetID("ContextD").Build()

	mdb.NewComponent().SetID("ComponentA").AddContext("ContextA").Build()
	mdb.NewComponent().SetID("ComponentB").AddContext("ContextB").Build()
	mdb.NewComponent().SetID("ComponentC").AddContext("ContextC").Build()

	cp := mdb.NewComponent().SetID("ComponentD").AddContext("ContextD")
	cp.NewMember().SetID("Value1").SetValue("int").Build()
	cp.NewMember().SetID("Value2").SetValue("int").Build()
	cp.Build()

	mdb.NewComponent().SetID("ComponentE").AddContext("ContextE").Build()

	return mdb.Build()
}

// Test ...
func Test(c *configuration.C) (*model.MD, error) {
	return createTestModel()
}
