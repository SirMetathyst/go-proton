package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Describe Context List ...
var _ = Describe("Context List", func() {

	// Adding a nil context ...
	It("will panic when a nil context is added", func() {
		contextList := proton.NewContextList()
		Expect(func() { contextList.AddContext(nil) }).To(Panic())
	})

	// Adding a valid context ...
	DescribeTable("adding a valid context",
		func(contextData ContextData) {
			contextList := proton.NewContextList()
			context, _ := proton.NewContext(contextData.id)
			Expect(contextList.AddContext(context)).To(BeNil())
		}, ValidContextEntrySlice...,
	)

	// Adding a duplicate context ...
	DescribeTable("adding a duplicate context",
		func(contextData ContextData) {
			contextList := proton.NewContextList()
			context, _ := proton.NewContext(contextData.id)
			contextList.AddContext(context)
			Expect(contextList.AddContext(context)).To(Equal(proton.ErrContextListTriedToAddDuplicateContextID))
		}, ValidContextEntrySlice...,
	)

	// Returning a slice of context's
	It("will return a slice of context's", func() {

		// Add context entries to ContextList
		addContextEntries := func(contextList *proton.ContextList, entries []TableEntry) []*proton.Context {
			contextSlice := make([]*proton.Context, 0)
			for _, entry := range entries {
				contextData := entry.Parameters[0].(ContextData)
				context, _ := proton.NewContext(contextData.id)
				contextList.AddContext(context)
				contextSlice = append(contextSlice, context)
			}
			return contextSlice
		}

		contextList := proton.NewContextList()
		contextSlice := addContextEntries(contextList, ValidContextEntrySlice)
		Expect(contextList.ContextSlice()).To(HaveLen(len(ValidContextEntrySlice)))

		for _, context := range contextSlice {
			Expect(contextList.ContextSlice()).To(ContainElement(context))
		}
	})

})
