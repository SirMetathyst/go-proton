package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Describe Alias Builder ...
var _ = Describe("Alias Builder", func() {

	// Creating an invalid AliasBuilder ...
	It("will panic when a nil alias list is used in creating an alias builder", func() {
		Expect(func() { proton.NewAliasBuilder(nil) }).To(Panic())
	})

	// Creating a valid AliasBuilder ...
	It("will return a pointer to an alias builder", func() {
		aliasList := proton.NewAliasList()
		Expect(func() { proton.NewAliasBuilder(aliasList) }).ToNot(Panic())
		Expect(proton.NewAliasBuilder(aliasList)).ToNot(BeNil())
	})

	// Creating a valid alias using AliasBuilder ...
	DescribeTable("creating a valid alias using alias builder",
		func(aliasData AliasData) {
			aliasList := proton.NewAliasList()
			aliasBuilder := proton.NewAliasBuilder(aliasList)
			aliasBuilder.SetID(aliasData.id)
			aliasBuilder.SetValue(aliasData.value)
			Expect(aliasBuilder.Build()).To(BeNil())
			Expect(aliasList.AliasWithID(aliasData.id)).ToNot(BeNil())

			// Building twice with the same AliasBuilder returns an error
			err := aliasBuilder.Build()
			Expect(err).To(Equal(proton.ErrAliasBuilderAliasAlreadyBuilt))

		}, ValidAliasEntrySlice...,
	)

	// AliasBuilder does not return an invalid alias ...
	DescribeTable("alias builder does not return invalid alias",
		func(aliasData AliasData) {
			aliasList := proton.NewAliasList()
			aliasBuilder := proton.NewAliasBuilder(aliasList)
			aliasBuilder.SetID(aliasData.id)
			aliasBuilder.SetValue(aliasData.value)
			Expect(aliasBuilder.Build()).ToNot(BeNil())
			Expect(aliasList.AliasWithID(aliasData.id)).To(BeNil())
		}, InvalidAliasEntrySlice...,
	)

	// AliasBuilder cannot create a duplicate alias ...
	DescribeTable("alias builder cannot create a duplicate alias",
		func(aliasData AliasData) {
			aliasList := proton.NewAliasList()

			buildAlias := func(value string) error {
				aliasBuilder := proton.NewAliasBuilder(aliasList)
				aliasBuilder.SetID(aliasData.id)
				aliasBuilder.SetValue(value)
				return aliasBuilder.Build()
			}

			// When no duplicate has been added ...
			Expect(buildAlias(aliasData.value)).To(BeNil())
			Expect(aliasList.AliasWithID(aliasData.id)).ToNot(BeNil())

			// When duplicate has been added ...
			Expect(buildAlias("value")).ToNot(BeNil())
			Expect(aliasList.AliasWithID(aliasData.id).Value()).ToNot(Equal("value"))

		}, ValidAliasEntrySlice...,
	)
})
