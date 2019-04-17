package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// ContextData ...
type ContextData struct {
	id          string
	expectedErr error
}

var (

	// InvalidContextEntrySlice ...
	InvalidContextEntrySlice = []TableEntry{
		Entry("", ContextData{"", proton.ErrContextIDUndefined}),                // 0
		Entry("", ContextData{"id ", proton.ErrContextIDContainsWhitespace}),    // 1
		Entry("", ContextData{"\tid", proton.ErrContextIDContainsWhitespace}),   // 5
		Entry("", ContextData{"\nid", proton.ErrContextIDContainsWhitespace}),   // 6
		Entry("", ContextData{"\r\nid", proton.ErrContextIDContainsWhitespace}), // 7
	}

	// ValidContextEntrySlice ...
	ValidContextEntrySlice = []TableEntry{
		Entry("", ContextData{"id", nil}),                         // 0
		Entry("", ContextData{"the_id", nil}),                     // 1
		Entry("", ContextData{"ID", nil}),                         // 2
		Entry("", ContextData{"the-id", nil}),                     // 3
		Entry("", ContextData{"~@*^£*(", nil}),                    // 4
		Entry("", ContextData{"F{^v@7Er4Fq&}CZ[", nil}),           // 5
		Entry("", ContextData{"%&ROPEEGG#_JACKXBOX#tokyoZ", nil}), // 6
	}
)

// Describe Context ...
var _ = Describe("Context", func() {

	// Creating an invalid context ...
	DescribeTable("creating an invalid context",
		func(contextData ContextData) {
			context, err := proton.NewContext(contextData.id)
			Expect(err).To(Equal(contextData.expectedErr))
			Expect(context).To(BeNil())
		}, InvalidContextEntrySlice...,
	)

	// Creating a valid context ...
	DescribeTable("creating a valid context",
		func(contextData ContextData) {
			context, err := proton.NewContext(contextData.id)
			Expect(err).To(BeNil())
			Expect(context).ToNot(BeNil())
			Expect(context.ID()).To(Equal(proton.String(contextData.id)))
		}, ValidContextEntrySlice...,
	)

})
