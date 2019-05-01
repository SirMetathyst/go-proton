package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// ComponentMemberBuilderData ...
type ComponentMemberBuilderData struct {
	id              string
	entityIndexType proton.EntityIndexType
	expectedErr     error
}

var (

	// InvalidComponentMemberBuilderEntrySlice ...
	InvalidComponentMemberBuilderEntrySlice = []TableEntry{
		Entry("", ComponentMemberBuilderData{"", proton.EntityIndexTypeNone, proton.ErrComponentMemberIDUndefined}),            // 0
		Entry("", ComponentMemberBuilderData{" ", proton.EntityIndexTypeNone, proton.ErrComponentMemberIDContainsWhitespace}),  // 1
		Entry("", ComponentMemberBuilderData{"id", proton.EntityIndexTypeNone, proton.ErrComponentMemberBuilderAliasNotFound}), // 2
		Entry("", ComponentMemberBuilderData{"id", 3, proton.ErrComponentMemberBuilderAliasNotFound}),                          // 3
		Entry("", ComponentMemberBuilderData{"id", -1, proton.ErrComponentMemberBuilderAliasNotFound}),                         // 4
	}
)

// Describe Component Member Builder ...
var _ = Describe("Component Member Builder", func() {

	// Creating an invalid ComponentMemberBuilder ...
	It("will panic when a nil component member list and alias list is used in creating a component member builder", func() {
		Expect(func() { proton.NewComponentMemberBuilder(nil, nil) }).To(Panic())
		Expect(func() { proton.NewComponentMemberBuilder(proton.NewAliasList(), nil) }).To(Panic())
	})

	// Creating a valid ComponentMemberBuilder ...
	It("will return a pointer to a component member builder", func() {
		aliasList := proton.NewAliasList()
		componentMemberList := proton.NewComponentMemberList()
		Expect(func() { proton.NewComponentMemberBuilder(aliasList, componentMemberList) }).ToNot(Panic())
		Expect(proton.NewComponentMemberBuilder(aliasList, componentMemberList)).ToNot(BeNil())
	})

	// Creating a valid component member using ComponentMemberBuilder ...
	DescribeTable("creating a valid component member using component member builder",
		func(componentMemberData ComponentMemberData) {
			aliasList := proton.NewAliasList()
			componentMemberList := proton.NewComponentMemberList()
			componentMemberBuilder := proton.NewComponentMemberBuilder(aliasList, componentMemberList)
			componentMemberBuilder.SetID(componentMemberData.id)
			componentMemberBuilder.SetValue(componentMemberData.value)
			componentMemberBuilder.SetEntityIndexType(componentMemberData.entityIndexType)

			Expect(componentMemberBuilder.Build()).To(BeNil())
			Expect(componentMemberList.ComponentMemberWithID(componentMemberData.id)).ToNot(BeNil())

			// Building twice with the same ContextBuilder returns an error
			err := componentMemberBuilder.Build()
			Expect(err).To(Equal(proton.ErrComponentMemberBuilderMemberAlreadyBuilt))

		}, ValidComponentMemberEntrySlice...,
	)

	// Creating a valid component member using ComponentMemberBuilder ...
	DescribeTable("creating a valid component member using component member builder with alias",
		func(componentMemberAliasData ComponentMemberAliasData) {
			aliasList := proton.NewAliasList()
			aliasList.AddAlias(componentMemberAliasData.alias)

			componentMemberList := proton.NewComponentMemberList()
			componentMemberBuilder := proton.NewComponentMemberBuilder(aliasList, componentMemberList)
			componentMemberBuilder.SetID(componentMemberAliasData.id)
			componentMemberBuilder.SetAlias(componentMemberAliasData.alias.ID().String())
			componentMemberBuilder.SetEntityIndexType(componentMemberAliasData.entityIndexType)

			Expect(componentMemberBuilder.Build()).To(BeNil())

			componentMember := componentMemberList.ComponentMemberWithID(componentMemberAliasData.id)

			Expect(componentMember).ToNot(BeNil())
			Expect(componentMember.Value().String()).To(Equal(componentMemberAliasData.alias.Value().String()))

			// Building twice with the same ContextBuilder returns an error
			err := componentMemberBuilder.Build()
			Expect(err).To(Equal(proton.ErrComponentMemberBuilderMemberAlreadyBuilt))

		}, ValidComponentMemberAliasEntrySlice...,
	)

	// Creating a valid component member using ComponentMemberBuilder with invalid data ...
	DescribeTable("creating a valid component member using component member builder with invalid data",
		func(componentMemberData ComponentMemberData) {
			aliasList := proton.NewAliasList()

			componentMemberList := proton.NewComponentMemberList()
			componentMemberBuilder := proton.NewComponentMemberBuilder(aliasList, componentMemberList)
			componentMemberBuilder.SetID(componentMemberData.id)
			componentMemberBuilder.SetValue(componentMemberData.value)
			componentMemberBuilder.SetEntityIndexType(componentMemberData.entityIndexType)

			Expect(componentMemberBuilder.Build()).To(Equal(componentMemberData.expectedErr))
			Expect(componentMemberList.ComponentMemberWithID(componentMemberData.id)).To(BeNil())

		}, InvalidComponentMemberEntrySlice...,
	)

	// Creating a valid component member using ComponentMemberBuilder should not be possible with invalid alias id ...
	DescribeTable("creating a valid component member using component member builder should not be possible with invalid alias id",
		func(componentMemberBuilderData ComponentMemberBuilderData) {
			aliasList := proton.NewAliasList()
			componentMemberList := proton.NewComponentMemberList()
			componentMemberBuilder := proton.NewComponentMemberBuilder(aliasList, componentMemberList)
			componentMemberBuilder.SetID(componentMemberBuilderData.id)
			componentMemberBuilder.SetAlias("does_not_exist")
			componentMemberBuilder.SetEntityIndexType(componentMemberBuilderData.entityIndexType)

			Expect(componentMemberBuilder.Build()).To(Equal(componentMemberBuilderData.expectedErr))
			Expect(componentMemberList.ComponentMemberWithID(componentMemberBuilderData.id)).To(BeNil())

		}, InvalidComponentMemberBuilderEntrySlice...,
	)

	// ComponentMemberBuilder cannot create a duplicate component member ...
	DescribeTable("component member builder cannot create a duplicate component member",
		func(componentMemberData ComponentMemberData) {
			aliasList := proton.NewAliasList()
			componentMemberList := proton.NewComponentMemberList()

			buildComponentMember := func() error {
				componentMemberBuilder := proton.NewComponentMemberBuilder(aliasList, componentMemberList)
				componentMemberBuilder.SetID(componentMemberData.id)
				componentMemberBuilder.SetValue(componentMemberData.value)
				componentMemberBuilder.SetEntityIndexType(componentMemberData.entityIndexType)
				return componentMemberBuilder.Build()
			}

			// When no duplicate has been added ...
			Expect(buildComponentMember()).To(BeNil())
			Expect(componentMemberList.ComponentMemberWithID(componentMemberData.id)).ToNot(BeNil())

			// When trying to build a duplicate ...
			Expect(buildComponentMember()).ToNot(BeNil())

		}, ValidComponentMemberEntrySlice...,
	)

})
