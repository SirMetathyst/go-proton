package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Describe Context Builder ...
var _ = Describe("Context Builder", func() {

	// Creating an invalid ContextBuilder ...
	It("will panic when a nil context list is used in creating a context builder", func() {
		Expect(func() { proton.NewContextBuilder(nil) }).To(Panic())
	})

	// Creating a valid ContextBuilder ...
	It("will return a pointer to a context builder", func() {
		contextList := proton.NewContextList()
		Expect(func() { proton.NewContextBuilder(contextList) }).ToNot(Panic())
		Expect(proton.NewContextBuilder(contextList)).ToNot(BeNil())
	})

	// Creating a valid context using ContextBuilder ...
	DescribeTable("creating a valid context using context builder",
		func(contextData ContextData) {
			contextList := proton.NewContextList()
			contextBuilder := proton.NewContextBuilder(contextList)
			contextBuilder.SetID(contextData.id)
			Expect(contextBuilder.Build()).To(BeNil())
			Expect(contextList.ContextWithID(contextData.id)).ToNot(BeNil())

			// Building twice with the same ContextBuilder returns an error
			err := contextBuilder.Build()
			Expect(err).To(Equal(proton.ErrContextBuilderContextAlreadyBuilt))

		}, ValidContextEntrySlice...,
	)

	// ContextBuilder does not return an invalid context ...
	DescribeTable("context builder does not return invalid context",
		func(contextData ContextData) {
			contextList := proton.NewContextList()
			contextBuilder := proton.NewContextBuilder(contextList)
			contextBuilder.SetID(contextData.id)
			Expect(contextBuilder.Build()).ToNot(BeNil())
			Expect(contextList.ContextWithID(contextData.id)).To(BeNil())
		}, InvalidContextEntrySlice...,
	)

	// ContextBuilder cannot create a duplicate context ...
	DescribeTable("context builder cannot create a duplicate context",
		func(contextData ContextData) {
			contextList := proton.NewContextList()

			buildContext := func() error {
				contextBuilder := proton.NewContextBuilder(contextList)
				contextBuilder.SetID(contextData.id)
				return contextBuilder.Build()
			}

			// When no duplicate has been added ...
			Expect(buildContext()).To(BeNil())
			Expect(contextList.ContextWithID(contextData.id)).ToNot(BeNil())

			// When trying to build a duplicate ...
			Expect(buildContext()).ToNot(BeNil())

		}, ValidContextEntrySlice...,
	)
})
