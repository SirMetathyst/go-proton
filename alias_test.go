package proton_test

import (
	proton "github.com/SirMetathyst/go-proton"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// AliasData ...
type AliasData struct {
	id          string
	value       string
	expectedErr error
}

var (

	// Valid Alias ...
	ValidAlias, _ = CreateAliasFromEntry(ValidAliasEntrySlice[0])

	// InvalidAliasEntrySlice ...
	InvalidAliasEntrySlice = []TableEntry{
		Entry("", AliasData{"", "", proton.ErrAliasIDUndefined}),                       // 0
		Entry("", AliasData{" ", "", proton.ErrAliasIDContainsWhitespace}),             // 0
		Entry("", AliasData{"", "value", proton.ErrAliasIDUndefined}),                  // 1
		Entry("", AliasData{"id ", "value", proton.ErrAliasIDContainsWhitespace}),      // 2
		Entry("", AliasData{"id", "  value", proton.ErrAliasValueContainsWhitespace}),  // 3
		Entry("", AliasData{"id", "", proton.ErrAliasValueUndefined}),                  // 4
		Entry("", AliasData{"\tid", "\tvalue", proton.ErrAliasIDContainsWhitespace}),   // 5
		Entry("", AliasData{"\nid", "\nvalue", proton.ErrAliasIDContainsWhitespace}),   // 6
		Entry("", AliasData{"\rid", "\r\nvalue", proton.ErrAliasIDContainsWhitespace}), // 7
	}

	// ValidAliasEntrySlice ...
	ValidAliasEntrySlice = []TableEntry{
		Entry("", AliasData{"id", "value", nil}),                                    // 0
		Entry("", AliasData{"the_id", "the_value", nil}),                            // 1
		Entry("", AliasData{"ID", "VALUE", nil}),                                    // 2
		Entry("", AliasData{"the-id", "the-value", nil}),                            // 3
		Entry("", AliasData{"~@*^£*(", "jd*(&)", nil}),                              // 4
		Entry("", AliasData{"F{^v@7Er4Fq&}CZ[", "q*$h+Bzcs6Z*8KgP", nil}),           // 5
		Entry("", AliasData{"%&ROPEEGG#_JACKXBOX#tokyoZ", "Y=Gmuz*e4KX-*j36", nil}), // 6
	}
)

// CreateAliasFromAliasData ...
func CreateAliasFromAliasData(aliasData AliasData) (*proton.Alias, error) {
	return proton.NewAlias(aliasData.id, aliasData.value)
}

// CreateAliasFromEntry
func CreateAliasFromEntry(entry TableEntry) (*proton.Alias, error) {
	return CreateAliasFromAliasData(entry.Parameters[0].(AliasData))
}

// Describe Alias ...
var _ = Describe("Alias", func() {

	// Creating an invalid alias ...
	DescribeTable("creating an invalid alias",
		func(aliasData AliasData) {
			alias, err := CreateAliasFromAliasData(aliasData)
			Expect(err).To(Equal(aliasData.expectedErr))
			Expect(alias).To(BeNil())
		}, InvalidAliasEntrySlice...,
	)

	// Creating a valid alias ...
	DescribeTable("creating a valid alias",
		func(aliasData AliasData) {
			alias, err := CreateAliasFromAliasData(aliasData)
			Expect(err).To(BeNil())
			Expect(alias).ToNot(BeNil())
			Expect(alias.ID()).To(Equal(proton.String(aliasData.id)))
			Expect(alias.Value()).To(Equal(proton.String(aliasData.value)))
		}, ValidAliasEntrySlice...,
	)

})
